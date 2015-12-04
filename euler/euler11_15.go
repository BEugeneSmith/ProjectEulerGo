package euler

import (
	"GoProjectEuler/goutils"
	"fmt"
)

// Euler11 exports the the solution to Project Euler no.11
func Euler11() {
	type MaxProd struct {
		Init int
		Vals []int
	}

	var dirs = make(map[string]*MaxProd)

	// maxoftypes returns the maximum combination of
	// product and product components
	dirs["diag"] = &MaxProd{1, []int{}}
	dirs["right"] = &MaxProd{1, []int{}}
	dirs["left"] = &MaxProd{1, []int{}}
	dirs["up"] = &MaxProd{1, []int{}}
	dirs["down"] = &MaxProd{1, []int{}}

	matrix := Euler11Matrix

	// traverse the matrix diagonally
	for y := 0; y < len(matrix)-3; y++ {
		for x := 0; x < len(matrix)-3; x++ {
			numset := []int{}
			for n := 0; n < 4; n++ {
				num := matrix[y+n][x+n]
				numset = append(numset, num)
			}
			setprod := goutils.GetProduct(numset)
			if setprod > dirs["diag"].Init {
				dirs["diag"].Init = setprod
				dirs["diag"].Vals = numset
			}
		}
	}

	// traverse the matrix in slices of len(4)
	// down from the point at [x,y]
	for y := 0; y < len(matrix)-3; y++ {
		for x := 0; x < len(matrix); x++ {
			numset := []int{}
			for n := 0; n < 4; n++ {
				num := matrix[y+n][x]
				numset = append(numset, num)
			}
			setprod := goutils.GetProduct(numset)
			if setprod > dirs["down"].Init {
				dirs["down"].Init = setprod
				dirs["down"].Vals = numset
			}
		}
	}

	// traverse the matrix in slices of len(4)
	// up from the point at [x,y]
	for y := 3; y < len(matrix); y++ {
		for x := 0; x < len(matrix); x++ {
			numset := []int{}
			for n := 0; n < 4; n++ {
				num := matrix[y-n][x]
				numset = append(numset, num)
			}
			setprod := goutils.GetProduct(numset)
			if setprod > dirs["up"].Init {
				dirs["up"].Init = setprod
				dirs["up"].Vals = numset
			}
		}
	}

	// traverse the matrix in slices of len(4)
	// right from the point at [x,y]
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix)-3; x++ {
			numset := []int{}
			for n := 0; n < 4; n++ {
				num := matrix[y][x+n]
				numset = append(numset, num)
			}
			setprod := goutils.GetProduct(numset)
			if setprod > dirs["right"].Init {
				dirs["right"].Init = setprod
				dirs["right"].Vals = numset
			}
		}
	}

	// traverse the matrix in slices of len(4)
	// left from the point at [x,y]
	for y := 0; y < len(matrix); y++ {
		for x := 3; x < len(matrix); x++ {
			numset := []int{}
			for n := 0; n < 4; n++ {
				num := matrix[y][x-n]
				numset = append(numset, num)
			}
			setprod := goutils.GetProduct(numset)
			if setprod > dirs["left"].Init {
				dirs["left"].Init = setprod
				dirs["left"].Vals = numset
			}
		}
	}

	for key := range dirs {
		fmt.Println(dirs[key])
	}
}
