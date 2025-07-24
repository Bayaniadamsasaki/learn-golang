// Control flow: if-else, switch, dan kondisi bersyarat
package main

import "fmt"

func main() {
	age := 20
	
	fmt.Println("=== IF-ELSE ===")
	if age >= 18 {
		fmt.Println("Dewasa")
	} else {
		fmt.Println("Anak-anak")
	}
	
	if age < 13 {
		fmt.Println("Kategori: Anak")
	} else if age < 20 {
		fmt.Println("Kategori: Remaja")
	} else if age < 60 {
		fmt.Println("Kategori: Dewasa")
	} else {
		fmt.Println("Kategori: Lansia")
	}
	
	fmt.Println("\n=== SWITCH ===")
	day := 3
	switch day {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	default:
		fmt.Println("Akhir pekan")
	}
	
	grade := 85
	switch {
	case grade >= 90:
		fmt.Println("Nilai A")
	case grade >= 80:
		fmt.Println("Nilai B")
	case grade >= 70:
		fmt.Println("Nilai C")
	case grade >= 60:
		fmt.Println("Nilai D")
	default:
		fmt.Println("Nilai E")
	}
	
	fmt.Println("\n=== SHORT VARIABLE DECLARATION IN IF ===")
	if score := 95; score >= 90 {
		fmt.Printf("Excellent! Score: %d\n", score)
	}
}
