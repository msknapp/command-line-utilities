package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

func Version() *cobra.Command {
    command := &cobra.Command{
        Use: "version",
        Short: "print the application version and then exit",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println(version)
        },
    }
    return command
}