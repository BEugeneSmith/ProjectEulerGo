package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"GoProjectEuler/euler"
)

func main() {

	name := os.Args[1]
	arg, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	RoutesFive(name, arg)

}

// RoutesFive returns a function
func RoutesFive(s string, i int) {
	switch s {
	case "euler1":
		fmt.Println(euler.Euler1(i))
	case "euler2":
		fmt.Println(euler.Euler2(i))
	case "euler3":
		fmt.Println(euler.Euler3(i))
	case "euler4":
		fmt.Println(euler.Euler4(i))
	case "euler5":
		fmt.Println(euler.Euler5(i))
	case "euler6":
		fmt.Println(euler.Euler6(i))
	case "euler7":
		fmt.Println(euler.Euler7(i))
	}
}
