package testgoutils

import (
	"GoProjectEuler/goutils"
	"testing"
)

// TestFibonacciNumber test FibonacciNumber
func TestFibonacciNumber(t *testing.T) {
	type FibonacciTest struct {
		test int // input
		expt int // expected result
	}

	var tt = []FibonacciTest{
		{3, 2},
		{10, 55},
	}

	for i := 0; i < len(tt); i++ {
		testInIndex := tt[i].test - 1
		testIn := goutils.FibonacciNumber(tt[i].test)[testInIndex]
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("FibonacciNumber test failed")
		} else {
			t.Log("Passed")
		}
	}
}

// TestFibonacciLimit tests FibonacciLimit
func TestFibonacciLimit(t *testing.T) {
	type FibonacciTest struct {
		test int // input
		expt int // expected result
	}

	var tt = []FibonacciTest{
		{10, 8},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.FibonacciLimit(tt[i].test)
		lastTestIn := testIn[len(testIn)-1]
		testExp := tt[i].expt

		if testExp != lastTestIn {
			t.Error("FibonacciLimit test failed")
		}
	}
}

// TestSumSquare tests SumSquare
func TestSumSquare(t *testing.T) {
	type SumSquareTest struct {
		test int // input
		expt int // expected result
	}

	var tt = []SumSquareTest{
		{10, 385},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.SumSquare(tt[i].test)
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("SumSquare test failed")
		}
	}
}

// TestSquareSum tests SquareSum
func TestSquareSum(t *testing.T) {
	type SquareSumTest struct {
		test int // input
		expt int // expected result
	}

	var tt = []SquareSumTest{
		{10, 3025},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.SquareSum(tt[i].test)
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("SquareSum test failed")
		}
	}
}
