package main

import "fmt"

func main() {
	var numbers [5]int // by default all values in array are 0 for int array

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)

	var arr1 = [3]int{1, 2, 3}

	fmt.Println(arr1)

	var names = [3]string{"Ved", "John", "Jack"}

	fmt.Println(names)

	fmt.Println(len(names))

	fmt.Println(cap(names)) // capacity of array

	for item := 0; item < len(arr1); item++ {
		fmt.Println(arr1[item])
	}

}
