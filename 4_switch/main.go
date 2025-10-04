package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("=== Simple switch example ===")
	number := 2
	switch number {
	case 1:
		fmt.Println("Number is One")
	case 2:
		fmt.Println("Number is Two")
	case 3:
		fmt.Println("Number is Three")
	default:
		fmt.Println("Number is something else")
	}

	fmt.Println("\n=== Multiple values in a single case ===")
	day := time.Now().Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Yay! It's the weekend!")
	default:
		fmt.Println("It's a weekday. Time to work!")
	}

	fmt.Println("\n=== Type switch example ===")
	describeType := func(i interface{}) {
		switch v := i.(type) { // type assertion
		case int:
			fmt.Println(v, "is an integer")
		case string:
			fmt.Println(v, "is a string")
		case bool:
			fmt.Println(v, "is a boolean")
		default:
			fmt.Println(v, "is of some other type")
		}
	}

	describeType(100)
	describeType("GoLang")
	describeType(true)
	describeType(3.14)

	fmt.Println("\nSwitch in Go is powerful and can handle values, multiple conditions, and even types!")

}
