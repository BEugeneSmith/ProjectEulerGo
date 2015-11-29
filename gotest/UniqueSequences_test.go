package testgoutils

import (
	"GoProjectEuler/goutils"
	"testing"
)

// TestGetFactors tests the GetFactors
func TestGetFactors(t *testing.T) {
	type FactorsTest struct {
		test int   // input
		expt []int // expected result
	}

	var tt = []FactorsTest{
		{10, []int{2, 5}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.GetFactors(tt[i].test)
		testExp := tt[i].expt

		if goutils.ArrayEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}
