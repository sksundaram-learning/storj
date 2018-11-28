// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package tally

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestIdentifyActiveNodes(t *testing.T) {

}

func TestCategorize(t *testing.T) {
	// logger := zap.NewNop()
	// pointerdb := pointerdb.NewServer(teststore.New(), &overlay.Cache{}, logger, pointerdb.Config{}, nil)

	// const N = 50
	// nodes := []*pb.Node{}
	// nodeIDs := []dht.NodeID{}
	// expectedOnline := []*pb.Node{}
	// for i := 0; i < N; i++ {
	// 	str := strconv.Itoa(i)
	// 	n := &pb.Node{Id: str, Address: &pb.NodeAddress{Address: str}}
	// 	nodes = append(nodes, n)
	// 	if i%(rand.Intn(5)+2) == 0 {
	// 		id := node.IDFromString("id" + str)
	// 		nodeIDs = append(nodeIDs, id)
	// 	} else {
	// 		id := node.IDFromString(str)
	// 		nodeIDs = append(nodeIDs, id)
	// 		expectedOnline = append(expectedOnline, n)
	// 	}
	// }
	// overlayServer := mocks.NewOverlay(nodes)
	// rootdir, cleanup := mktempdir(t, "kademlia")
	// defer cleanup()
	// kad, err := kademlia.NewKademlia(node.IDFromString("foo"), []pb.Node{}, "127.0.0.1:8080", nil, &provider.FullIdentity{}, rootdir, 1)
	// assert.NoError(t, err)
	// limit := 0
	// interval := time.Second

	// tally := newTally(ctx, pointerdb, overlayServer, kad, limit, logger, interval)
	// online, err := tally.categorize(ctx, nodeIDs)
	// assert.NoError(t, err)
	// assert.Equal(t, expectedOnline, online)
}

func TestTallyAtRestStorage(t *testing.T) {

}

func TestUpdateRawTable(t *testing.T) {

}

// func mktempdir(t *testing.T, dir string) (string, func()) {
// 	rootdir, err := ioutil.TempDir("", dir)
// 	assert.NoError(t, err)
// 	cleanup := func() {
// 		assert.NoError(t, os.RemoveAll(rootdir))
// 	}
// 	return rootdir, cleanup
// }
