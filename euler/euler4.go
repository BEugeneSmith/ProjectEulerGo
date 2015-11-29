package euler

import (
	"GoProjectEuler/goutils"
	"math"
	"strconv"
	"strings"
)

// Euler4 solution
func Euler4(p int) string {
	p--
	upperL := int(math.Pow(10, float64(p))) - 1

	p--
	lowerL := int(math.Pow(10, float64(p)))

	for j := upperL; j > lowerL; j-- {
		for k := upperL; k > lowerL; k-- {
			prod := strconv.Itoa(j * k)

			if goutils.IsPalindrome(prod) {
				str := []string{strconv.Itoa(j), ",", strconv.Itoa(k)}
				nums := strings.Join(str, "")
				return nums

			}
		}
	}
	return ""
}
