package multiple_test

import (
	"multiple"
	"testing"
)

func TestMultipls(t *testing.T) {
	cases := []struct {
		caseName       string
		number, result int
	}{
		{"base case 23", 10, 23},
		{"case 0", 0, 0},
	}
	for _, c := range cases {
		got := multiple.Multiple3And5(c.number)
		if c.result != got {
			t.Errorf("Case %q, expect %d, got %d", c.caseName, c.result, got)
		}
	}
}
