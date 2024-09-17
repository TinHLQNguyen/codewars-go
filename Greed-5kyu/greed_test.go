package greed_test

import (
	"greed"
	"testing"
)

func TestGreed(t *testing.T) {
	cases := []struct {
		name     string
		rolls    [5]int
		maxScore int
	}{
		{"weird dice sneak in", [5]int{9, 3, 4, 6, 2}, 0},
		{"worthless rolls", [5]int{2, 3, 4, 6, 2}, 0},
		{"should value this triplet correctly", [5]int{4, 4, 4, 3, 3}, 400},
		{"should value only triplet", [5]int{6, 6, 6, 6, 6}, 600},
		{"should value this shuffled triplet correctly", [5]int{4, 3, 4, 4, 3}, 400},
		{"should only use as triplet", [5]int{4, 4, 4, 4, 4}, 400},
		{"should use for all applicable rules", [5]int{1, 1, 1, 1, 1}, 1200},
		{"should value this mixed set correctly", [5]int{2, 4, 4, 5, 4}, 450},
		{"only single rewards", [5]int{5, 1, 1, 5, 2}, 300},
		{"random test", [5]int{1, 2, 3, 4, 6}, 100},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := greed.Score(c.rolls)
			if got != c.maxScore {
				t.Errorf("GOT %d, WANT %d", got, c.maxScore)
			}
		})
	}
}
