// Package utils: utility functions untuk matematika dasar
package utils

import "math"

var PI = math.Pi

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func CircleArea(radius float64) float64 {
	return PI * radius * radius
}

func privateFunction() string {
	return "This function is not exported"
}

type Calculator struct {
	Result float64
}

func (c *Calculator) Add(value float64) {
	c.Result += value
}

func (c *Calculator) Multiply(value float64) {
	c.Result *= value
}

func (c *Calculator) Reset() {
	c.Result = 0
}
