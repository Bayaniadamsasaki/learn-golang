// Fungsi: definisi, parameter, return value, dan variadic functions
package main

import "fmt"

func greet() {
	fmt.Println("Hello dari fungsi greet!")
}

func greetWithName(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int) int {
	return a + b
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("tidak bisa dibagi dengan nol")
	}
	return a / b, nil
}

func getPersonInfo() (string, int) {
	return "John Doe", 25
}

func namedReturn(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("=== FUNGSI SEDERHANA ===")
	greet()
	greetWithName("Alice")
	
	fmt.Println("\n=== FUNGSI DENGAN RETURN ===")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	product := multiply(4.5, 2.0)
	fmt.Printf("4.5 * 2.0 = %.1f\n", product)
	
	fmt.Println("\n=== FUNGSI DENGAN MULTIPLE RETURN ===")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}
	
	_, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	
	name, age := getPersonInfo()
	fmt.Printf("Nama: %s, Umur: %d\n", name, age)
	
	fmt.Println("\n=== NAMED RETURN ===")
	s, p := namedReturn(4, 5)
	fmt.Printf("Sum: %d, Product: %d\n", s, p)
	
	fmt.Println("\n=== VARIADIC FUNCTION ===")
	total1 := sum(1, 2, 3)
	total2 := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum(1,2,3) = %d\n", total1)
	fmt.Printf("Sum(1,2,3,4,5) = %d\n", total2)
	
	numbers := []int{10, 20, 30}
	total3 := sum(numbers...)
	fmt.Printf("Sum dari slice = %d\n", total3)
}
