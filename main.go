package main

import (
	"fmt"
	"os"

	"github.com/hiden2000/taskmaster/internal/storage"
	"github.com/hiden2000/taskmaster/pkg/cli"
)

func main() {
	// Init storage
	store := storage.NewStorage()

	// Init CLI
	taskCLI := cli.NewCLI(store)

	// Run CLI tool
	if err := taskCLI.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v\n", err)
		os.Exit(1)
	}

}
