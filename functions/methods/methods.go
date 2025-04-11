package main

import "fmt"

type Person struct {
	gender     string
	firstname  string
	lastname   string
	age        int
	profession string
	isMarried  bool
}

func (p Person) printInfo() {
	fmt.Println("I'm", p.firstname, p.lastname, "Gender:", p.gender, "Age:", p.age, "Profession:", p.profession, "Married:", p.isMarried)
}

func main() {
	ved := Person{
		gender:     "male",
		firstname:  "Ved",
		lastname:   "Dhake",
		age:        20,
		profession: "Programmer",
		isMarried:  false,
	}

	ved.printInfo()
}
