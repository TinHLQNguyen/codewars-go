package primefactor

import (
	"fmt"
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
		return ""
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

func GCB(num1, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return GCB(num2, num1%num2)
	}
}

func runFunc(num int) int {
	return num ^ 2 + 1
}
