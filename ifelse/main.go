package main

import "fmt"

func main() {
    // Basic if statement
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult")
    }

    // If-else statement
    score := 75
    if score >= 90 {
        fmt.Println("Grade: A")
    } else {
        fmt.Println("Grade: Below A")
    }

    // If-else if-else chain
    temperature := 28
    if temperature < 0 {
        fmt.Println("Freezing!")
    } else if temperature < 20 {
        fmt.Println("Cool weather")
    } else if temperature < 30 {
        fmt.Println("Warm weather")
    } else {
        fmt.Println("Hot!")
    }

    // If with initialization statement
    // The variable 'count' is only accessible within the if block
    if count := 5; count > 0 {
        fmt.Println("Count is positive")
    }

    // Nested if statements
    num := 15
    if num > 0 {
        if num%2 == 0 {
            fmt.Println("Positive even number")
        } else {
            fmt.Println("Positive odd number")
        }
    }

    // Using logical operators in conditions
    username := "admin"
    password := "secret123"
    if username == "admin" && password == "secret123" {
        fmt.Println("Login successful")
    }

    // Using OR operator
    holiday := "Sunday"
    if holiday == "Saturday" || holiday == "Sunday" {
        fmt.Println("It's weekend!")
    }

    // Error handling pattern in Go
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}

// Helper function to demonstrate error handling
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
