// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"flag"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/provider"
	"storj.io/storj/pkg/transport"
)

var (
	// Addr is the address of Capt Planet from command flags
	Addr = flag.String("address", "[::1]:7778", "address of captplanet to inspect")

	// ErrIdentity is for errors during identity creation for this CLI
	ErrIdentity = errs.Class("error creating identity:")

	// ErrInspectorDial throws when there are errors dialing the inspector server
	ErrInspectorDial = errs.Class("error dialing inspector server:")

	// ErrRequest is for gRPC request errors after dialing
	ErrRequest = errs.Class("error processing request:")

	rootCmd = &cobra.Command{
		Use:   "pointerdb",
		Short: "commands for exploring pointerdb",
	}
	getCmd = &cobra.Command{
		Use:   "get <path>",
		Short: "get pointer at a path",
	}
	listCmd = &cobra.Command{
		Use:   "list <start_after> <end_before> <limit>",
		Short: "list pointers from a starting path to a limit",
		Args:  cobra.MinimumNArgs(1),
		RunE:  ListPointers,
	}
)

// PointerInspector gives access to pointerdb
type PointerInspector struct {
	identity *provider.FullIdentity
	client   pb.PointerInspectorClient
}

// NewPointerInspector creates a new inspector server for accessing pointerdb
func NewPointerInspector(address string) (*PointerInspector, error) {
	ctx := context.Background()
	identity, err := provider.NewFullIdentity(ctx, 12, 4)
	if err != nil {
		return &PointerInspector{}, ErrIdentity.Wrap(err)
	}
	tc := transport.NewClient(identity)
	conn, err := tc.DialAddress(ctx, address)
	if err != nil {
		return &PointerInspector{}, ErrInspectorDial.Wrap(err)
	}

	c := pb.NewPointerInspectorClient(conn)

	return &PointerInspector{
		identity: identity,
		client:   c,
	}, nil
}

// ListPointers lists pointers
func ListPointers(*cobra.Command, []string) error {
	inspector, err := NewPointerInspector(*Addr)
	if err != nil {
		return ErrInspectorDial.Wrap(err)
	}
	_, err = inspector.client.ListPointers(context.Background(), &pb.ListPointersRequest{})
	if err != nil {
		return ErrRequest.Wrap(err)
	}
	return nil
}
