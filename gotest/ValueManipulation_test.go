package testgoutils

import (
	"GoProjectEuler/goutils"
	"testing"
)

// TestReverseString tests ReverseString
func TestReverseString(t *testing.T) {
	type ReverseStringTest struct {
		test string // input
		expt string // expected result
	}

	var tt = []ReverseStringTest{
		{"nai", "ian"},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.ReverseString(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("ReverseString Failed")
		}
	}
}

// TestReverseIntArray tests ReverseString
func TestReverseIntArray(t *testing.T) {
	type ReverseIntArrayTest struct {
		test []int // input
		expt []int // expected result
	}

	var tt = []ReverseIntArrayTest{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.ReverseIntArray(tt[i].test)
		testExp := tt[i].expt

		if goutils.ArrayEqual(testExp, testIn) == false {

			t.Error("ReverseIntArray Failed")
		}
	}
}
