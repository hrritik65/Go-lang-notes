package main

import "fmt"

func main() {
	// ===========================
	// 1️ Basic Values
	// ===========================
	fmt.Println("=== Basic Values ===")

	// Integers
	fmt.Println("Integer:", 1+1)

	// Strings
	fmt.Println("String:", "hello golang")

	// Booleans
	fmt.Println("Boolean:", true, false)

	// Floats
	fmt.Println("Float:", 10.5)
	fmt.Println("Float division:", 7.0/3.0)

	fmt.Println()

	// ===========================
	// 2️ Variables
	// ===========================
	fmt.Println("=== Variables ===")

	// Explicit type declaration
	var name string = "golang"
	fmt.Println("Name (explicit):", name)

	// Type inference
	var isAdult = true
	fmt.Println("IsAdult (inferred):", isAdult)

	// Shorthand declaration (only inside functions)
	age := 30
	fmt.Println("Age (shorthand):", age)

	// Reassigning variable
	var price float32 = 50.5
	fmt.Println("Price (before):", price)
	price = 60.0
	fmt.Println("Price (after):", price)

	fmt.Println()

	// ===========================
	// 3️ Constants
	// ===========================
	fmt.Println("=== Constants ===")

	// Single constant
	const appName = "GoLang Tutorial"
	fmt.Println("App Name:", appName)

	// Multiple constants
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println("Server running at", host, "on port", port)
}
