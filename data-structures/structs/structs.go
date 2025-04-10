package main

import "fmt"

type Employee struct {
	name   string
	age    int
	job    string
	salary int
}

type Person struct {
	firstName string
	lastName  string
}

type Programmer struct { // nested structs
	Person
	isProgrammer bool
}

func main() {
	var userOne Employee

	userOne.name = "Ved"
	userOne.age = 20
	userOne.job = "SWE"
	userOne.salary = 100000

	fmt.Println(userOne)

	v := Person{
		firstName: "Ved",
		lastName:  "Dhake",
	}

	fmt.Println(v)

	h := Programmer{
		Person: Person{
			firstName: "John",
			lastName:  "Doe",
		},
		isProgrammer: true,
	}
	fmt.Println(h.firstName, h.lastName, h.isProgrammer)

	x := struct { // anonymous structs
		name    string
		address string
	}{
		name:    "Ved",
		address: "ABCDSIIS",
	}

	fmt.Println(x)
}
