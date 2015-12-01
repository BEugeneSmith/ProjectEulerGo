package testeuler

import (
	"GoProjectEuler/euler"
	"testing"
)

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

func TestEuler10(t *testing.T) {
	type Euler10Test struct {
		test int // input
		expt int // expected result
	}

	var tt = []Euler10Test{
		{10, 17},
		{200, 4227},
	}

	for i := 0; i < len(tt); i++ {
		testIn := euler.Euler10(tt[i].test)
		testExp := tt[i].expt

		if testIn != testExp {
			t.Error("Euler 10 failed")
		}
	}
}
