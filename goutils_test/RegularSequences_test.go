package testgoutils

import (
	"GoProjectEuler/goutils"
	"testing"
)

// TestFibonacciNumber test FibonacciNumber
func TestFibonacciNumber(t *testing.T) {
	var tt = []goutils.TTII{
		{3, 2},
		{10, 55},
	}

	for i := 0; i < len(tt); i++ {
		testInIndex := tt[i].Test - 1
		testIn := goutils.FibonacciNumber(tt[i].Test)[testInIndex]
		testExp := tt[i].Expt

		if testExp != testIn {
			t.Error("FibonacciNumber test failed")
		} else {
			t.Log("Passed")
		}
	}
}

// TestFibonacciLimit tests FibonacciLimit
func TestFibonacciLimit(t *testing.T) {
	var tt = []goutils.TTII{
		{10, 8},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.FibonacciLimit(tt[i].Test)
		lastTestIn := testIn[len(testIn)-1]
		testExp := tt[i].Expt

		if testExp != lastTestIn {
			t.Error("FibonacciLimit test failed")
		}
	}
}

// TestSumSquare tests SumSquare
func TestSumSquare(t *testing.T) {
	var tt = []goutils.TTII{
		{10, 385},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.SumSquare(tt[i].Test)
		testExp := tt[i].Expt

		if testExp != testIn {
			t.Error("SumSquare test failed")
		}
	}
}

// TestSquareSum tests SquareSum
func TestSquareSum(t *testing.T) {
	var tt = []goutils.TTII{
		{10, 3025},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.SquareSum(tt[i].Test)
		testExp := tt[i].Expt

		if testExp != testIn {
			t.Error("SquareSum test failed")
		}
	}
}

func TestPrimeNumber(t *testing.T) {
	var tt = []goutils.TTIA{
		{6, []int{2, 3, 5, 7, 11, 13}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.PrimeNumber(tt[i].Test)
		testExp := tt[i].Expt

		if goutils.ArrayEqual(testExp, testIn) == false {
			t.Error("PrimeNumber test failed")
		}
	}
}

func TestPrimeLimit(t *testing.T) {
	var tt = []goutils.TTIA{
		{20, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{10, []int{2, 3, 5, 7}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.PrimeLimit(tt[i].Test)
		testExp := tt[i].Expt

		if goutils.ArrayEqual(testExp, testIn) == false {
			t.Error("PrimeLimit test failed")
		}
	}
}
