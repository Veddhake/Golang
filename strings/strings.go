package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Ved Dhake"
	fmt.Println(name)

	intro := `Hello 	
						my name is
						Ved `

	fmt.Println(intro)

	greetings := "Hello" + " my name is" + " Ved"

	fmt.Println(greetings)

	fmt.Printf("%c", name[0]) // Println makes use of unicode characters hence Printf is required for formatting. %c is used for formatting it as a characters

	fmt.Println()

	fmt.Println(len(name))

	msg1 := "one"
	msg2 := "one"
	msg3 := "two"

	fmt.Println(strings.Compare(msg1, msg2))
	fmt.Println(strings.Compare(msg2, msg3))

	findChar := "Golang Programming Language"
	fmt.Println(strings.Contains(findChar, "Golang"))
	fmt.Println(strings.Contains(findChar, "Hey"))

	stringOne := "Hello"
	fmt.Println(strings.ToLower(stringOne))
	fmt.Println(strings.ToUpper(stringOne))

}
