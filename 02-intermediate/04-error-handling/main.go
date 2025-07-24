// Error handling: menangani kesalahan dengan elegant di Go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("failed to parse age '%s': %w", s, err)
	}
	if age < 0 {
		return 0, errors.New("age cannot be negative")
	}
	if age > 150 {
		return 0, errors.New("age cannot be greater than 150")
	}
	return age, nil
}

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s' with value '%v': %s", ve.Field, ve.Value, ve.Message)
}

type User struct {
	Name  string
	Email string
	Age   int
}

func validateUser(u User) error {
	if u.Name == "" {
		return ValidationError{
			Field:   "name",
			Value:   u.Name,
			Message: "name cannot be empty",
		}
	}
	
	if len(u.Name) < 2 {
		return ValidationError{
			Field:   "name",
			Value:   u.Name,
			Message: "name must be at least 2 characters",
		}
	}
	
	if u.Email == "" {
		return ValidationError{
			Field:   "email",
			Value:   u.Email,
			Message: "email cannot be empty",
		}
	}
	
	if u.Age < 0 || u.Age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   u.Age,
			Message: "age must be between 0 and 150",
		}
	}
	
	return nil
}

func processUser(name, email, ageStr string) error {
	age, err := parseAge(ageStr)
	if err != nil {
		return fmt.Errorf("error processing user: %w", err)
	}
	
	user := User{
		Name:  name,
		Email: email,
		Age:   age,
	}
	
	if err := validateUser(user); err != nil {
		return fmt.Errorf("user validation failed: %w", err)
	}
	
	fmt.Printf("User processed successfully: %+v\n", user)
	return nil
}

func multipleOperations() error {
	result1, err := divide(10, 2)
	if err != nil {
		return fmt.Errorf("first division failed: %w", err)
	}
	
	result2, err := divide(result1, 0)
	if err != nil {
		return fmt.Errorf("second division failed: %w", err)
	}
	
	fmt.Printf("Final result: %.2f\n", result2)
	return nil
}

func recoverFromPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	fmt.Println("About to panic...")
	panic("something went wrong!")
	fmt.Println("This line will not be executed")
}

func main() {
	fmt.Println("=== BASIC ERROR HANDLING ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
	
	fmt.Println("\n=== ERROR WRAPPING ===")
	age, err := parseAge("25")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed age: %d\n", age)
	}
	
	age, err = parseAge("invalid")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	age, err = parseAge("-5")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	fmt.Println("\n=== CUSTOM ERROR TYPES ===")
	validUser := User{Name: "John Doe", Email: "john@example.com", Age: 30}
	if err := validateUser(validUser); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("User is valid")
	}
	
	invalidUser := User{Name: "", Email: "invalid@example.com", Age: 25}
	if err := validateUser(invalidUser); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		
		var validationErr ValidationError
		if errors.As(err, &validationErr) {
			fmt.Printf("Field: %s, Value: %v\n", validationErr.Field, validationErr.Value)
		}
	}
	
	fmt.Println("\n=== ERROR CHAIN ===")
	if err := processUser("Alice", "alice@example.com", "28"); err != nil {
		fmt.Printf("Process error: %v\n", err)
	}
	
	if err := processUser("", "bob@example.com", "invalid"); err != nil {
		fmt.Printf("Process error: %v\n", err)
	}
	
	fmt.Println("\n=== MULTIPLE ERROR HANDLING ===")
	if err := multipleOperations(); err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		
		fmt.Println("Unwrapping errors:")
		currentErr := err
		for currentErr != nil {
			fmt.Printf("  - %v\n", currentErr)
			currentErr = errors.Unwrap(currentErr)
		}
	}
	
	fmt.Println("\n=== ERROR CHECKING ===")
	testErr := errors.New("division by zero")
	wrappedErr := fmt.Errorf("calculation failed: %w", testErr)
	
	if errors.Is(wrappedErr, testErr) {
		fmt.Println("wrappedErr contains testErr")
	}
	
	fmt.Println("\n=== PANIC AND RECOVER ===")
	recoverFromPanic()
	fmt.Println("Program continues after panic recovery")
	
	fmt.Println("\n=== DEFER WITH ERROR ===")
	func() {
		defer func() {
			fmt.Println("Cleanup operation in defer")
		}()
		
		_, err := divide(5, 0)
		if err != nil {
			fmt.Printf("Deferred function will still run. Error: %v\n", err)
			return
		}
	}()
}
