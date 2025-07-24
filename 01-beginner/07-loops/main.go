// Perulangan: for loop dengan berbagai variasi
package main

import "fmt"

func main() {
	fmt.Println("=== FOR LOOP DASAR ===")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iterasi %d\n", i)
	}
	
	fmt.Println("\n=== WHILE-LIKE LOOP ===")
	counter := 1
	for counter <= 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}
	
	fmt.Println("\n=== INFINITE LOOP dengan BREAK ===")
	num := 1
	for {
		if num > 3 {
			break
		}
		fmt.Printf("Angka: %d\n", num)
		num++
	}
	
	fmt.Println("\n=== CONTINUE ===")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Printf("Nilai: %d\n", i)
	}
	
	fmt.Println("\n=== NESTED LOOP ===")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== RANGE OVER SLICE ===")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	
	fmt.Println("\n=== RANGE OVER STRING ===")
	text := "Hello"
	for i, char := range text {
		fmt.Printf("Index: %d, Char: %c\n", i, char)
	}
}
