// Arrays dan Slices: koleksi data dengan ukuran tetap dan dinamis
package main

import "fmt"

func main() {
	fmt.Println("=== ARRAYS ===")
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	
	fruits := [3]string{"apple", "banana", "orange"}
	colors := [...]string{"red", "green", "blue", "yellow"}
	
	fmt.Printf("Numbers array: %v\n", numbers)
	fmt.Printf("Fruits array: %v\n", fruits)
	fmt.Printf("Colors array: %v (length: %d)\n", colors, len(colors))
	
	for i, value := range numbers {
		fmt.Printf("Index %d: %d\n", i, value)
	}
	
	fmt.Println("\n=== SLICES ===")
	var scores []int
	scores = append(scores, 85)
	scores = append(scores, 90, 75, 88)
	
	names := []string{"Alice", "Bob", "Charlie"}
	grades := make([]float64, 3, 5)
	grades[0] = 85.5
	grades[1] = 92.0
	grades[2] = 78.5
	
	fmt.Printf("Scores: %v (len: %d, cap: %d)\n", scores, len(scores), cap(scores))
	fmt.Printf("Names: %v (len: %d, cap: %d)\n", names, len(names), cap(names))
	fmt.Printf("Grades: %v (len: %d, cap: %d)\n", grades, len(grades), cap(grades))
	
	fmt.Println("\n=== SLICE OPERATIONS ===")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("nums[2:5]: %v\n", nums[2:5])
	fmt.Printf("nums[:4]: %v\n", nums[:4])
	fmt.Printf("nums[6:]: %v\n", nums[6:])
	fmt.Printf("nums[:]: %v\n", nums[:])
	
	subSlice := nums[2:7]
	fmt.Printf("Sub-slice: %v\n", subSlice)
	subSlice[0] = 100
	fmt.Printf("After modifying sub-slice: %v\n", nums)
	
	fmt.Println("\n=== COPY SLICE ===")
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)
	copied[0] = 999
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copied: %v\n", copied)
	
	fmt.Println("\n=== MULTI-DIMENSIONAL ===")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %d ", i, j, value)
		}
		fmt.Println()
	}
}
