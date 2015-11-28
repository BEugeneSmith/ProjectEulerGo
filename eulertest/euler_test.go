package TestEuler

import "testing"

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
		testIn := euler.euler1(tt[i].test)
		testExp := tt[i].expt

		if testExp != testIn {
			t.Error("Test failed")
		}
	}
}
