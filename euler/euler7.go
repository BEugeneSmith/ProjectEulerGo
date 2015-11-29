package euler

import "GoProjectEuler/goutils"

// Euler7 solution
func Euler7(p int) int {
	primes := goutils.PrimeNumber(p)

	return primes[len(primes)-1]
}
