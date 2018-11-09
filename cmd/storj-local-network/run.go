// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run peers",
		RunE:  cmdRun,
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}

func cmdRun(cmd *cobra.Command, args []string) (err error) {
	processes := NewProcesses(1, 100)

	ctx, cleanup := NewCLIContext(context.Background())
	defer cleanup()

	return processes.Run(ctx)
}
