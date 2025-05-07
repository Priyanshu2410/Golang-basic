package main

import "fmt"

func main() {
    // 1. Basic for loop with initialization, condition, and increment
    fmt.Println("1. Standard for loop:")
    for i := 0; i < 3; i++ {
        fmt.Println("Iteration:", i)
    }

    // 2. For loop as while loop
    fmt.Println("\n2. For loop as while loop:")
    counter := 0
    for counter < 3 {
        fmt.Println("Counter is:", counter)
        counter++
    }

    // 3. Infinite loop with break
    fmt.Println("\n3. Infinite loop with break:")
    count := 0
    for {
        if count >= 3 {
            break
        }
        fmt.Println("Count is:", count)
        count++
    }

    // 4. For loop with continue statement
    fmt.Println("\n4. For loop with continue:")
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println("Odd number:", i)
    }

    // 5. For loop for iterating over slices
    fmt.Println("\n5. Range loop for slice:")
    fruits := []string{"apple", "banana", "orange"}
    for index, value := range fruits {
        fmt.Printf("Index: %d, Value: %s\n", index, value)
    }

    // 6. For loop for iterating over maps
    fmt.Println("\n6. Range loop for map:")
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }
    for key, value := range colors {
        fmt.Printf("Key: %s, Value: %s\n", key, value)
    }

    // 7. Nested for loops
    fmt.Println("\n7. Nested for loops:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 2; j++ {
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
}
