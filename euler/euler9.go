package euler

import (
	"GoProjectEuler/goutils"
	"strconv"
	"strings"
)

// Euler9 solution
func Euler9(n int) string {
	var allstrings []string

	for i := 1; i < n; i++ {
		for j := 2; j < n; j++ {
			for k := 3; k < n; k++ {
				if (i + j + k) == n {
					if goutils.IsEuclidianTriad(i, j, k) {

						var strSlice = []string{strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k)}
						str := strings.Join(strSlice, ",")
						allstrings = append(allstrings, str)
					}
				}
			}
		}
	}
	return allstrings[len(allstrings)-1]
}
