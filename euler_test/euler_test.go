package testeuler

import (
	"GoProjectEuler/euler"
	"GoProjectEuler/goutils"
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

		if goutils.ArrayEqual(testIn, testExp) == false {
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

func TestEuler8(t *testing.T) {
	type Euler8Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler8Test{
		{4, 5832},
		{13, 23514624000},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler8(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 8 failed")
		}
	}
}

func TestEuler9(t *testing.T) {
	type Euler9Test struct {
		test int    // input
		expt string // expected result
	}

	var tt = []Euler9Test{
		{12, "3,4,5"},
		{100, "32,33,35"},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler9(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 9 failed")
		}
	}
}
