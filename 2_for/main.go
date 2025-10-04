package main

import "fmt"

// Demonstrating the "for" loop in Go
// Note: "for" is the only looping construct in Go

func main() {

	fmt.Println("=== Classic for loop ===")
	// Classic for loop: initializer; condition; post
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n=== While-style loop ===")
	// While-style loop: using only a condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n=== Infinite loop (use break to exit) ===")
	// Infinite loop example
	count := 0
	for {
		fmt.Println("Looping infinitely")
		count++
		if count >= 3 {
			break // exit after 3 iterations
		}
	}

	fmt.Println("\n=== Using range in for loop ===")
	// Using range to iterate over slices or arrays
	numbers := []int{10, 20, 30}
	for idx, val := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
	}

}
