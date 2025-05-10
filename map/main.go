package main

import "fmt"

func main() {
	// Declaring and initializing a map
	// Syntax: map[KeyType]ValueType{key:value}
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Creating an empty map using make
	scores := make(map[string]int)

	// Adding key-value pairs
	scores["John"] = 89
	scores["Alice"] = 95

	// Accessing values
	fmt.Println("Red hex code:", colors["red"])

	// Checking if key exists
	value, exists := colors["yellow"]
	if !exists {
		fmt.Println("Yellow color not found")
	}

	// Deleting a key-value pair
	delete(colors, "blue")

	// Iterating over a map
	for key, value := range colors {
		fmt.Printf("Color: %s, Hex: %s\n", key, value)
	}

	// Map of custom types
	type Person struct {
		age  int
		city string
	}

	var people = map[string]Person{
		"Bob":   {age: 25, city: "New York"},
		"Alice": {age: 30, city: "London"},
	}

	// Nested maps
	cityPopulation := map[string]map[string]int{
		"USA": {
			"New York": 8400000,
			"Chicago":  2700000,
		},
		"India": {
			"Mumbai": 12400000,
			"Delhi":  11000000,
		},
	}

	// Length of map
	fmt.Printf("Number of colors: %d\n", len(colors))

	// Clear all elements from map (Go 1.21+)
	clear(scores)

	// Common map operations
	fmt.Println("\nCity Populations:")
	for country, cities := range cityPopulation {
		fmt.Printf("%s:\n", country)
		for city, population := range cities {
			fmt.Printf("  %s: %d\n", city, population)
		}
	}

	// Map as a set
	uniqueNumbers := map[int]bool{
		1: true,
		2: true,
		3: true,
	}

	// Check if number exists in set
	if uniqueNumbers[2] {
		fmt.Println("2 exists in the set")
	}
}
