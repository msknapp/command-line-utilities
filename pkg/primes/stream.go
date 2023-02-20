package primes

type SixNApprox struct {
	n, i int
}

func NewSixNApprox() *SixNApprox {
	return &SixNApprox{
		i: 1,
	}
}

func (a *SixNApprox) Advance() {
	a.n += (a.i + 1) / 2
	a.i *= -1
}

func (a *SixNApprox) Get() int {
	return 6*a.n + a.i
}

func (a *SixNApprox) Next() int {
	a.Advance()
	return a.Get()
}

type PrimeStream struct {
	approx *SixNApprox
	index  int
}

func NewPrimeStream() *PrimeStream {
	tmp := &PrimeStream{
		approx: NewSixNApprox(),
		index:  -1,
	}
	tmp.approx.n = 16
	tmp.approx.i = 1
	return tmp
}

func (s *PrimeStream) Advance() {
	if s.index < len(QUICK_PRIMES)+len(PRIMES_UNDER_97) {
		s.index++
	} else {
		s.approx.Advance()
		for !IsPrime(s.approx.Get()) {
			s.approx.Advance()
		}
	}
}

func (s *PrimeStream) Get() int {
	if s.index < len(QUICK_PRIMES) {
		return QUICK_PRIMES[s.index]
	} else if s.index < len(QUICK_PRIMES)+len(PRIMES_UNDER_97) {
		return PRIMES_UNDER_97[s.index-len(QUICK_PRIMES)]
	}
	return s.approx.Get()
}

func (s *PrimeStream) Next() int {
	s.Advance()
	return s.Get()
}
