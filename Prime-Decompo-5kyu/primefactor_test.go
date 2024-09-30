package primefactor_test

import (
	"primefactor"
	"testing"
)

func TestPrimeFactor(t *testing.T) {
	cases := []struct {
		name, pfd string
		input     int
	}{
		{"given 2", "(2)", 2},
		{"only 1 prime 1", "(23)", 23},
		{"only 1 prime 1-1", "(5)", 5},
		{"only 1 prime 2", "(3**2)", 9},
		{"random SMALL number 1", "(11)(23)", 253},
		{"random SMALL number 2", "(2**2)(3**3)(5)", 540},
		{"random SMALL number 3", "(2**2)(5)(3967)", 79340},
		{"random BIG number 1", "(2**2)(3**3)(5)(7)(11**2)(17)", 7775460},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := primefactor.PrimeFactors(c.input)
			if got != c.pfd {
				t.Errorf("GOT %q, WANT %q", got, c.pfd)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	cases := []struct {
		name  string
		input [2]int
		gbd   int
	}{
		{"normal case 1", [2]int{8, 12}, 4},
		{"normal case 2", [2]int{54, 24}, 6},
		{"denominator case", [2]int{6, 3}, 3},
		{"coprime case", [2]int{3, 4}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := primefactor.GCD(c.input[0], c.input[1])
			if got != c.gbd {
				t.Errorf("GOT %q, WANT %q", got, c.gbd)
			}
		})
	}
}
