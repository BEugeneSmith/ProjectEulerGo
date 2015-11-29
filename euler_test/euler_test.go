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
			t.Error("Euler 1 failed")
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
			t.Error("Euler 2 failed")
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
			t.Error("Euler 3 failed")
		}
	}
}

func TestEuler4(t *testing.T) {
	type Euler4Test struct {
		test int    // input
		expt string // expected result
	}

	var tt = []Euler4Test{
		{3, "99,91"},
		{4, "995,583"},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler4(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 4 failed")
		}
	}
}

func TestEuler5(t *testing.T) {
	type Euler5Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler5Test{
		{10, 2520},
		{6, 60},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler5(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 5 failed")
		}
	}
}

func TestEuler6(t *testing.T) {
	type Euler6Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler6Test{
		{10, 2640},
		// {100, 25164150},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler6(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 6 failed")
		}
	}
}

func TestEuler7(t *testing.T) {
	type Euler7Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler7Test{
		{6, 13},
		// {10001, 104743},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler7(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 7 failed")
		}
	}
}
