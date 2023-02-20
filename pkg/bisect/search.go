package bisect

type MODE int

const (
	EQUAL MODE = iota
	BEFORE
	AFTER
)

func IndexOf(x []int, v int) int {
	i, ok := closestIndex(x, v)
	if ok {
		return i
	}
	return -1
}

func Contains(x []int, v int) bool {
	_, ok := closestIndex(x, v)
	return ok
}

func IndexOfOrBefore(x []int, v int) int {
	i, _ := closestIndex(x, v)
	if i == -2 {
		return len(x)
	}
	if i == -1 {
		return 0
	}
	return i
}

func closestIndex(x []int, v int) (int, bool) {
	if x == nil || len(x) < 1 {
		return -1, false
	}
	if x[0] == v {
		return 0, true
	} else if x[0] > v {
		return -1, false
	}
	if x[len(x)-1] == v {
		return len(x) - 1, true
	} else if x[len(x)-1] < v {
		return -2, false
	}
	lo := 0
	hi := len(x)
	var mid int
	for hi-lo > 1 {
		mid = (lo + hi) / 2
		mv := x[mid]
		if v == mv {
			return mid, true
		}
		if v > mv {
			lo = mid
		}
		if v < mv {
			hi = mid
		}
	}
	return hi, false
}
