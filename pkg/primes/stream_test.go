package primes

import (
	"fmt"
	"testing"
	"time"
)

func TestStream(t *testing.T) {
	s := NewPrimeStream()
	for i := 1; i <= 1000; i++ {
		p := s.Next()
		fmt.Printf("%d, ", p)
		if i > 0 && i%20 == 0 {
			fmt.Println()
		}
	}
}

func TestChecks(t *testing.T) {
	x := time.Now()
	for i := 15000000; i <= 15001000; i++ {
		if IsPrime(i) {
			fmt.Printf("%d, ", i)
		}
	}
	d := time.Since(x)
	fmt.Printf("\ntook %d micros\n", d.Microseconds())
}
