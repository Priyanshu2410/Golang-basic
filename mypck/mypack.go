package mypck
import "fmt"

// ExportedName is an exported variable that can be accessed from other packages
var ExportedName string = "John Doe"

// ExportedAge is an exported variable that can be accessed from other packages
var ExportedAge int = 30

func Print(s string) {
    fmt.Println(s)
    fmt.Println(ExportedName)
    fmt.Println(ExportedAge)
}