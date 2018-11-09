package main

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "setup peers",
		RunE:  cmdSetup,
	}
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

func cmdSetup(cmd *cobra.Command, args []string) (err error) {
	processes := NewProcesses(1, 100)

	ctx, cleanup := NewCLIContext(context.Background())
	defer cleanup()

	return processes.Setup(ctx)
}
