package lms

import "math"

// Do all the things with primes!

// primesBelow is a primegenerator that finds primes below a number
func PrimesBelow(max int) (primes []int) {
	if max < 1 {
		return []int{0}
	}
	primes = []int{2}             // We don't have to look for 2.
	for a := 3; a < max; a += 2 { // Only check odd numbers, continue until we've reached the limit
		foundDivider := false
		highestRoot := math.Pow(float64(a), .5) + 1
		for _, b := range primes { // We will try if a is divisable by the allready found primes.
			if b > int(highestRoot) { // It's no use continuing above the root of the number.
				break
			}
			if a%b == 0 { // We've found a prime number that is a divisor, so this is no prime.
				foundDivider = true
				break
			}
		}
		if !foundDivider {
			primes = append(primes, a)
		}
	}
	return primes
}
