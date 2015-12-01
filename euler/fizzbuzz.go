package euler

import (
	"GoProjectEuler/goutils"
	"fmt"
	"strconv"
)

// FizzBuzzPop fun!!!
func FizzBuzzPop() {
	for i := 1; i <= 100; i++ {
		var str string

		if goutils.IsDiv(i, 3) {
			str += "Fizz"
		}
		if goutils.IsDiv(i, 5) {
			str += "Buzz"
		}
		if goutils.IsDiv(i, 7) {
			str += "Pop"
		}
		if str == "" {
			istring := strconv.Itoa(i)
			str += istring
		}
		fmt.Println(str)

	}
}
