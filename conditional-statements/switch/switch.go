package main

import "fmt"

func main() {
	day := 8

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	num := 500

	switch {
	case (num > 5):
		fmt.Println("num is greater than 5")
	case (num < 5):
		fmt.Println("num is less than 5")
	case (num == 5):
		fmt.Println("num is equal to 5")
	default:
		fmt.Println("Invalid")
	}
}
