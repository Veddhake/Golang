package main

// Go does not have a while loop so you use a for loop to make a while loop

import "fmt"

func main() {
	x := 1

	for x <= 20 {
		fmt.Println(x)
		x++
	}

	i := 1

	for i <= 5 {
		fmt.Println("Hello")
		i++
	}
}
