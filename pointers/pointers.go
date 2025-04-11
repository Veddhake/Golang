package main

import "fmt"

func main() {
	address := 10

	fmt.Println(&address)

	x := 20

	var pointerToX *int = &x

	fmt.Println("Memory address of x:", pointerToX)
	fmt.Println("Value pointed by pointer:", *pointerToX)
}
