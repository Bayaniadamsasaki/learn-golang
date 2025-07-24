// Konstanta: nilai yang tidak dapat diubah setelah dideklarasikan
package main

import "fmt"

const PI = 3.14159
const COMPANY_NAME = "Tech Corp"

const (
	MONDAY    = "Senin"
	TUESDAY   = "Selasa"
	WEDNESDAY = "Rabu"
)

func main() {
	const localConst = "Konstanta lokal"
	
	fmt.Println("PI:", PI)
	fmt.Println("Nama perusahaan:", COMPANY_NAME)
	fmt.Println("Hari:", MONDAY, TUESDAY, WEDNESDAY)
	fmt.Println("Konstanta lokal:", localConst)
	
	radius := 5.0
	area := PI * radius * radius
	fmt.Printf("Luas lingkaran dengan radius %.1f = %.2f\n", radius, area)
}
