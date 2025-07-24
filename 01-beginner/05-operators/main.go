// Operator: aritmatika, perbandingan, logika, dan assignment
package main

import "fmt"

func main() {
	a := 10
	b := 3
	
	fmt.Println("=== OPERATOR ARITMATIKA ===")
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	
	fmt.Println("\n=== OPERATOR PERBANDINGAN ===")
	fmt.Printf("%d == %d = %t\n", a, b, a == b)
	fmt.Printf("%d != %d = %t\n", a, b, a != b)
	fmt.Printf("%d > %d = %t\n", a, b, a > b)
	fmt.Printf("%d < %d = %t\n", a, b, a < b)
	fmt.Printf("%d >= %d = %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d = %t\n", a, b, a <= b)
	
	fmt.Println("\n=== OPERATOR LOGIKA ===")
	x, y := true, false
	fmt.Printf("%t && %t = %t\n", x, y, x && y)
	fmt.Printf("%t || %t = %t\n", x, y, x || y)
	fmt.Printf("!%t = %t\n", x, !x)
	
	fmt.Println("\n=== OPERATOR ASSIGNMENT ===")
	c := 5
	fmt.Printf("c = %d\n", c)
	c += 3
	fmt.Printf("c += 3, c = %d\n", c)
	c -= 2
	fmt.Printf("c -= 2, c = %d\n", c)
	c *= 2
	fmt.Printf("c *= 2, c = %d\n", c)
	c /= 3
	fmt.Printf("c /= 3, c = %d\n", c)
}
