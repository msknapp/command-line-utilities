package divisors

import (
	"math"
	"sort"
)

func Divisors(x int, proper bool, sorted bool) []int {
	if x < 2 {
		return []int{}
	}
	out := make([]int, 0, 6)
	CollectDivisors(x, proper, func(y int) {
		out = append(out, y)
	})
	if sorted {
		sort.Ints(out)
	}
	return out
}

func CollectDivisors(x int, proper bool, callback func(int)) {
	if x < 2 {
		return
	}
	rt := math.Sqrt(float64(x))
	rti := int(rt)
	if proper {
		callback(1)
	}
	for i := 2; i <= rti; i++ {
		if x%i == 0 {
			callback(i)
			if float64(i) < rt {
				callback(x / i)
			}
		}
	}
}
