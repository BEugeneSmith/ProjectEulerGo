package goutils

// IsDiv tests whether a number is even
func IsDiv(n, d int) bool {
	if n == 0 || d == 0 {
		return false
	} else if n%d == 0 {
		return true
	} else {
		return false
	}
}

// IsPrime tests whether a number is prime
func IsPrime(n int) bool {
	if n == 0 {
		return false
	} else if n <= 2 {
		return true
	} else if IsDiv(n, 2) {
		return false
	}
	for i := 3; i < int(n); i += 2 {
		if IsDiv(n, i) {
			return false
		}
	}
	return true
}

// IsPalindrome tests if a string is palindromic
func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	if s == string(runes) {
		return true
	}
	return false
}

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

// GetFactors returns a list of a number's factors
func GetFactors(n int) []int {
	var factors []int
	if IsDiv(n, 2) {
		factors = append(factors, 2)
		for i := 2; i <= (n / 2); i++ {
			if IsDiv(n, i) {
				factors = append(factors, i)
			}
		}
		factors = factors[1:]

	} else {
		for i := 3; i <= (n / 2); i += 2 {
			if IsDiv(n, i) {
				factors = append(factors, i)
			}
		}
	}

	return factors
}
