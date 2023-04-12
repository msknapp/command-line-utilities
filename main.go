package main

import (
	"fmt"
	"os"

	"github.com/msknapp/command-line-utilities/cmd"
)

func main() {
	command := cmd.Root()
	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
