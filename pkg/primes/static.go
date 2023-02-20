package primes

import (
	"math"
	"sync"

	"github.com/msknapp/command-line-utilities/pkg/bisect"
)

// IsPrimeContinuousCheck always correctly identifies a prime, but it may be slower than needed
// for larger numbers.  It checks if the number is divisible by any integer up to its square
// root.  The problem is it checks numbers that are not themselves primes, which is not necessary.
// if a number is not divisible by two, then there is no point in checking if it is divisible by four.
func IsPrimeContinuousCheck(x int) bool {
	if x < 2 {
		return false
	}
	rt := int(math.Sqrt(float64(x)))
	for i := 2; i <= rt; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// IsPrime attempts to be a very fast method of checking any integer to see if it is a prime.
// For prime numbers under 7921, it is very fast at checking them.
func IsPrime(x int) bool {
	if x <= 20 {
		// runs in O(3)
		return isPrimeUnder20(x)
	} else if x < 97 {
		// runs in O(4)
		return isPrimeFrom20to97(x)
	}
	// runs in O(8)
	if !isPrimeUnder529(x) {
		return false
	}
	if x < 529 {
		return true
	}
	// runs in O(16)
	if !possiblyPrimeUnder7927(x) {
		return false
	}
	if x < 7927 {
		return true
	}
	// now we get into a slower method of checking primes.
	if x < 5e5 {
		// it should run in O(root of x, divided by 3)
		return possiblyPrimeSixCheck(x, 16, false)
	}
	// after half a million, it could pay off to use some concurrency.
	return possiblyPrimeSixCheckConcurrent(x, 16, false)
}

// isPrimeUnder20 is always correct if x is less than 20.  It runs in O(3).
func isPrimeUnder20(x int) bool {
	return bisect.Contains(QUICK_PRIMES, x)
}

// isPrimeUnder20 is always correct if x is at least 20 and less than 97.
// It runs in O(4).
func isPrimeFrom20to97(x int) bool {
	return bisect.Contains(PRIMES_UNDER_97, x)
}

// isPrimeUnder529 determines if the input is a prime so long as it is under 367.
// It is always correct if x is under 367, and runs in O(8).
func isPrimeUnder529(x int) bool {
	return !doesSetContainAFactorOf(x, QUICK_PRIMES)
}

// possiblyPrimeUnder7927 returns true if the input does not have a factor that is a prime
// from 23 to 89.  If you combine this with isPrimeUnder367 and x is less than 7921,
// then it correctly detects all primes.
func possiblyPrimeUnder7927(x int) bool {
	return !doesSetContainAFactorOf(x, PRIMES_UNDER_97)
}

func isPrimeUnder7927(x int) bool {
	if !isPrimeUnder529(x) {
		return false
	}
	if x < 529 {
		return true
	}
	// runs in O(16)
	if !possiblyPrimeUnder7927(x) {
		return false
	}
	if x < 7921 {
		return true
	}
	return false
}

func failurePointOf(initialValue int, f func(int) bool) int {
	for i := initialValue; i < 1000*initialValue; i++ {
		xx := f(i)
		yy := IsPrimeContinuousCheck(i)
		if xx != yy {
			return i
		}
	}
	return -1
}

func doesSetContainAFactorOf(x int, primes []int) bool {
	rt := int(math.Sqrt(float64(x)))
	for _, p := range primes {
		if p > rt {
			break
		}
		if x%p == 0 {
			return true
		}
	}
	return false
}

// possiblyPrimeSixCheck determines if an integer is prime by checking if it is divisibly by
// any value of the function 6n+i, where i is +/- 1, and n > 0.  It stops when the value exceeds
// the square root of the input.  This method is correct as long as x > 5.  It is faster than
// the IsPrimeContinuousCheck method because it skips over all multiples of two or three.
// It does not check for division with primes below the initial 6n+i value.  If you combine
// it with other functions with lower prime checks, it will always be correct.  However,
// it will sometimes check if the input x has a factor which is not itself a prime.
// For very large primes, the numbers of factors that it must check can get prohibitively high,
// so the function may split the checks up into separate goroutines and check them in parallel.
func possiblyPrimeSixCheck(x int, initialN int, negativeI bool) bool {
	rt := int(math.Sqrt(float64(x)))
	n := initialN
	i := 1
	if negativeI {
		i = -1
	}
	v := 6*n + i
	for v <= rt {
		if x%v == 0 {
			return false
		}
		n += (i + 1) / 2
		i *= -1
		v = 6*n + i
	}
	return true
}

// possiblyPrimeSixCheckConcurrent is similar to possiblyPrimeSixCheck except it runs checks
// in parallel.  A suggestion is to use this when each routine needs to make over 64 checks,
// and there would be at least two such routines.  Why 64?  because I felt like it.
// So arbitrarily it should be chosen when there are at least 128 factors to check.  Since
// factors come two at a time for every six ints, it means we should use this if there are at
// least 128 * 3
func possiblyPrimeSixCheckConcurrent(x int, initialN int, negativeI bool) bool {
	rt := int(math.Sqrt(float64(x)))
	const blockSize = 64
	maxN := (rt / 6) + 1
	// you get about 10 groups when the input is over 15.5 million.
	groups := 1 + (maxN-initialN)/blockSize
	canBePrime := true
	wg := new(sync.WaitGroup)
	wg.Add(groups)
	for gr := 0; gr < groups; gr++ {
		go func(groupNumber int) {
			// prevent deadlock by always reporting when you are done.
			defer wg.Done()
			// if some other routine already said it found a factor, then stop immediately.
			if !canBePrime {
				return
			}
			n := initialN + (blockSize * groupNumber)
			stopN := initialN + (blockSize * (groupNumber + 1)) + 1
			i := -1
			for n < stopN {
				// there will be some unnecessary checks in here, where v is not actually a prime,
				// but it checks it anyways.
				v := 6*n + i
				if x%v == 0 {
					canBePrime = false
					break
				}
				if v > rt {
					break
				}
				n += (1 + i) / 2
				i *= -1
			}
		}(gr)
	}
	wg.Wait()
	return canBePrime
}

func ExpressAsSixN(x int) (int, int) {
	n := x / 6
	i := x % 6
	if i > 2 {
		i = 6 - i
		n++
	}
	return n, i
}
