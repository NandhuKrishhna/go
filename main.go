package main

import "fmt"

// Function that returns the sum of two numbers
func add(a int, b int) int {
	return a + b
}

// Struct example
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello, World!")

	// For loop example
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Calling a function
	sum := add(10, 20)
	fmt.Println("Sum of 10 and 20 is:", sum)

	// Creating and printing a struct
	person := Person{Name: "Nandhu", Age: 22}
	fmt.Printf("Person: %+v\n", person)
}
