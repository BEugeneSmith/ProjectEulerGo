package testgoutils

import (
	"GoProjectEuler/goutils"
	"testing"
)

// TestArrayEqual tests ArrayEqual
func TestArrayEqual(t *testing.T) {

	var l1 = []int{1, 2, 3}
	var l2 = []int{1, 2, 3}

	equal := goutils.ArrayEqual(l1, l2)

	if equal != true {
		t.Error("ArrayEqual Failed")
	}
}

// TestGetProduct tests GetProduct
func TestGetProduct(t *testing.T) {
	var tt = []goutils.TTAI{
		{[]int{1, 2, 3}, 6},
		{[]int{4, 5, 6}, 120},
	}

	for i := 0; i < len(tt); i++ {
		testIn := goutils.GetProduct(tt[i].Test)
		testExp := tt[i].Expt

		if testExp != testIn {
			t.Error("GetProduct test failed")
		} else {
			t.Log("Passed")
		}
	}
}
