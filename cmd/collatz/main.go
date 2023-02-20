package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/msknapp/command-line-utilities/pkg/collatz"
	"github.com/spf13/cobra"
)

func main() {
	var delim string
	c := &cobra.Command{
		Use:   "collatz",
		Short: "prints a collatz sequence",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "pass the initial value as an argument")
				os.Exit(1)
			}
			x, e := strconv.Atoi(args[0])
			if e != nil {
				fmt.Fprintf(os.Stderr, e.Error())
				os.Exit(1)
			}
			if x < 1 {
				fmt.Fprintf(os.Stderr, "the initial value must be positive")
				os.Exit(1)
			}
			fmt.Printf("%d", x)
			for x > 1 {
				fmt.Printf(delim)
				x2 := collatz.NextCollatz(x)
				fmt.Printf("%d", x2)
				x = x2
			}
			fmt.Println()
		},
	}
	c.Flags().StringVarP(&delim, "delimiter", "d", ", ", "sets the delimiter")
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
