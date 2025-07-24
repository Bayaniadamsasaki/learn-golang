// Package management: membuat dan menggunakan package untuk modularitas
package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== BUILT-IN PACKAGES ===")
	
	fmt.Println("\n--- fmt package ---")
	name := "John"
	age := 25
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
	fmt.Sprintf("Formatted string: %s is %d", name, age)
	
	fmt.Println("\n--- strings package ---")
	text := "Hello, World!"
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("Upper: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower: %s\n", strings.ToLower(text))
	fmt.Printf("Contains 'World': %t\n", strings.Contains(text, "World"))
	fmt.Printf("Replace 'World' with 'Go': %s\n", strings.Replace(text, "World", "Go", 1))
	
	words := strings.Split("apple,banana,orange", ",")
	fmt.Printf("Split result: %v\n", words)
	joined := strings.Join(words, " | ")
	fmt.Printf("Joined: %s\n", joined)
	
	fmt.Println("\n--- math package ---")
	fmt.Printf("Pi: %.6f\n", math.Pi)
	fmt.Printf("E: %.6f\n", math.E)
	fmt.Printf("Sqrt(16): %.2f\n", math.Sqrt(16))
	fmt.Printf("Pow(2, 3): %.0f\n", math.Pow(2, 3))
	fmt.Printf("Max(10, 20): %.0f\n", math.Max(10, 20))
	fmt.Printf("Min(10, 20): %.0f\n", math.Min(10, 20))
	fmt.Printf("Abs(-5): %.0f\n", math.Abs(-5))
	fmt.Printf("Ceil(4.3): %.0f\n", math.Ceil(4.3))
	fmt.Printf("Floor(4.7): %.0f\n", math.Floor(4.7))
	
	fmt.Println("\n--- time package ---")
	now := time.Now()
	fmt.Printf("Current time: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("Unix timestamp: %d\n", now.Unix())
	
	birthday := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Birthday: %s\n", birthday.Format("January 2, 2006"))
	
	duration := now.Sub(birthday)
	fmt.Printf("Age in days: %.0f\n", duration.Hours()/24)
	
	fmt.Println("\n--- math/rand package ---")
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Random int: %d\n", rand.Intn(100))
	fmt.Printf("Random float: %.3f\n", rand.Float64())
	
	numbers := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Printf("Shuffled numbers: %v\n", numbers)
	
	fmt.Println("\n=== PACKAGE ALIASES ===")
	
	fmt.Println("\n=== BLANK IDENTIFIER ===")
	
	fmt.Println("\n=== INIT FUNCTION ===")
	fmt.Println("Init functions are called automatically when package is imported")
	
	fmt.Println("\n=== PACKAGE VISIBILITY ===")
	fmt.Println("Exported names start with capital letter (public)")
	fmt.Println("Unexported names start with lowercase letter (private)")
}
