package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		showMenu()
		choice := getChoice()
		
		if choice == 5 {
			fmt.Println("Terima kasih!")
			break
		}
		
		if choice < 1 || choice > 5 {
			fmt.Println("Pilihan tidak valid!")
			continue
		}
		
		num1 := getNumber("Masukkan angka pertama: ")
		num2 := getNumber("Masukkan angka kedua: ")
		
		switch choice {
		case 1:
			result := add(num1, num2)
			fmt.Printf("Hasil: %.2f + %.2f = %.2f\n", num1, num2, result)
		case 2:
			result := subtract(num1, num2)
			fmt.Printf("Hasil: %.2f - %.2f = %.2f\n", num1, num2, result)
		case 3:
			result := multiply(num1, num2)
			fmt.Printf("Hasil: %.2f * %.2f = %.2f\n", num1, num2, result)
		case 4:
			if num2 == 0 {
				fmt.Println("Error: Tidak bisa dibagi dengan nol!")
			} else {
				result := divide(num1, num2)
				fmt.Printf("Hasil: %.2f / %.2f = %.2f\n", num1, num2, result)
			}
		}
		
		fmt.Println()
	}
}

func showMenu() {
	fmt.Println("=== KALKULATOR SEDERHANA ===")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Perkalian")
	fmt.Println("4. Pembagian")
	fmt.Println("5. Keluar")
}

func getChoice() int {
	var choice int
	fmt.Print("Pilih operasi (1-5): ")
	fmt.Scanln(&choice)
	return choice
}

func getNumber(prompt string) float64 {
	var num float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Input tidak valid!")
		os.Exit(1)
	}
	return num
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}
