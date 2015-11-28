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
	} else if n < 2 {
		return true
	} else if IsDiv(n, 2) {
		return false
	}
	for i := 3; i < int(n/2); i += 2 {
		if IsDiv(n, i) {
			return false
		}
	}
	return true
}
