package collatz

func NextCollatz(x int) int {
	if x < 2 {
		return -1
	}
	if x%2 == 0 {
		return x / 2
	} else {
		return 3*x + 1
	}
}
