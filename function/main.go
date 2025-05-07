package main

import "fmt"

// Basic function with no parameters and no return value
func sayHello() {
    fmt.Println("Hello!")
}

// Function with parameters
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Function with multiple parameters
func add(x, y int) {
    fmt.Printf("Sum: %d\n", x+y)
}

// Function with return value
func multiply(x, y int) int {
    return x * y
}

// Function with multiple return values
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return x / y, nil
}

// Function with named return values
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Function with closure (anonymous function)
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// Recursive function
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    // Calling basic function
    sayHello()

    // Calling function with argument
    greet("Alice")

    // Calling function with multiple arguments
    add(5, 3)

    // Using function with return value
    result := multiply(4, 6)
    fmt.Printf("Multiplication result: %d\n", result)

    // Using function with multiple return values
    quotient, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Division result: %.2f\n", quotient)
    }

    // Using function with named return values
    area, perim := rectangle(5.0, 3.0)
    fmt.Printf("Rectangle - Area: %.2f, Perimeter: %.2f\n", area, perim)

    // Using variadic function
    total := sum(1, 2, 3, 4, 5)
    fmt.Printf("Sum of numbers: %d\n", total)

    // Using closure
    increment := counter()
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
    fmt.Println(increment()) // 3

    // Using recursive function
    fact := factorial(5)
    fmt.Printf("Factorial of 5: %d\n", fact)
}
