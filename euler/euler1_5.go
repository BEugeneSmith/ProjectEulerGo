package euler

import (
	"GoProjectEuler/goutils"
	"math"
	"strconv"
	"strings"
)

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

// Euler2 solution
func Euler2(l int) int {
	FibSeq := goutils.FibonacciLimit(l)
	var SumOfEvens int

	for i := 1; i < len(FibSeq); i++ {
		FibElem := FibSeq[i]
		if FibElem < l && goutils.IsDiv(FibElem, 2) {
			SumOfEvens += FibElem
		}
	}
	return SumOfEvens
}

// Euler3 solution
func Euler3(l int) []int {
	allFactors := goutils.GetFactors(l)
	var primeFactors []int

	for i := 0; i < len(allFactors); i++ {
		if goutils.IsPrime(allFactors[i]) {
			primeFactors = append(primeFactors, allFactors[i])
		}
	}

	return primeFactors
}

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

// Euler5 solution
func Euler5(p int) int {
	allProducts := goutils.FilterList(p)
	return allProducts[0]
}
