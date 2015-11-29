package euler

import "GoProjectEuler/goutils"

// Euler5 solution
func Euler5(p int) int {
	allProducts := goutils.FilterList(p)
	return allProducts[0]
}
