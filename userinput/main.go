package main
import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	// ReadString reads until the first occurrence of the delimiter ('\n' in this case)
	// It returns the string read and an error value
	// Here we're ignoring the error with _ but in production code you should handle it
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, " + name + "!")
}