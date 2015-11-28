package main

import (
	"goutils"
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
			t.Error("Test failed")
		}
	}
}

// TestIsPrime tests the IsPrime function
func TestIsPrime(t *testing.T) {
	type IsPrimeTest struct {
		test int  // input
		expt bool // expected result
	}

	var tt = []IsPrimeTest{
		{5, true},
		{4, false},
		{9, false},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.IsPrime(tt[i].test)
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("Test failed")
		}
	}
}
