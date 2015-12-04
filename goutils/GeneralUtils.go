package goutils

// ArrayEqual checks the equivalency of two arrays
func ArrayEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

// GetProduct returns the product of a int slice
func GetProduct(l []int) int {
	prod := 1
	for i := 0; i < len(l); i++ {
		prod *= l[i]
	}
	return prod
}
