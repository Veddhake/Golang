package main

// slices are dynamically sized arrays

import (
	"fmt"
)

func main() {
	num := []int{10, 20, 30, 40, 50}
	fmt.Println(num[:])
	fmt.Println(num[1:4]) // [a:b] , a is inclusive but b is exclusive
	fmt.Println(num[3:])

	ans := []int{10, 20}
	fmt.Println(ans)
	ans = append(ans, 30, 40, 50)
	fmt.Println(ans)

	num1 := make([]int, 5, 20) // this take type, length , capacity
	fmt.Println(num1)

	num1[0] = 10
	fmt.Println(num1)

	num1[1] = 20
	num1[2] = 30
	num1[3] = 40
	num1[4] = 50

	fmt.Println(num1)
	fmt.Println(cap(num1))

	ppls := []string{"Ved", "John", "Doe"}

	for i := 0; i < len(ppls); i++ {
		fmt.Println(ppls[i])
	}

	for index, value := range ppls {
		fmt.Println(index, value)
	}
}
