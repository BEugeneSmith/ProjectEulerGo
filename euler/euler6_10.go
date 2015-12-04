package euler

import (
	"GoProjectEuler/goutils"
	"log"
	"strconv"
	"strings"
)

// Euler6 solution
func Euler6(p int) int {
	sumSquare := goutils.SumSquare(p)
	squareSum := goutils.SquareSum(p)

	squareDiff := squareSum - sumSquare
	return squareDiff
}

// Euler7 solution
func Euler7(p int) int {
	primes := goutils.PrimeNumber(p)

	return primes[len(primes)-1]
}

// Euler8 solution
func Euler8(size int) int {
	num := Euler8string

	// var maxSet []string // for testing
	max := 0
	for i := 0; i < (len(num) - (size - 1)); i++ {
		set := num[i:(i + size)]
		if strings.Contains(set, "0") {
			continue
		}
		splitStr := strings.Split(set, "")
		prod := 1

		for j := 0; j < len(splitStr); j++ {
			conv, err := strconv.Atoi(splitStr[j])
			if err != nil {
				log.Fatal(err)
			}
			prod *= conv
		}
		if prod > max {
			// maxSet = splitStr // for testing
			max = prod
		}
	}
	return max
}

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

// Euler10 solution
func Euler10(n int) int {
	set := goutils.PrimeLimit(n)
	sum := 0
	for i := 0; i < len(set); i++ {
		sum += set[i]
	}
	return sum
}
