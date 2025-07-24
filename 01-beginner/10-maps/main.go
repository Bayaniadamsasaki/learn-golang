// Maps: struktur data key-value untuk menyimpan pasangan data
package main

import "fmt"

func main() {
	fmt.Println("=== MEMBUAT MAP ===")
	var ages map[string]int
	ages = make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	
	scores := map[string]float64{
		"Math":    95.5,
		"Physics": 87.0,
		"Chemistry": 92.5,
	}
	
	fmt.Printf("Ages: %v\n", ages)
	fmt.Printf("Scores: %v\n", scores)
	
	fmt.Println("\n=== MENGAKSES MAP ===")
	mathScore := scores["Math"]
	fmt.Printf("Math score: %.1f\n", mathScore)
	
	englishScore, exists := scores["English"]
	if exists {
		fmt.Printf("English score: %.1f\n", englishScore)
	} else {
		fmt.Println("English score tidak ditemukan")
	}
	
	fmt.Println("\n=== MEMODIFIKASI MAP ===")
	scores["English"] = 89.0
	scores["Math"] = 98.0
	fmt.Printf("Updated scores: %v\n", scores)
	
	delete(scores, "Physics")
	fmt.Printf("After deleting Physics: %v\n", scores)
	
	fmt.Println("\n=== ITERASI MAP ===")
	for subject, score := range scores {
		fmt.Printf("%s: %.1f\n", subject, score)
	}
	
	fmt.Println("\n=== MAP dengan VALUE SLICE ===")
	studentGrades := map[string][]float64{
		"Alice":   {85.5, 92.0, 78.5},
		"Bob":     {90.0, 88.5, 95.0},
		"Charlie": {75.0, 80.0, 85.0},
	}
	
	for student, grades := range studentGrades {
		total := 0.0
		for _, grade := range grades {
			total += grade
		}
		average := total / float64(len(grades))
		fmt.Printf("%s - Average: %.2f\n", student, average)
	}
	
	fmt.Println("\n=== NESTED MAP ===")
	company := map[string]map[string]interface{}{
		"employee1": {
			"name":   "John",
			"age":    30,
			"salary": 50000,
		},
		"employee2": {
			"name":   "Jane",
			"age":    28,
			"salary": 55000,
		},
	}
	
	for id, employee := range company {
		fmt.Printf("%s: %v\n", id, employee)
	}
	
	fmt.Println("\n=== MAP OPERATIONS ===")
	fmt.Printf("Length of scores map: %d\n", len(scores))
	
	if _, exists := scores["Biology"]; !exists {
		fmt.Println("Biology tidak ada dalam map")
	}
}
