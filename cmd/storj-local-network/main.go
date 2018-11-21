// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"storj.io/storj/pkg/process"
)

var (
	rootCmd = &cobra.Command{
		Use:   "storj-local-network",
		Short: "Storj full-network",
	}
	defaultDir = "local-network"
)

func main() {
	// process.Exec will load this for this command.
	runCmd.Flags().String("config", filepath.Join(defaultDir, "config.yaml"), "path to configuration")
	setupCmd.Flags().String("config", filepath.Join(defaultDir, "setup.yaml"), "path to configuration")
	process.Exec(rootCmd)
}
