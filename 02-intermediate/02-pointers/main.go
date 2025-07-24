// Pointers: referensi memori untuk mengelola data secara efisien
package main

import "fmt"

func modifyValue(x int) {
	x = 100
}

func modifyValueByPointer(x *int) {
	*x = 100
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) AreaByValue() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) AreaByPointer() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	fmt.Println("=== BASIC POINTERS ===")
	num := 42
	fmt.Printf("Value: %d, Address: %p\n", num, &num)
	
	var ptr *int
	ptr = &num
	fmt.Printf("Pointer value: %p\n", ptr)
	fmt.Printf("Value at pointer: %d\n", *ptr)
	
	*ptr = 99
	fmt.Printf("Modified num: %d\n", num)
	
	fmt.Println("\n=== ZERO VALUE OF POINTER ===")
	var nilPtr *int
	fmt.Printf("Nil pointer: %v\n", nilPtr)
	if nilPtr == nil {
		fmt.Println("Pointer is nil")
	}
	
	fmt.Println("\n=== FUNCTION PARAMETERS ===")
	x := 50
	fmt.Printf("Before modifyValue: %d\n", x)
	modifyValue(x)
	fmt.Printf("After modifyValue: %d\n", x)
	
	fmt.Printf("Before modifyValueByPointer: %d\n", x)
	modifyValueByPointer(&x)
	fmt.Printf("After modifyValueByPointer: %d\n", x)
	
	a, b := 10, 20
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)
	
	fmt.Println("\n=== POINTERS WITH ARRAYS ===")
	arr := [3]int{1, 2, 3}
	arrPtr := &arr
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Array via pointer: %v\n", *arrPtr)
	
	(*arrPtr)[0] = 100
	fmt.Printf("Modified array: %v\n", arr)
	
	fmt.Println("\n=== POINTERS WITH SLICES ===")
	slice := []int{10, 20, 30}
	slicePtr := &slice
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Slice via pointer: %v\n", *slicePtr)
	
	*slicePtr = append(*slicePtr, 40)
	fmt.Printf("Extended slice: %v\n", slice)
	
	fmt.Println("\n=== STRUCT POINTERS ===")
	rect := Rectangle{Width: 10, Height: 5}
	rectPtr := &rect
	
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area (by value): %.2f\n", rect.AreaByValue())
	fmt.Printf("Area (by pointer): %.2f\n", rectPtr.AreaByPointer())
	
	fmt.Printf("Before scaling: %+v\n", rect)
	rectPtr.Scale(2.0)
	fmt.Printf("After scaling: %+v\n", rect)
	
	fmt.Println("\n=== POINTER ARITHMETIC (NOT ALLOWED IN GO) ===")
	numbers := []int{1, 2, 3, 4, 5}
	firstPtr := &numbers[0]
	fmt.Printf("First element: %d at %p\n", *firstPtr, firstPtr)
	
	fmt.Println("\n=== NEW FUNCTION ===")
	newPtr := new(int)
	fmt.Printf("new(int): %p, value: %d\n", newPtr, *newPtr)
	*newPtr = 42
	fmt.Printf("After assignment: %d\n", *newPtr)
	
	newRectPtr := new(Rectangle)
	newRectPtr.Width = 15
	newRectPtr.Height = 8
	fmt.Printf("New rectangle: %+v\n", *newRectPtr)
}
