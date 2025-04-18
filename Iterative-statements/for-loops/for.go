package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("O")
		}
		fmt.Println()
	}
}
