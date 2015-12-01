package goutils

// TTII is a test table int to int
type TTII struct {
	Test int
	Expt int
}

// TTIS is a test table, int to string
type TTIS struct {
	Test int
	Expt string
}

// TTIA is a test table, int to array
type TTIA struct {
	Test int
	Expt []int
}

// TTIB is a test table, int to bool
type TTIB struct {
	Test int
	Expt bool
}

// TTSB is a test table, string to boolean
type TTSB struct {
	Test string
	Expt bool
}
