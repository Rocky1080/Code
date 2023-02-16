package main

import "fmt"

func main() {
	fmt.Println("Welcome to the topic on pointers")

	// var ptr *int

	// fmt.Println("The value of the pointer", ptr)
	number := 25

	var ptr = &number

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)
}
