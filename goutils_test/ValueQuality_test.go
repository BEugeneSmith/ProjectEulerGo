package testgoutils

import (
	"GoProjectEuler/goutils"
	"strconv"
	"testing"
)

// TestIsDiv tests the IsDiv function
func TestIsDiv(t *testing.T) {
	type IsDivTest struct {
		test [2]int // input
		expt bool   // expected result
	}

	var tt = []IsDivTest{
		{[2]int{4, 2}, true},
		{[2]int{9, 2}, false},
		{[2]int{10, 3}, false},
		{[2]int{0, 2}, false},
		{[2]int{2, 0}, false},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.IsDiv(tt[i].test[0], tt[i].test[1])
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("IsDiv test failed")
		}
	}
}

// TestIsPrime tests the IsPrime function
func TestIsPrime(t *testing.T) {
	var tt = []goutils.TTIB{
		{2, true},
		{5, true},
		{4, false},
		{9, false},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.IsPrime(tt[i].Test)
		testExp := tt[i].Expt

		if testExp != testIn {
			t.Error("IsPrime test failed")
		}
	}
}

// TestIsPalindrome tests IsPalindrome
func TestIsPalindrome(t *testing.T) {
	var tt = []goutils.TTSB{
		{"me", false},
		{"bob", true},
		{"90", false},
		{"909", true},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.IsPalindrome(tt[i].Test)
		testExp := tt[i].Expt

		if testIn != testExp {
			t.Error("IsPalindrome test failed")
		}
	}
}

// TestIsPalindrome tests IsPalindrome
func TestIsEuclidianTriad(t *testing.T) {
	type IsEuclidianTriadTest struct {
		a    int
		b    int
		c    int
		expt bool
	}

	var tt = []IsEuclidianTriadTest{
		{1, 2, 3, true},
		{3, 2, 1, false},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.IsEuclidianTriad(tt[i].a, tt[i].b, tt[i].c)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("IsEuclidianTriad test failed")
			t.Log("example " + strconv.Itoa(i+1))
		}
	}
}
