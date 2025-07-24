// Belajar variabel: deklarasi, inisialisasi, dan tipe data dasar
package main

import "fmt"

func main() {
	var name string = "John Doe"
	var age int = 25
	var height float64 = 175.5
	var isStudent bool = true
	
	shortName := "Jane"
	shortAge := 22
	
	fmt.Println("Nama:", name)
	fmt.Println("Umur:", age)
	fmt.Println("Tinggi:", height, "cm")
	fmt.Println("Status mahasiswa:", isStudent)
	fmt.Printf("Nama pendek: %s, Umur: %d\n", shortName, shortAge)
}
