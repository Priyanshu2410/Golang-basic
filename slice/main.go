package main

import "fmt"

func main() {
	// Declaring and initializing a slice
	// Syntax: []datatype{values}
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	// Creating a slice using make()
	// make([]datatype, length, capacity)
	slice1 := make([]int, 3, 5)
	fmt.Println("Slice created with make():", slice1)

	// Slicing an existing slice
	// Syntax: slice[startIndex:endIndex]
	// endIndex is exclusive
	slice2 := numbers[1:4]
	fmt.Println("Sliced portion:", slice2)

	// Appending elements to a slice
	numbers = append(numbers, 6, 7)
	fmt.Println("After append:", numbers)

	// Length and Capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(numbers), cap(numbers))

	// Creating slice from array
	array := [5]int{10, 20, 30, 40, 50}
	sliceFromArray := array[1:4]
	fmt.Println("Slice from array:", sliceFromArray)

	// Modifying slice (affects underlying array)
	sliceFromArray[0] = 25
	fmt.Println("Modified array:", array)

	// Multi-dimensional slices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D slice:", matrix)

	// Copy slices
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Copied slice:", dst)

	// Slice with zero value (nil)
	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice == nil)

	// Check if slice is empty
	emptySlice := []int{}
	fmt.Println("Empty slice length:", len(emptySlice))
}
