// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package tally

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"

	"storj.io/storj/pkg/accounting/accountingdb"
	"storj.io/storj/pkg/dht"
	"storj.io/storj/pkg/kademlia"
	"storj.io/storj/pkg/node"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pointerdb"
	"storj.io/storj/pkg/utils"
	"storj.io/storj/storage"
)

// Tally is the service for accounting for data stored on each storage node
type Tally interface {
	Run(ctx context.Context) error
}

type tally struct {
	pointerdb  *pointerdb.Server
	overlay    pb.OverlayServer
	kademlia   *kademlia.Kademlia
	limit      int
	logger     *zap.Logger
	ticker     *time.Ticker
	nodes      map[string]int64
	nodeClient node.Client
	db         *accountingdb.Database
}

func newTally(ctx context.Context, db *accountingdb.Database, pointerdb *pointerdb.Server, overlay pb.OverlayServer, kademlia *kademlia.Kademlia, limit int, interval time.Duration, logger *zap.Logger) (*tally, error) {
	rt, err := kademlia.GetRoutingTable(ctx)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	self := rt.Local()
	identity, err := node.NewFullIdentity(ctx, 12, 4) //TODO: what values go here?
	client, err := node.NewNodeClient(identity, self, kademlia)
	if err != nil {
		return nil, Error.Wrap(err)
	}
	return &tally{
		pointerdb:  pointerdb,
		overlay:    overlay,
		kademlia:   kademlia,
		limit:      limit,
		logger:     logger,
		ticker:     time.NewTicker(interval),
		nodes:      make(map[string]int64),
		nodeClient: client,
		db:         db,
	}, nil
}

// Run the tally loop
func (t *tally) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	for {
		err = t.identifyActiveNodes(ctx)
		if err != nil {
			zap.L().Error("Tally failed", zap.Error(err))
		}

		select {
		case <-t.ticker.C: // wait for the next interval to happen
		case <-ctx.Done(): // or the tally is canceled via context
			err = t.db.Close()
			if err != nil {
				zap.L().Error("error closing connection to accountingdb", zap.Error(err))
			}
			return ctx.Err()
		}
	}
}

// identifyActiveNodes iterates through pointerdb and identifies nodes that have storage on them
func (t *tally) identifyActiveNodes(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	t.logger.Debug("entering pointerdb iterate")
	err = t.pointerdb.Iterate(ctx, &pb.IterateRequest{Recurse: true},
		func(it storage.Iterator) error {
			var item storage.ListItem
			lim := t.limit
			if lim <= 0 || lim > storage.LookupLimit {
				lim = storage.LookupLimit
			}
			for ; lim > 0 && it.Next(&item); lim-- {
				pointer := &pb.Pointer{}
				err = proto.Unmarshal(item.Value, pointer)
				if err != nil {
					return Error.Wrap(err)
				}
				pieces := pointer.Remote.RemotePieces
				var nodeIDs []dht.NodeID
				for _, p := range pieces {
					nodeIDs = append(nodeIDs, node.IDFromString(p.NodeId))
				}
				online, err := t.categorize(ctx, nodeIDs)
				if err != nil {
					return Error.Wrap(err)
				}
				err = t.tallyAtRestStorage(ctx, pointer, online)
				if err != nil {
					return Error.Wrap(err)
				}
			}
			return nil
		},
	)
	return t.updateRawTable(ctx)
}

func (t *tally) categorize(ctx context.Context, nodeIDs []dht.NodeID) (online []*pb.Node, err error) {
	responses, err := t.overlay.BulkLookup(ctx, utils.NodeIDsToLookupRequests(nodeIDs))
	if err != nil {
		return []*pb.Node{}, err
	}
	nodes := utils.LookupResponsesToNodes(responses)
	for _, n := range nodes {
		if n != nil {
			online = append(online, n)
		} else {
			t.nodes[n.Id] = 0
		}
	}
	return online, nil
}

func (t *tally) tallyAtRestStorage(ctx context.Context, pointer *pb.Pointer, nodes []*pb.Node) error {
	segmentSize := pointer.GetSegmentSize()
	minReq := pointer.Remote.Redundancy.GetMinReq()
	if minReq <= 0 {
		return Error.New("minReq must be an int greater than 0")
	}
	pieceSize := segmentSize / int64(minReq)
	for _, n := range nodes {
		connected, err := t.nodeClient.Ping(ctx, *n)
		if connected {
			val, ok := t.nodes[n.Id]
			if ok {
				t.nodes[n.Id] = val + pieceSize
			} else {
				t.nodes[n.Id] = pieceSize
			}
		}
		if !connected || err != nil {
			zap.L().Error("ping failed for node: " + n.Id)
			t.nodes[n.Id] = 0
			continue
		}
	}
	return nil
}

func (t *tally) updateRawTable(ctx context.Context) error {
	//TODO
	return nil
}
