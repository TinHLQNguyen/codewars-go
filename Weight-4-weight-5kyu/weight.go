package weight

import (
	"sort"
	"strconv"
	"strings"
)

// NOTE: Given a string with the weights of in normal order,
// give this string ordered by "weights" of these numbers?
// The weight of a number will be from now on the sum of its digits.
// For example 99 will have "weight" 18, 100 will have "weight" 1 so in the list 100 will come before 99
// When two numbers have the same "weight", let us class them as if they were strings (alphabetical ordering) and not numbers
// 180 is before 90 since, having the same "weight" (9), it comes before as a string.
// All numbers in the list are positive numbers and the list can be empty.

// / we use cheatWeight as keys, any realWeight with same keys will be added to the same key.
type cheatWeightMap map[int][]string

func OrderWeight(input string) string {
	realWeights := strings.Fields(input)
	cheatWeights := []int{}
	cheatWeightsMapping := make(cheatWeightMap)
	cheatOrderWeight := []string{}

	for _, weight := range realWeights {
		cheatWeight := convertCheatWeight(weight)
		if _, ok := cheatWeightsMapping[cheatWeight]; !ok {
			cheatWeights = append(cheatWeights, cheatWeight)
			cheatWeightsMapping[cheatWeight] = []string{weight}
		} else {
			cheatWeightsMapping[cheatWeight] = append(cheatWeightsMapping[cheatWeight], weight)
		}
	}
	sort.Ints(cheatWeights)
	for _, cheat := range cheatWeights {
		sort.Strings(cheatWeightsMapping[cheat])
		cheatOrderWeight = append(cheatOrderWeight, cheatWeightsMapping[cheat]...)
	}
	return strings.Join(cheatOrderWeight, " ")
}

func convertCheatWeight(weight string) int {
	// NOTE: this work because we're dealing with only number
	sum := 0
	for _, digit := range weight {
		val, _ := strconv.Atoi(string(digit))
		sum += val
	}
	return sum
}
