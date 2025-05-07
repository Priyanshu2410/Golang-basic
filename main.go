package main

import (
	"fmt"
	"mygo/mypck"
)

// main is the entry point of the program
func main() {
    // Print "Hello, World!" to console
    fmt.Println("Hello, World!")
    // Print a separator line
    fmt.Println("--------------")
    // Call Print function from mypck package
    mypck.Print("Hello, World!")

    // Declaring variables using var keyword
    var name string = "John"
    var age int = 25
    var height float64 = 5.9
    var isStudent bool = true

    // Short declaration syntax
    city := "New York"
    temperature := 72.5

    // Multiple variable declaration
    firstName, lastName := "Jane", "Doe"

    // Print all variables
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("City:", city)
    fmt.Println("Temperature:", temperature)
    fmt.Printf("Full Name: %s %s\n", firstName, lastName)

    // Updating variables
    age = 26
    city = "Los Angeles"
    fmt.Println("\nAfter updates:")
    fmt.Println("New Age:", age)
    fmt.Println("New City:", city)

    // Access exported variables from mypck package
    fmt.Println("Exported variables:", mypck.ExportedName, mypck.ExportedAge)
}
