// Performance optimization: profiling, benchmarking, dan memory management
package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

func inefficientStringConcatenation() string {
	result := ""
	for i := 0; i < 1000; i++ {
		result += fmt.Sprintf("Item %d ", i)
	}
	return result
}

func efficientStringConcatenation() string {
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf("Item %d ", i))
	}
	return builder.String()
}

func inefficientSliceGrowth() []int {
	var slice []int
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}
	return slice
}

func efficientSliceGrowth() []int {
	slice := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}
	return slice
}

func objectPoolExample() {
	pool := &sync.Pool{
		New: func() interface{} {
			return &Person{}
		},
	}
	
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			person := pool.Get().(*Person)
			person.Name = fmt.Sprintf("Person%d", id)
			person.Age = id
			
			time.Sleep(time.Millisecond)
			
			person.Name = ""
			person.Age = 0
			pool.Put(person)
		}(i)
	}
	wg.Wait()
}

func memoryAllocationExample() {
	fmt.Println("=== MEMORY ALLOCATION PATTERNS ===")
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Initial allocation: %d KB\n", m.Alloc/1024)
	
	largeSlice := make([]int, 1000000)
	runtime.ReadMemStats(&m)
	fmt.Printf("After large slice: %d KB\n", m.Alloc/1024)
	
	for i := range largeSlice {
		largeSlice[i] = i
	}
	runtime.ReadMemStats(&m)
	fmt.Printf("After filling slice: %d KB\n", m.Alloc/1024)
	
	largeSlice = nil
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("After GC: %d KB\n", m.Alloc/1024)
}

func stackVsHeapAllocation() {
	fmt.Println("\n=== STACK VS HEAP ALLOCATION ===")
	
	stackVariable := 42
	fmt.Printf("Stack variable address: %p\n", &stackVariable)
	
	heapVariable := new(int)
	*heapVariable = 42
	fmt.Printf("Heap variable address: %p\n", heapVariable)
	
	escapeToHeap := func() *int {
		localVar := 42
		return &localVar
	}
	
	escaped := escapeToHeap()
	fmt.Printf("Escaped variable address: %p\n", escaped)
}

func unsafeOperations() {
	fmt.Println("\n=== UNSAFE OPERATIONS ===")
	
	s := "hello"
	fmt.Printf("String: %s, Address: %p\n", s, &s)
	
	ptr := unsafe.Pointer(&s)
	fmt.Printf("Unsafe pointer: %p\n", ptr)
	
	bytes := *(*[]byte)(unsafe.Pointer(&s))
	fmt.Printf("String as bytes: %v\n", bytes)
	
	var x int64 = 42
	fmt.Printf("int64 size: %d bytes\n", unsafe.Sizeof(x))
	
	var person Person
	fmt.Printf("Person struct size: %d bytes\n", unsafe.Sizeof(person))
	fmt.Printf("Name offset: %d\n", unsafe.Offsetof(person.Name))
	fmt.Printf("Age offset: %d\n", unsafe.Offsetof(person.Age))
}

func cacheLineOptimization() {
	fmt.Println("\n=== CACHE LINE OPTIMIZATION ===")
	
	type BadStruct struct {
		a bool    // 1 byte
		b int64   // 8 bytes  
		c bool    // 1 byte
		d int64   // 8 bytes
	}
	
	type GoodStruct struct {
		b, d int64   // 16 bytes together
		a, c bool    // 2 bytes together
	}
	
	fmt.Printf("BadStruct size: %d bytes\n", unsafe.Sizeof(BadStruct{}))
	fmt.Printf("GoodStruct size: %d bytes\n", unsafe.Sizeof(GoodStruct{}))
}

func benchmarkComparison() {
	fmt.Println("\n=== BENCHMARK COMPARISON ===")
	
	start := time.Now()
	inefficientStringConcatenation()
	inefficientTime := time.Since(start)
	
	start = time.Now()
	efficientStringConcatenation()
	efficientTime := time.Since(start)
	
	fmt.Printf("Inefficient string concatenation: %v\n", inefficientTime)
	fmt.Printf("Efficient string concatenation: %v\n", efficientTime)
	fmt.Printf("Improvement: %.2fx faster\n", float64(inefficientTime)/float64(efficientTime))
	
	start = time.Now()
	inefficientSliceGrowth()
	inefficientSliceTime := time.Since(start)
	
	start = time.Now()
	efficientSliceGrowth()
	efficientSliceTime := time.Since(start)
	
	fmt.Printf("Inefficient slice growth: %v\n", inefficientSliceTime)
	fmt.Printf("Efficient slice growth: %v\n", efficientSliceTime)
	fmt.Printf("Improvement: %.2fx faster\n", float64(inefficientSliceTime)/float64(efficientSliceTime))
}

func garbageCollectionTuning() {
	fmt.Println("\n=== GARBAGE COLLECTION ===")
	
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("GC runs: %d\n", stats.NumGC)
	fmt.Printf("Total pause time: %v\n", time.Duration(stats.PauseTotalNs))
	
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024)
	}
	
	runtime.ReadMemStats(&stats)
	fmt.Printf("After allocation - GC runs: %d\n", stats.NumGC)
	fmt.Printf("Heap size: %d KB\n", stats.HeapAlloc/1024)
	
	data = nil
	runtime.GC()
	
	runtime.ReadMemStats(&stats)
	fmt.Printf("After forced GC - GC runs: %d\n", stats.NumGC)
	fmt.Printf("Heap size: %d KB\n", stats.HeapAlloc/1024)
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	
	memoryAllocationExample()
	stackVsHeapAllocation()
	unsafeOperations()
	cacheLineOptimization()
	benchmarkComparison()
	garbageCollectionTuning()
	
	fmt.Println("\n=== OBJECT POOL DEMO ===")
	start := time.Now()
	objectPoolExample()
	fmt.Printf("Object pool demo completed in: %v\n", time.Since(start))
	
	fmt.Println("\n=== RUNTIME STATISTICS ===")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	fmt.Printf("Allocated memory: %d KB\n", m.Alloc/1024)
	fmt.Printf("Total allocations: %d\n", m.TotalAlloc/1024)
	fmt.Printf("System memory: %d KB\n", m.Sys/1024)
	fmt.Printf("GC cycles: %d\n", m.NumGC)
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
}
