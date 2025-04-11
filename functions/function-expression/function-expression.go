package main

import "fmt"

func main() {
	add := func(num1 int, num2 int) int { // these are functions stored in variables
		return num1 + num2
	}

	res := add(10, 20)
	fmt.Println(res)
}
