package weight_test

import (
	"testing"
	"weight"
)

func TestWeight4Weight(t *testing.T) {
	cases := []struct {
		name, input, expectedOutput string
	}{
		{"Unique weight", "103 123 4444 99 2000", "2000 103 123 4444 99"},
		{"More whitespace", "  103 123 4444 99   2000  ", "2000 103 123 4444 99"},
		{"Same weight", "2000 10003 1234000 44444444 9999 11 11 22 123", "11 11 2000 10003 22 123 1234000 44444444 9999"},
		{"Empty", "", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := weight.OrderWeight(c.input)
			if got != c.expectedOutput {
				t.Errorf("GOT %q, WANT %q", got, c.expectedOutput)
			}
		})
	}
}
