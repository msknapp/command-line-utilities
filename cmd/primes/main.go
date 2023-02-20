package main

import (
	"fmt"
	"os"

	"github.com/msknapp/command-line-utilities/pkg/primes"
	"github.com/spf13/cobra"
)

func main() {
	var theInt int64
	var theStartNumber int64
	var theStopNumber int64
	var theStartIndex int64
	var theStopIndex int64
	var printIt bool
	delimiter := " "
	perLine := -1
	c := &cobra.Command{
		Use:   "check primes or print them",
		Short: "Checks if numbers are primes, or prints a sequence of them.",
		Run: func(cmd *cobra.Command, args []string) {
			if theInt < 0 {
				fmt.Fprint(os.Stderr, "the integer to check must be positive")
				os.Exit(1)
			}
			if theInt > 0 {
				isPrime := primes.IsPrime(int(theInt))
				if printIt {
					fmt.Printf("%t\n", isPrime)
					os.Exit(0)
				} else {
					code := 1
					if isPrime {
						code = 0
					}
					os.Exit(code)
				}
			}
			// they want a range of integers.
			if theStartNumber > -1 && theStopNumber > theStartNumber {
				startOfLine := true
				count := 0
				for i := theStartNumber; i < theStopNumber; i++ {
					if primes.IsPrime(int(i)) {
						if !startOfLine {
							fmt.Print(delimiter)
						}
						fmt.Printf("%d", i)
						startOfLine = false
						count++
						if perLine > 0 && count%perLine == 0 {
							fmt.Println()
							startOfLine = true
						}
					}
				}
				fmt.Println()
				os.Exit(0)
			}
			// they want integers by their index
			str := primes.NewPrimeStream()
			startOfLine := true
			for i := 0; i < int(theStopIndex); i++ {
				v := str.Next()
				if i >= int(theStartIndex) {
					if !startOfLine {
						fmt.Print(delimiter)
					}
					fmt.Printf("%d", v)
					startOfLine = false
					if perLine > 0 && (i-int(theStartIndex))%perLine == 0 {
						fmt.Println()
						startOfLine = true
					}
				}
			}
			fmt.Println()
		},
	}
	c.Flags().Int64VarP(&theInt, "integer", "i", 0, "the integer to assert that it is a prime")
	c.Flags().Int64VarP(&theStartNumber, "from", "f", 0, "the integer to start at")
	c.Flags().Int64VarP(&theStopNumber, "to", "t", 0, "the integer to stop at, exclusive")
	c.Flags().Int64VarP(&theStartIndex, "start-index", "S", 0, "the index to start at")
	c.Flags().Int64VarP(&theStopIndex, "stop-index", "P", 0, "the index to stop at")
	c.Flags().BoolVarP(&printIt, "print", "p", false, "print true or false instead of changing the exit status code")
	c.Flags().StringVarP(&delimiter, "delimiter", "d", " ", "the delimiter between values")
	c.Flags().IntVarP(&perLine, "per-line", "l", -1, "the number of integers to print on one line.  by default it never wraps onto a new line.")
	if err := c.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
