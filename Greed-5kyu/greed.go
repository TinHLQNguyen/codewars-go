package greed

// NOTE: Greed is a dice game played with five six-sided dice. You throw the dices ant get the Score.
// Below are the rules
// Three 1's => 1000 points
// Three 6's =>  600 points
// Three 5's =>  500 points
// Three 4's =>  400 points
// Three 3's =>  300 points
// Three 2's =>  200 points
// One   1   =>  100 points
// One   5   =>   50 points
// A single die can only be counted once in each roll.
// For example, a given "5" can only count as part of a triplet (contributing to the 500 points)
// or as a single 50 points, but not both in the same roll.

func Score(rolls [5]int) int {
	return 0
}
