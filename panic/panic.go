package main

import "fmt"

func employee(name *string, age int) {
	if age > 65 {
		panic("Age cannot be greater than retirement age")
	}
	fmt.Println("Employee:", *name, "Age:", age)
}

func main() {
	empname := "Ved"
	age := 70
	employee(&empname, age)
}
