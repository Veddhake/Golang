package main

import "fmt"

func main() {
	userData := map[string]int{
		"Ved":   20,
		"John":  29,
		"Harry": 40,
	}

	fmt.Println(userData)

	userData["John"] = 23
	fmt.Println(userData)

	fmt.Println(userData["Ved"])

	for prop, value := range userData {
		fmt.Println(prop, value)
	}

}
