package testgoutils

import (
	"GoProjectEuler/goutils"
	"testing"
)

// TestGetFactors tests the GetFactors
func TestGetFactors(t *testing.T) {
	var tt = []goutils.TTIA{
		{10, []int{2, 5}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.GetFactors(tt[i].Test)
		testExp := tt[i].Expt

		if goutils.ArrayEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}

// TestInitFilteredList tests InitFilteredList
func TestInitFilteredList(t *testing.T) {
	var tt = []goutils.TTIA{
		{6, []int{6, 12, 18, 24, 30, 36, 42, 48, 54, 60}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.InitFilteredList(tt[i].Test)
		testExp := tt[i].Expt

		if goutils.ArrayEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}

// TestInitFilteredList tests InitFilteredList
func TestFilterList(t *testing.T) {
	var tt = []goutils.TTIA{
		{6, []int{60}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.FilterList(tt[i].Test)
		testExp := tt[i].Expt

		if goutils.ArrayEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}
