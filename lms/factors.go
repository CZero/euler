package lms

import (
	"sort"
)

// FindFactors finds the factors of a number. You should be able to do some voodoo with PrimeFactors, but
// I'm not using that yet.
func FindFactors(input int) []int {
	found := make(map[int]bool, input/2+2)
	answer := []int{}
	for i := 1; !found[i]; i++ {
		if !found[i] {
			if input%i == 0 {
				answer = append(answer, i, input/i)
				found[i] = true
				found[input/i] = true
			}
		}
	}
	sort.Ints(answer)
	return answer
}
