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
		{"random number 1", "(2**2)(3**3)(5)(7)(11**2)(17)", 7775460},
		{"random number 2", "(2**2)(5)(3967)", 79340},
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
