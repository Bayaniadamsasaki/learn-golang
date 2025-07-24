// Structs: tipe data kustom untuk mengelompokkan data terkait
package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person
	ID       int
	Position string
	Salary   float64
}

func (p Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func (e Employee) GetInfo() string {
	return fmt.Sprintf("Employee ID: %d, Position: %s", e.ID, e.Position)
}

func main() {
	fmt.Println("=== BASIC STRUCT ===")
	var person1 Person
	person1.Name = "John Doe"
	person1.Age = 30
	person1.Email = "john@example.com"
	
	person2 := Person{
		Name:  "Jane Smith",
		Age:   25,
		Email: "jane@example.com",
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
	}
	
	person3 := Person{"Bob Johnson", 35, "bob@example.com", Address{}}
	
	fmt.Printf("Person1: %+v\n", person1)
	fmt.Printf("Person2: %+v\n", person2)
	fmt.Printf("Person3: %+v\n", person3)
	
	fmt.Println("\n=== ACCESSING FIELDS ===")
	fmt.Printf("Name: %s, Email: %s\n", person2.Name, person2.Email)
	fmt.Printf("Address: %s, %s, %s\n", person2.Address.Street, person2.Address.City, person2.Address.Country)
	
	fmt.Println("\n=== STRUCT METHODS ===")
	person2.Greet()
	fmt.Printf("Age before birthday: %d\n", person2.Age)
	person2.HaveBirthday()
	fmt.Printf("Age after birthday: %d\n", person2.Age)
	
	fmt.Println("\n=== EMBEDDED STRUCT ===")
	emp := Employee{
		Person: Person{
			Name:  "Alice Brown",
			Age:   28,
			Email: "alice@company.com",
		},
		ID:       101,
		Position: "Software Engineer",
		Salary:   75000.0,
	}
	
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Println(emp.GetInfo())
	emp.Greet()
	fmt.Printf("Accessing embedded field directly: %s\n", emp.Name)
	
	fmt.Println("\n=== ANONYMOUS STRUCT ===")
	point := struct {
		X, Y int
		Name string
	}{
		X:    10,
		Y:    20,
		Name: "Point A",
	}
	fmt.Printf("Point: %+v\n", point)
	
	fmt.Println("\n=== POINTER TO STRUCT ===")
	personPtr := &person2
	fmt.Printf("Via pointer: %s\n", personPtr.Name)
	personPtr.Age = 26
	fmt.Printf("Modified age: %d\n", person2.Age)
	
	fmt.Println("\n=== STRUCT COMPARISON ===")
	addr1 := Address{"123 Main St", "NYC", "USA"}
	addr2 := Address{"123 Main St", "NYC", "USA"}
	addr3 := Address{"456 Oak Ave", "LA", "USA"}
	
	fmt.Printf("addr1 == addr2: %t\n", addr1 == addr2)
	fmt.Printf("addr1 == addr3: %t\n", addr1 == addr3)
}
