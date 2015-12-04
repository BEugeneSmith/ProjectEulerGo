package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"GoProjectEuler/euler"
)

var num string
var arg int

func init() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}

	flag.StringVar(&num, "num", os.Args[1], "number")

	i, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	flag.IntVar(&arg, "arg", i, "argument")
}

func main() {

	Routes(num, arg)

}

// Routes returns a function
func Routes(s string, i int) {
	switch s {
	case "1":
		fmt.Println(euler.Euler1(i))
	case "2":
		fmt.Println(euler.Euler2(i))
	case "3":
		fmt.Println(euler.Euler3(i))
	case "4":
		fmt.Println(euler.Euler4(i))
	case "5":
		fmt.Println(euler.Euler5(i))
	case "6":
		fmt.Println(euler.Euler6(i))
	case "7":
		fmt.Println(euler.Euler7(i))
	case "8":
		fmt.Println(euler.Euler8(i))
	case "9":
		fmt.Println(euler.Euler9(i))
	case "10":
		fmt.Println(euler.Euler10(i))
	case "11":
		euler.Euler11()
	default:
		euler.FizzBuzzPop()
	}
}
