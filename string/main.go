package main

import (
	"fmt"
	"strings"
)

func main() {
	// Basic string declaration
	str := "Hello, Go Strings!"

	// 1. String Length
	fmt.Printf("Length of string: %d\n", len(str))

	// 2. String Conversion
	// Convert to uppercase
	upperStr := strings.ToUpper(str)
	fmt.Printf("Uppercase: %s\n", upperStr)
	// Convert to lowercase
	lowerStr := strings.ToLower(str)
	fmt.Printf("Lowercase: %s\n", lowerStr)

	// 3. String Contains
	// Check if string contains substring
	contains := strings.Contains(str, "Go")
	fmt.Printf("Contains 'Go': %v\n", contains)

	// 4. String Replacement
	// Replace all occurrences
	replaced := strings.Replace(str, "o", "0", -1)
	fmt.Printf("Replaced 'o' with '0': %s\n", replaced)

	// 5. String Splitting
	// Split string by delimiter
	words := strings.Split(str, " ")
	fmt.Printf("Split words: %v\n", words)

	// 6. String Joining
	// Join strings with delimiter
	joined := strings.Join([]string{"Hello", "World"}, "-")
	fmt.Printf("Joined string: %s\n", joined)

	// 7. Trim Functions
	// Remove leading and trailing spaces
	spacedStr := "   Trim me   "
	trimmed := strings.TrimSpace(spacedStr)
	fmt.Printf("Trimmed: '%s'\n", trimmed)

	// 8. String Prefix/Suffix
	// Check if string starts with prefix
	hasPrefix := strings.HasPrefix(str, "Hello")
	fmt.Printf("Starts with 'Hello': %v\n", hasPrefix)
	// Check if string ends with suffix
	hasSuffix := strings.HasSuffix(str, "!")
	fmt.Printf("Ends with '!': %v\n", hasSuffix)

	// 9. String Count
	// Count occurrences of substring
	count := strings.Count(str, "l")
	fmt.Printf("Count of 'l': %d\n", count)

	// 10. String Index
	// Find first occurrence of substring
	index := strings.Index(str, "Go")
	fmt.Printf("Index of 'Go': %d\n", index)

	// 11. String Repeat
	// Repeat string multiple times
	repeated := strings.Repeat("Go!", 3)
	fmt.Printf("Repeated string: %s\n", repeated)

	// 12. String Fields
	// Split string by whitespace
	text := "This is a    sample   text"
	fields := strings.Fields(text)
	fmt.Printf("Fields: %v\n", fields)

	// 13. String Builder
	// Efficient string concatenation
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	result := builder.String()
	fmt.Printf("Built string: %s\n", result)
}
