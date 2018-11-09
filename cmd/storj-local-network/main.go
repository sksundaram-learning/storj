// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"path/filepath"

	"github.com/spf13/cobra"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/process"
)

var (
	mon = monkit.Package()

	rootCmd = &cobra.Command{
		Use:   "captplanet",
		Short: "Captain Planet! With our powers combined!",
	}

	defaultDir = "local-network"
)

func main() {
	// process.Exec will load this for this command.
	runCmd.Flags().String("config", filepath.Join(defaultDir, "config.yaml"), "path to configuration")
	setupCmd.Flags().String("config", filepath.Join(defaultDir, "setup.yaml"), "path to configuration")
	process.Exec(rootCmd)
}
