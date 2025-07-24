// Generics: type parameters untuk membuat code yang reusable dan type-safe
package main

import (
	"fmt"
	"constraints"
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

type Pair[T, U any] struct {
	First  T
	Second U
}

func MakePair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

func Sort[T Comparable](slice []T) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

type Container[T any] interface {
	Add(T)
	Remove() (T, bool)
	Size() int
	IsEmpty() bool
}

type List[T any] struct {
	items []T
}

func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

func (l *List[T]) Remove() (T, bool) {
	if len(l.items) == 0 {
		var zero T
		return zero, false
	}
	item := l.items[len(l.items)-1]
	l.items = l.items[:len(l.items)-1]
	return item, true
}

func (l *List[T]) Size() int {
	return len(l.items)
}

func (l *List[T]) IsEmpty() bool {
	return len(l.items) == 0
}

func ProcessContainer[T any](container Container[T], items []T) {
	for _, item := range items {
		container.Add(item)
	}
	
	fmt.Printf("Container size: %d\n", container.Size())
	
	for !container.IsEmpty() {
		item, _ := container.Remove()
		fmt.Printf("Removed: %v\n", item)
	}
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Number](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Average[T Number](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := Sum(numbers)
	return float64(sum) / float64(len(numbers))
}

type Cache[K comparable, V any] struct {
	data map[K]V
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}

func (c *Cache[K, V]) Keys() []K {
	keys := make([]K, 0, len(c.data))
	for k := range c.data {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	fmt.Println("=== BASIC GENERIC FUNCTIONS ===")
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))
	
	fmt.Printf("Min(10, 20): %d\n", Min(10, 20))
	fmt.Printf("Min(3.14, 2.71): %.2f\n", Min(3.14, 2.71))
	
	fmt.Println("\n=== GENERIC SLICE OPERATIONS ===")
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubled)
	
	strings := []string{"hello", "world", "go", "generics"}
	lengths := Map(strings, func(s string) int { return len(s) })
	fmt.Printf("Strings: %v\n", strings)
	fmt.Printf("Lengths: %v\n", lengths)
	
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)
	
	longStrings := Filter(strings, func(s string) bool { return len(s) > 3 })
	fmt.Printf("Long strings: %v\n", longStrings)
	
	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Sum: %d\n", sum)
	
	product := Reduce(numbers, 1, func(acc, x int) int { return acc * x })
	fmt.Printf("Product: %d\n", product)
	
	fmt.Println("\n=== GENERIC DATA STRUCTURES ===")
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	
	fmt.Printf("Stack size: %d\n", intStack.Size())
	for !intStack.IsEmpty() {
		item, _ := intStack.Pop()
		fmt.Printf("Popped: %d\n", item)
	}
	
	stringQueue := &Queue[string]{}
	stringQueue.Enqueue("first")
	stringQueue.Enqueue("second")
	stringQueue.Enqueue("third")
	
	for !stringQueue.IsEmpty() {
		item, _ := stringQueue.Dequeue()
		fmt.Printf("Dequeued: %s\n", item)
	}
	
	fmt.Println("\n=== GENERIC PAIRS ===")
	pair1 := MakePair(1, "one")
	pair2 := MakePair("hello", 42)
	pair3 := MakePair(3.14, true)
	
	fmt.Printf("Pair1: %s\n", pair1)
	fmt.Printf("Pair2: %s\n", pair2)
	fmt.Printf("Pair3: %s\n", pair3)
	
	fmt.Println("\n=== GENERIC SORTING ===")
	intSlice := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Before sorting: %v\n", intSlice)
	Sort(intSlice)
	fmt.Printf("After sorting: %v\n", intSlice)
	
	stringSlice := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Before sorting: %v\n", stringSlice)
	Sort(stringSlice)
	fmt.Printf("After sorting: %v\n", stringSlice)
	
	fmt.Println("\n=== GENERIC INTERFACES ===")
	list := &List[string]{}
	ProcessContainer(list, []string{"a", "b", "c"})
	
	fmt.Println("\n=== NUMERIC OPERATIONS ===")
	intNumbers := []int{1, 2, 3, 4, 5}
	floatNumbers := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	
	fmt.Printf("Int sum: %d\n", Sum(intNumbers))
	fmt.Printf("Int average: %.2f\n", Average(intNumbers))
	
	fmt.Printf("Float sum: %.1f\n", Sum(floatNumbers))
	fmt.Printf("Float average: %.2f\n", Average(floatNumbers))
	
	fmt.Println("\n=== GENERIC CACHE ===")
	stringCache := NewCache[string, int]()
	stringCache.Set("apple", 5)
	stringCache.Set("banana", 6)
	stringCache.Set("cherry", 6)
	
	if value, exists := stringCache.Get("apple"); exists {
		fmt.Printf("apple: %d\n", value)
	}
	
	fmt.Printf("Cache keys: %v\n", stringCache.Keys())
	
	intCache := NewCache[int, string]()
	intCache.Set(1, "one")
	intCache.Set(2, "two")
	intCache.Set(3, "three")
	
	if value, exists := intCache.Get(2); exists {
		fmt.Printf("2: %s\n", value)
	}
	
	fmt.Printf("Int cache keys: %v\n", intCache.Keys())
}
