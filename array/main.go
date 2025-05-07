package main

import "fmt"

func main() {
    // Arrays in Go are fixed-length sequences of elements of the same type
    // Declaring an array with a specific size and type
    var numbers [5]int // Creates an array of 5 integers, initialized with zero values

    // Array initialization with values
    fruits := [3]string{"apple", "banana", "orange"} // Array literal with size 3

    // Array with implicit size using ...
    colors := [...]string{"red", "green", "blue"} // Compiler counts the elements

    // Accessing and modifying array elements
    numbers[0] = 1 // Arrays are zero-indexed
    numbers[1] = 2
    numbers[2] = 3
    numbers[3] = 4
    numbers[4] = 5

    // Iterating over array using traditional for loop
    fmt.Println("Traditional for loop:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
    }

    // Iterating using range
    fmt.Println("\nRange-based for loop:")
    for index, value := range fruits {
        fmt.Printf("fruits[%d] = %s\n", index, value)
    }

    // Multi-dimensional arrays
    // Creating a 2x3 array
    matrix := [2][3]int{
        {1, 2, 3},   // First row
        {4, 5, 6},   // Second row
    }

    // Accessing elements in multi-dimensional array
    fmt.Println("\nMatrix elements:")
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[0]); j++ {
            fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
        }
    }

    // Array length is part of its type
    fmt.Printf("\nArray lengths:\n")
    fmt.Printf("numbers length: %d\n", len(numbers))
    fmt.Printf("fruits length: %d\n", len(fruits))
    fmt.Printf("colors length: %d\n", len(colors))

    // Copying arrays
    // Arrays are value types in Go, so assigning creates a copy
    numbersCopy := numbers
    numbersCopy[0] = 100 // This won't affect the original array
    
    fmt.Println("\nOriginal array:", numbers)
    fmt.Println("Copied array:", numbersCopy)

    // Array comparison
    // Arrays of the same type can be compared using == operator
    arr1 := [3]int{1, 2, 3}
    arr2 := [3]int{1, 2, 3}
    arr3 := [3]int{3, 2, 1}

    fmt.Printf("\nArray comparison:\n")
    fmt.Printf("arr1 == arr2: %v\n", arr1 == arr2) // true
    fmt.Printf("arr1 == arr3: %v\n", arr1 == arr3) // false
}
