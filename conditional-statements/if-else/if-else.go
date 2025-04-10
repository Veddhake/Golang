package main

import "fmt"

func main() {

	password := "12345678"

	if len(password) > 7 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Invalid Password")
	}

	num := 5

	if num > 5 {
		fmt.Println("Number is greater than 5")
	} else if num < 5 {
		fmt.Println("Number is less than 5")
	} else {
		fmt.Println("Number is equal to 5")
	}
}
