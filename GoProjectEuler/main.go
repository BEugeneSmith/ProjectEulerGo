package main

import (
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

	euler.RoutesFive(name, arg)

}
