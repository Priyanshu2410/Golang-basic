package main

import "fmt"

// Define a basic struct for a Person
type Person struct {
	Name    string
	Age     int
	Address string
}

// Define a nested struct for an Employee that includes Person
type Employee struct {
	Person    // Embedded struct (inheritance-like behavior)
	CompanyID string
	Salary    float64
}

// Method associated with Person struct
func (p Person) Greet() string {
	// Sprintf formats and returns a string using the provided format string and arguments
	// The formatted string introduces the person with their name and age
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

// Method to update salary for Employee
func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

func main() {
	// Creating a basic struct instance
	person1 := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	// Another way to create struct instance
	var person2 Person
	person2.Name = "Jane Smith"
	person2.Age = 25
	person2.Address = "456 Oak Ave"

	// Creating an Employee struct instance
	employee1 := Employee{
		Person: Person{
			Name:    "Bob Wilson",
			Age:     35,
			Address: "789 Pine Rd",
		},
		CompanyID: "EMP123",
		Salary:    50000,
	}

	// Using struct methods
	fmt.Println(person1.Greet())
	
	// Updating employee salary using pointer method
	employee1.UpdateSalary(55000)
	
	// Anonymous struct - useful for one-time use
	tempPerson := struct {
		Name  string
		Email string
	}{
		Name:  "Temporary User",
		Email: "temp@example.com",
	}

	// Printing various struct instances
	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 2: %+v\n", person2)
	fmt.Printf("Employee 1: %+v\n", employee1)
	fmt.Printf("Temporary Person: %+v\n", tempPerson)

	// Demonstrating struct comparison
	person3 := Person{Name: "John Doe", Age: 30, Address: "123 Main St"}
	fmt.Printf("person1 equals person3: %v\n", person1 == person3)
}
