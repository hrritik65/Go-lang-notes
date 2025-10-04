package main

import "fmt"

func main() {

	fmt.Println("=== Simple if-else ===")
	age := 21
	if age >= 18 {
		fmt.Println("You are legally an adult.")
	} else {
		fmt.Println("You are still a minor.")
	}

	fmt.Println("\n=== if-else if-else ===")
	age = 14
	if age >= 18 {
		fmt.Println("You are grown up!")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	fmt.Println("\n=== Logical AND in if statement ===")
	role := "editor"
	hasAccess := true
	if role == "editor" && hasAccess {
		fmt.Println("Editor access granted.")
	} else {
		fmt.Println("No access for you.")
	}

	fmt.Println("\n=== Variable declaration inside if ===")
	if tempAge := 17; tempAge >= 18 {
		fmt.Println("Eligible for voting:", tempAge)
	} else if tempAge >= 13 {
		fmt.Println("Not eligible for voting yet:", tempAge)
	}

	fmt.Println("\nNote: Go does not have a ternary operator, so always use if-else for conditional expressions.")

}
