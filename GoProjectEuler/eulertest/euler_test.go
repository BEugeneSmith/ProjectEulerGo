package testeuler

import (
	"GoProjectEuler/euler"
	"reflect"
	"testing"
)

// TestEuler1 tests Project euler solution 1
func TestEuler1(t *testing.T) {
	type Euler1Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler1Test{
		{10, 23},
		{1000, 233168},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler1(tt[i].test)
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("Test failed")
		}
	}
}

func TestEuler2(t *testing.T) {
	type Euler2Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler2Test{
		{10, 10},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler2(tt[i].test)
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("Test failed")
		}
	}
}

func TestEuler3(t *testing.T) {
	type Euler3Test struct {
		test int   // input
		expt []int // expected result
	}

	var tt = []Euler3Test{
		{20, []int{2, 5}},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler3(tt[i].test)
		testExp := tt[i].expt

		if reflect.DeepEqual(testExp, testIn) {
			t.SkipNow()
		} else {
			t.Error("Test failed")
		}
	}
}
