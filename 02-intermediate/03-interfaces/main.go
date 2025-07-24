// Interfaces: kontrak yang mendefinisikan behavior tanpa implementasi
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Drawable interface {
	Draw()
}

type ShapeDrawer interface {
	Shape
	Drawable
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Draw() {
	fmt.Printf("Drawing rectangle %.1fx%.1f\n", r.Width, r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Draw() {
	fmt.Printf("Drawing circle with radius %.1f\n", c.Radius)
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + math.Sqrt(t.Base*t.Base+t.Height*t.Height)
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func drawShape(d Drawable) {
	d.Draw()
}

func processShapeDrawer(sd ShapeDrawer) {
	sd.Draw()
	fmt.Printf("Area: %.2f\n", sd.Area())
}

type Writer interface {
	Write(string)
}

type FileWriter struct {
	filename string
}

func (fw FileWriter) Write(data string) {
	fmt.Printf("Writing '%s' to file: %s\n", data, fw.filename)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) {
	fmt.Printf("Console output: %s\n", data)
}

func writeData(w Writer, data string) {
	w.Write(data)
}

func main() {
	fmt.Println("=== BASIC INTERFACE ===")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}
	triangle := Triangle{Base: 6, Height: 4}
	
	fmt.Println("Rectangle:")
	printShapeInfo(rect)
	
	fmt.Println("Circle:")
	printShapeInfo(circle)
	
	fmt.Println("Triangle:")
	printShapeInfo(triangle)
	
	fmt.Println("\n=== INTERFACE SLICE ===")
	shapes := []Shape{rect, circle, triangle}
	
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		printShapeInfo(shape)
	}
	
	fmt.Println("\n=== MULTIPLE INTERFACES ===")
	drawShape(rect)
	drawShape(circle)
	
	fmt.Println("\n=== EMBEDDED INTERFACE ===")
	processShapeDrawer(rect)
	processShapeDrawer(circle)
	
	fmt.Println("\n=== EMPTY INTERFACE ===")
	var anything interface{}
	
	anything = 42
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)
	
	anything = "Hello"
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)
	
	anything = rect
	fmt.Printf("anything = %v (type: %T)\n", anything, anything)
	
	fmt.Println("\n=== TYPE ASSERTION ===")
	var shapeInterface Shape = circle
	
	if c, ok := shapeInterface.(Circle); ok {
		fmt.Printf("It's a circle with radius: %.1f\n", c.Radius)
	}
	
	if _, ok := shapeInterface.(Rectangle); !ok {
		fmt.Println("It's not a rectangle")
	}
	
	fmt.Println("\n=== TYPE SWITCH ===")
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case Rectangle:
			fmt.Printf("Rectangle: %.1fx%.1f\n", v.Width, v.Height)
		case Circle:
			fmt.Printf("Circle with radius: %.1f\n", v.Radius)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}
	
	checkType(42)
	checkType("Hello")
	checkType(rect)
	checkType(circle)
	checkType(3.14)
	
	fmt.Println("\n=== INTERFACE POLYMORPHISM ===")
	fileWriter := FileWriter{filename: "data.txt"}
	consoleWriter := ConsoleWriter{}
	
	writers := []Writer{fileWriter, consoleWriter}
	
	for i, writer := range writers {
		writeData(writer, fmt.Sprintf("Message %d", i+1))
	}
	
	fmt.Println("\n=== NIL INTERFACE ===")
	var nilShape Shape
	fmt.Printf("Nil interface: %v\n", nilShape)
	
	if nilShape == nil {
		fmt.Println("Shape interface is nil")
	}
}
