package lms

// GetTriangularNumber finds the Nth triangular number.
// Formule Triangular numbers: https://www.mathsisfun.com/algebra/triangular-numbers.html.
// Rule: xn = n(n+1)/2
func GetTriangularNumber(nth int) int {
	return nth * (nth + 1) / 2
}

func FindDivsors() {
	panic("Not implemented")
}
