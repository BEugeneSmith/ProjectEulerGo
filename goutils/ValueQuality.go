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
	rev := ReverseString(s)

	if s == rev {
		return true
	}
	return false
}

// IsEuclidianTriad tests whether a given ordered trio of numbers is a Euclidian triad
func IsEuclidianTriad(a, b, c int) bool {
	if a < b && b < c {
		return true
	}
	return false
}
