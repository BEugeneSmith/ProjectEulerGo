package euler

import (
	"GoProjectEuler/goutils"
	"fmt"
)

// RoutesFive returns a function
func RoutesFive(s string, i int) {
	switch s {
	case "euler1":
		fmt.Println(Euler1(i))
	}
}

// Euler1 solution
func Euler1(l int) int {
	sum := 0
	for i := 1; i < l; i++ {
		if goutils.IsDiv(i, 3) || goutils.IsDiv(i, 5) {
			sum += i
		}
	}
	return sum
}

func euler2() {

}
