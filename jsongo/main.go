package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// User represents a basic user structure
type User struct {
	ID       int    `json:"id"`   // json tag for field mapping
	Name     string `json:"name"` // custom json field name
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`     // snake_case in JSON
	Age      *int   `json:"age,omitempty"` // omitempty skips null/zero values
}

// Product demonstrates nested structs and arrays
type Product struct {
	ID       int                    `json:"id"`
	Name     string                 `json:"name"`
	Price    float64                `json:"price"`
	Tags     []string               `json:"tags,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func main() {
	// Marshal (Struct to JSON)
	age := 25
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		IsActive: true,
		Age:      &age,
	}

	// Convert struct to JSON bytes
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("Marshaled JSON: %s\n", jsonBytes)

	// Unmarshal (JSON to Struct)
	jsonStr := `{
		"id": 2,
		"name": "Jane Smith",
		"email": "jane@example.com",
		"is_active": true
	}`

	var newUser User
	err = json.Unmarshal([]byte(jsonStr), &newUser)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	fmt.Printf("Unmarshaled struct: %+v\n", newUser)

	// Working with nested structures
	product := Product{
		ID:    1,
		Name:  "Laptop",
		Price: 999.99,
		Tags:  []string{"electronics", "computers"},
		Metadata: map[string]interface{}{
			"manufacturer": "TechCo",
			"model":        "X100",
			"inStock":      true,
		},
	}

	// Pretty printing JSON
	prettyJSON, err := json.MarshalIndent(product, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	fmt.Printf("\nPretty printed JSON:\n%s\n", prettyJSON)

	// Decoding JSON stream
	decoder := json.NewDecoder(strings.NewReader(jsonStr)) // Add "strings" package to imports
	var decodedUser User
	err = decoder.Decode(&decodedUser)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		return
	}

	// Encoding JSON stream
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(user)
	if err != nil {
		fmt.Printf("Error encoding: %v\n", err)
		return
	}
}
