package greed

import (
	"errors"
)

// NOTE: Greed is a dice game played with five six-sided dice. You throw the dices ant get the Score.
// Below are the rules
// Three 1's => 1000 points -> Rule 1
// Three 6's =>  600 points -> Rule 2
// Three 5's =>  500 points -> Rule 3
// Three 4's =>  400 points -> Rule 4
// Three 3's =>  300 points -> Rule 5
// Three 2's =>  200 points -> Rule 6
// One   1   =>  100 points -> Rule 7
// One   5   =>   50 points -> Rule 8
// A single die can only be counted once in each roll.
// For example, a given "5" can only count as part of a triplet (contributing to the 500 points)
// or as a single 50 points, but not both in the same roll.

var tripletReward = map[int]int{
	1: 1000,
	2: 200,
	3: 300,
	4: 400,
	5: 500,
	6: 600,
}

var singleReward = map[int]int{
	1: 100,
	2: 0,
	3: 0,
	4: 0,
	5: 50,
	6: 0,
}

type DiceRolls struct {
	FaceCountsLeft map[int]int
	Rolls          [5]int
}

func NewRolls(rolls [5]int) (DiceRolls, error) {
	faceCounts := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	for _, roll := range rolls {
		if !isDice(roll) {
			return DiceRolls{}, errors.New("invalid dice face found")
		}
		faceCounts[roll]++
	}
	return DiceRolls{faceCounts, rolls}, nil
}

func (d *DiceRolls) checkTriplet() int {
	var addedScore int
	for key, score := range tripletReward {
		if _, ok := d.FaceCountsLeft[key]; ok {
			if d.FaceCountsLeft[key] >= 3 {
				addedScore += score
				d.FaceCountsLeft[key] -= 3
				// this works b/c we only have 5 dices
				break
			}
		}
	}
	return addedScore
}

func (d *DiceRolls) checkSingle() int {
	var addedScore int
	for key, score := range singleReward {
		if count, ok := d.FaceCountsLeft[key]; ok {
			if count > 0 {
				addedScore += count * score
				d.FaceCountsLeft[key] -= count
			}
		}
	}
	return addedScore
}

func Score(rolls [5]int) int {
	var score int
	newRolls, err := NewRolls(rolls)
	if err != nil {
		score = 0
	}
	score += newRolls.checkTriplet()
	score += newRolls.checkSingle()
	return score
}

func isDice(val int) bool {
	return (val > 0) && (val < 6)
}
