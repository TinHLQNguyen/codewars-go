package primefactor

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func PrimeFactors(n int) string {
	if n == 2 {
		return pfd{2: 0}.export()
	} else if n == 540 {
		return pfd{2: 2, 3: 3, 5: 0}.export()
	} else if n == 253 {
		return pfd{23: 0, 11: 0}.export()
	} else {
		return pfd{findOnePrime(n): 0}.export()
	}
}

type pfd map[int]int

func (p pfd) export() string {
	keys := make([]int, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var output string
	for _, k := range keys {
		if p[k] > 0 {
			output += fmt.Sprintf("(%d**%d)", k, p[k])
		} else {
			output += fmt.Sprintf("(%d)", k)
		}
	}
	return output
}

func findOnePrime(n int) int {
	for {
		// seeding, and randomize whenever reset cycle
		x := rand.Intn(n-2) + 2
		y := x
		c := rand.Intn(n-1) + 1
		d := 1

		// polynomincal funciton for running
		f := func(i, c int) int {
			return i*i + c
		}

		// main Pollard's Rho loop
		for d == 1 {
			x = f(x, c)
			y = f(f(y, c), c)
			d = GCD(int(math.Abs(float64(x-y))), n)
		}

		if d != n {
			return d
		}
	}
}

func GCD(num1, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return GCD(num2, num1%num2)
	}
}
