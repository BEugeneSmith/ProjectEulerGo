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
	num := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

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
