package main

import (
	"fmt"
)

var number = 20 // declaration using var

func main() {

	fmt.Println(number)

	num := 1000 // declaration using :=
	fmt.Println(num)

	one, two, three := 1, 2, 3

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)

}
