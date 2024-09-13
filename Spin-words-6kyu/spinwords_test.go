package spinwords_test

import (
	"spinwords"
	"testing"
)

func TestSpinSingleWord(t *testing.T) {
	cases := []struct {
		input, expectedOutput string
	}{
		{"Welcome", "emocleW"},
		{"to", "to"},
		{"CodeWars", "sraWedoC"},
		{"", ""},
	}
	for _, c := range cases {
		got := spinwords.SpinWords(c.input)
		assertEqual(t, got, c.expectedOutput)
	}
}

func TestSpinMultipleWords(t *testing.T) {
	cases := []struct {
		input, expectedOutput string
	}{
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"Burgers are my favorite fruit", "sregruB are my etirovaf tiurf"},
		{"Pizza is the best vegetable", "azziP is the best elbategev"},
	}
	for _, c := range cases {
		got := spinwords.SpinWords(c.input)
		assertEqual(t, got, c.expectedOutput)
	}
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("WANT %q, GOT %q", want, got)
	}
}
