package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/msknapp/command-line-utilities/pkg/divisors"
	"github.com/spf13/cobra"
)

func main() {
	var proper bool
	var unsorted bool
	var delim string
	var printPrime bool
	c := &cobra.Command{
		Use:   "divisors",
		Short: "Determines the divisors or factors of a positive integer",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Fprintf(os.Stderr, "provide an integer argument\n")
				os.Exit(1)
			}
			z := args[0]
			n, e := strconv.Atoi(z)
			if e != nil {
				fmt.Fprintf(os.Stderr, "input is not an integer: %s\n", e.Error())
				os.Exit(1)
			}
			x := divisors.Divisors(n, proper, !unsorted)
			if len(x) < 1 {
				if printPrime {
					fmt.Printf("prime\n")
				}
				os.Exit(0)
			}
			for i, y := range x {
				if i > 0 {
					fmt.Printf(delim)
				}
				fmt.Printf("%d", y)
			}
			fmt.Println()
		},
	}
	c.Flags().BoolVarP(&proper, "proper", "p", false, "don't print 1")
	c.Flags().BoolVarP(&unsorted, "unsorted", "s", false, "don't sort factors, it may be slightly faster")
	c.Flags().BoolVarP(&printPrime, "printPrime", "P", false, "if it is a prime then print the word prime")
	c.Flags().StringVarP(&delim, "delimiter", "d", ", ", "Sets the delimiter between values")
	if err := c.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
