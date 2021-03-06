package goutils

import "math"

// FibonacciNumber returns a fibonacci sequence up to the lth element
func FibonacciNumber(l int) []int {
	seq := []int{1, 1}
	ix1 := 0
	ix2 := 1

	for len(seq) < (l + 1) {
		newIx := seq[ix1] + seq[ix2]
		seq = append(seq, newIx)
		ix1++
		ix2++
	}
	return seq
}

// FibonacciLimit returns a Fibonacci sequence up to a limit
func FibonacciLimit(l int) []int {
	seq := []int{1, 1}
	ix1 := 0
	ix2 := 1

	for seq[(len(seq)-1)] < l {
		newIx := seq[ix1] + seq[ix2]
		seq = append(seq, newIx)
		ix1++
		ix2++
	}
	return seq[:len(seq)-1]
}

// PrimeNumber returns a prime sequence up to the lth element
func PrimeNumber(l int) []int {
	seq := []int{2}
	for i := 3; len(seq) < l; i += 2 {
		if IsPrime(i) {
			seq = append(seq, i)
		}
	}
	return seq
}

// PrimeLimit returns a prime sequence up to the lth element
func PrimeLimit(l int) []int {
	seq := []int{2}

	for i := 3; seq[(len(seq)-1)] < l; i += 2 {
		if IsPrime(i) {
			seq = append(seq, i)
		}
	}
	return seq[:(len(seq) - 1)]
}

// SumSquare calculates the sum of squares for the first n natural numbers
func SumSquare(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		prod := int(math.Pow(float64(i), 2))
		sum += prod
	}
	return sum
}

// SquareSum calculates the square of sums for the first n natural numbers
func SquareSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// TriangleNumber returns a triangle sequence up to the lth element
// TODO: Needs to be tested
func TriangleNumber(l int) []int {
	seq := []int{1}

	for i := 1; i <= l; i++ {
		newIx := seq[(i-1)] + 1 + i
		seq = append(seq, newIx)
	}
	return seq
}

// TriangleLimit returns a prime sequence up to a limit
// TODO: Needs to be tested
func TriangleLimit(l int) []int {
	seq := []int{1}

	for i := 1; seq[len(seq)-1] <= l; i++ {
		newIx := seq[(i-1)] + 1 + i
		seq = append(seq, newIx)
	}
	return seq
}
