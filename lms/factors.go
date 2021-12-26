package lms

import "sort"

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

// FindNumFactors finds the number of factors of a number. You should be able to do some voodoo with PrimeFactors, but
// I'm not using that yet.
func FindNumFactors(input int, primes []int) int {
	// This page explains how to use primefactors to calculate faster: https://www.geeksforgeeks.org/no-factors-n/
	// The formula: n! = (a1+1) x (a2+1) x (a3+1)……………(ak+1)
	// If you go through number theory, you will find an efficient way to find the number of factors. If we take a number,
	// say in this case 30, then the prime factors of 30 will be 2, 3, 5 with count of each of these being 1 time,
	// so total number of factors of 30 will be (1+1)*(1+1)*(1+1) = 8. https://www.geeksforgeeks.org/efficient-program-print-number-factors-n-numbers/
	var answer = 1
	if len(primes) < 2 {
		primes = PrimesBelow(input)
	}
	exponents := make([]int, len(primes))
	for i, prime := range primes {
		for {
			if input%prime == 0 {
				exponents[i]++
				input = input / prime
			} else {
				break
			}
		}
	}
	for _, exp := range exponents {
		if exp != 0 {
			answer = (exp + 1) * answer
		}
	}
	return answer
}

// Factorial returns the factorial(!) of a number (5!=1x2x3x4x5=120)
// https://www.mathsisfun.com/numbers/factorial.html
// The formula = n! = n × (n−1)!
func Factorial(input int) int {
	if input == 0 {
		return 1
	}
	answers := []int{1}
	for i := 2; i <= input; i++ {
		answers = append(answers, answers[i-2]*i)
	}
	return answers[len(answers)-1]
}
