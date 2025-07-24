// JSON handling: encoding dan decoding data JSON
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type Company struct {
	Name      string   `json:"name"`
	Founded   int      `json:"founded"`
	Employees []Person `json:"employees"`
}

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Tags        []string `json:"tags"`
	InStock     bool     `json:"in_stock"`
	Description string   `json:"description,omitempty"`
}

func main() {
	fmt.Println("=== BASIC JSON ENCODING ===")
	person := Person{
		Name:      "John Doe",
		Age:       30,
		Email:     "john@example.com",
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("Error encoding: %v\n", err)
		return
	}
	
	fmt.Printf("JSON: %s\n", jsonData)
	
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding: %v\n", err)
		return
	}
	
	fmt.Printf("Pretty JSON:\n%s\n", prettyJSON)
	
	fmt.Println("\n=== BASIC JSON DECODING ===")
	jsonString := `{"name":"Jane Smith","age":25,"email":"jane@example.com","is_active":false,"created_at":"2023-01-15T10:30:00Z"}`
	
	var decodedPerson Person
	err = json.Unmarshal([]byte(jsonString), &decodedPerson)
	if err != nil {
		fmt.Printf("Error decoding: %v\n", err)
		return
	}
	
	fmt.Printf("Decoded person: %+v\n", decodedPerson)
	
	fmt.Println("\n=== NESTED STRUCTURES ===")
	company := Company{
		Name:    "Tech Corp",
		Founded: 2010,
		Employees: []Person{
			{Name: "Alice", Age: 28, Email: "alice@techcorp.com", IsActive: true, CreatedAt: time.Now()},
			{Name: "Bob", Age: 32, Email: "bob@techcorp.com", IsActive: true, CreatedAt: time.Now()},
		},
	}
	
	companyJSON, err := json.MarshalIndent(company, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding company: %v\n", err)
		return
	}
	
	fmt.Printf("Company JSON:\n%s\n", companyJSON)
	
	fmt.Println("\n=== WORKING WITH SLICES ===")
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 999.99, Tags: []string{"electronics", "computer"}, InStock: true},
		{ID: 2, Name: "Mouse", Price: 29.99, Tags: []string{"electronics", "accessory"}, InStock: false},
		{ID: 3, Name: "Book", Price: 19.99, Tags: []string{"education", "reading"}, InStock: true, Description: "Programming guide"},
	}
	
	productsJSON, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding products: %v\n", err)
		return
	}
	
	fmt.Printf("Products JSON:\n%s\n", productsJSON)
	
	fmt.Println("\n=== DECODING INTO SLICE ===")
	productJSONString := `[
		{"id":4,"name":"Keyboard","price":79.99,"tags":["electronics","input"],"in_stock":true},
		{"id":5,"name":"Monitor","price":299.99,"tags":["electronics","display"],"in_stock":false}
	]`
	
	var decodedProducts []Product
	err = json.Unmarshal([]byte(productJSONString), &decodedProducts)
	if err != nil {
		fmt.Printf("Error decoding products: %v\n", err)
		return
	}
	
	for _, product := range decodedProducts {
		fmt.Printf("Product: %+v\n", product)
	}
	
	fmt.Println("\n=== INTERFACE{} FOR DYNAMIC JSON ===")
	dynamicJSON := `{
		"string_field": "hello",
		"number_field": 42,
		"boolean_field": true,
		"array_field": [1, 2, 3],
		"object_field": {
			"nested_string": "world",
			"nested_number": 3.14
		}
	}`
	
	var dynamicData interface{}
	err = json.Unmarshal([]byte(dynamicJSON), &dynamicData)
	if err != nil {
		fmt.Printf("Error decoding dynamic JSON: %v\n", err)
		return
	}
	
	fmt.Printf("Dynamic data: %+v\n", dynamicData)
	
	if dataMap, ok := dynamicData.(map[string]interface{}); ok {
		for key, value := range dataMap {
			fmt.Printf("Key: %s, Value: %v, Type: %T\n", key, value, value)
		}
	}
	
	fmt.Println("\n=== MAP FOR FLEXIBLE JSON ===")
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(dynamicJSON), &jsonMap)
	if err != nil {
		fmt.Printf("Error decoding to map: %v\n", err)
		return
	}
	
	if stringField, exists := jsonMap["string_field"]; exists {
		fmt.Printf("String field: %s\n", stringField)
	}
	
	if numberField, exists := jsonMap["number_field"]; exists {
		if num, ok := numberField.(float64); ok {
			fmt.Printf("Number field: %.0f\n", num)
		}
	}
	
	fmt.Println("\n=== JSON TAGS ===")
	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"-"`
		FullName string `json:"full_name,omitempty"`
		Internal string `json:"internal,omitempty"`
	}
	
	user := User{
		ID:       1,
		Username: "johndoe",
		Password: "secret123",
		FullName: "",
		Internal: "internal_data",
	}
	
	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding user: %v\n", err)
		return
	}
	
	fmt.Printf("User JSON (note omitted fields):\n%s\n", userJSON)
	
	fmt.Println("\n=== ERROR HANDLING ===")
	invalidJSON := `{"name": "John", "age": "not_a_number"}`
	
	var invalidPerson Person
	err = json.Unmarshal([]byte(invalidJSON), &invalidPerson)
	if err != nil {
		fmt.Printf("Expected error decoding invalid JSON: %v\n", err)
	}
}
