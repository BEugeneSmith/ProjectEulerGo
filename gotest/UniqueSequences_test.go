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

// TestInitFilteredList tests InitFilteredList
func TestInitFilteredList(t *testing.T) {
	type InitFilteredListTest struct {
		test int   // input
		expt []int // expected result
	}

	var tt = []InitFilteredListTest{
		{6, []int{6, 12, 18, 24, 30, 36, 42, 48, 54, 60}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.InitFilteredList(tt[i].test)
		testExp := tt[i].expt

		if goutils.ArrayEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}

// TestInitFilteredList tests InitFilteredList
func TestFilterList(t *testing.T) {
	type FilterListTest struct {
		test int   // input
		expt []int // expected result
	}

	var tt = []FilterListTest{
		{6, []int{60}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.FilterList(tt[i].test)
		testExp := tt[i].expt

		if goutils.ArrayEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}
