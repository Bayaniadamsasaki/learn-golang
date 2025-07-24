// Tipe data dasar: integer, float, boolean, string dan operasinya
package main

import "fmt"

func main() {
	var intVar int = 42
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807
	
	var uintVar uint = 42
	var uint8Var uint8 = 255
	
	var float32Var float32 = 3.14
	var float64Var float64 = 3.141592653589793
	
	var boolVar bool = true
	var stringVar string = "Hello, Golang!"
	
	var byteVar byte = 65
	var runeVar rune = 'A'
	
	fmt.Printf("int: %d, type: %T\n", intVar, intVar)
	fmt.Printf("int8: %d, type: %T\n", int8Var, int8Var)
	fmt.Printf("int16: %d, type: %T\n", int16Var, int16Var)
	fmt.Printf("int32: %d, type: %T\n", int32Var, int32Var)
	fmt.Printf("int64: %d, type: %T\n", int64Var, int64Var)
	fmt.Printf("uint: %d, type: %T\n", uintVar, uintVar)
	fmt.Printf("uint8: %d, type: %T\n", uint8Var, uint8Var)
	fmt.Printf("float32: %.2f, type: %T\n", float32Var, float32Var)
	fmt.Printf("float64: %.15f, type: %T\n", float64Var, float64Var)
	fmt.Printf("bool: %t, type: %T\n", boolVar, boolVar)
	fmt.Printf("string: %s, type: %T\n", stringVar, stringVar)
	fmt.Printf("byte: %d (%c), type: %T\n", byteVar, byteVar, byteVar)
	fmt.Printf("rune: %d (%c), type: %T\n", runeVar, runeVar, runeVar)
}
