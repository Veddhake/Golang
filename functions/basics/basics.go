package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello from a function")
}

func printName(username string) {
	fmt.Println(username)
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func showNumbers(numbers ...int) { // this is used to enter unlimited number of numbers, you can use it for any data types

	fmt.Println(numbers)
}

func showStuff(word string, numbers ...int) { // the unlimited parameter must be at the end of function parameters
	fmt.Println(word, numbers)
}

func main() {

	hello()

	printName("Ved")

	greet("Ved")

	showNumbers(1, 2, 3, 4, 5)

	showStuff("Ved", 1, 2, 3, 4, 5)
}
