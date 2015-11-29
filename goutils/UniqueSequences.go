package goutils

// GetFactors returns a list of a number's factors
func GetFactors(n int) []int {
	var factors []int
	if IsDiv(n, 2) {
		factors = append(factors, 2)
		for i := 2; i <= (n / 2); i++ {
			if IsDiv(n, i) {
				factors = append(factors, i)
			}
		}
		factors = factors[1:]

	} else {
		for i := 3; i <= (n / 2); i += 2 {
			if IsDiv(n, i) {
				factors = append(factors, i)
			}
		}
	}

	return factors
}

//InitFilteredList produces a list to start the number sieve
func InitFilteredList(lf int) []int {
	var max float64 = 1
	for i := 1; i < lf; i++ {
		max *= float64(i)
	}

	list := []int{lf}
	maxLim := int(max / 2)

	for i := lf; i < maxLim; i += lf {
		list = append(list, (lf + list[len(list)-1]))
	}
	return list
}

// FilterList filters values from a list
func FilterList(lf int) []int {

	fullList := InitFilteredList(lf)
	var newList []int

	for i := lf; i > 1; i-- {
		for j := 0; j < len(fullList); j++ {
			k := fullList[j]
			if IsDiv(k, i) {
				newList = append(newList, k)
			}
		}
		fullList = newList
		newList = newList[:0]
	}
	return fullList
}
