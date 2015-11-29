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
