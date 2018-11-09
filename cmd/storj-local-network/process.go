// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sync/errgroup"

	"storj.io/storj/internal/processgroup"
)

type Processes struct {
	List []*Process
}

func NewProcesses(satelliteCount, storageNode int) *Processes {
	processes := &Processes{}

	for i := 0; i < satelliteCount; i++ {
		name := fmt.Sprintf("satellite/%d", i)
		process := NewProcess(name, "satellite", filepath.Join(".", defaultDir, "satellite", fmt.Sprint(i)))
		processes.List = append(processes.List, process)
	}

	for i := 0; i < storageNode; i++ {
		name := fmt.Sprintf("storage/%d", i)
		process := NewProcess(name, "storagenode", filepath.Join(".", defaultDir, "storagenode", fmt.Sprint(i)))
		processes.List = append(processes.List, process)
	}

	return processes
}

func (processes *Processes) all(fn func(process *Process) error) error {
	var group errgroup.Group
	for _, p := range processes.List {
		process := p

		process.Stdout.Hook(os.Stdout)
		process.Stderr.Hook(os.Stderr)

		group.Go(func() error {
			return fn(process)
		})
	}
	return group.Wait()
}

func (processes *Processes) Run(ctx context.Context) error {
	return processes.all(func(process *Process) error {
		return process.Run(ctx)
	})
}

func (processes *Processes) Setup(ctx context.Context) error {
	return processes.all(func(process *Process) error {
		return process.Setup(ctx)
	})
}

type Process struct {
	Name       string
	Directory  string
	Executable string

	Stdout Buffer
	Stderr Buffer
}

func NewProcess(name, executable, directory string) *Process {
	os.MkdirAll(directory, 0644)
	return &Process{
		Name:       name,
		Directory:  directory,
		Executable: executable,
	}
}

func (process *Process) Run(ctx context.Context) error {
	cmd := exec.Command(process.Executable, "run",
		"--base-path", process.Directory,
	)
	cmd.Dir = process.Directory
	cmd.Stdout, cmd.Stderr = &process.Stdout, &process.Stderr

	processgroup.Setup(cmd)

	err := cmd.Run()
	return err
}

func (process *Process) Setup(ctx context.Context) error {
	cmd := exec.Command(process.Executable, "setup",
		"--base-path", process.Directory,
	)
	cmd.Dir = process.Directory
	cmd.Stdout, cmd.Stderr = &process.Stdout, &process.Stderr

	processgroup.Setup(cmd)

	err := cmd.Run()
	return err
}
