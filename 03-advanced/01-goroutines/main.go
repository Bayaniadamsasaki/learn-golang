// Goroutines: concurrency dengan lightweight threads
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func simpleTask(id int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Task %d: Step %d\n", id, i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Task %d completed\n", id)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func counter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("Counter: %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	fmt.Printf("Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("Number of Goroutines: %d\n", runtime.NumGoroutine())
	
	fmt.Println("\n=== BASIC GOROUTINES ===")
	go simpleTask(1)
	go simpleTask(2)
	
	simpleTask(3)
	
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n=== ANONYMOUS GOROUTINES ===")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Anonymous goroutine %d\n", id)
		}(i)
	}
	
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("\n=== WAITGROUP ===")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("WaitGroup task %d started\n", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("WaitGroup task %d completed\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All WaitGroup tasks completed")
	
	fmt.Println("\n=== WORKER POOL PATTERN ===")
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	
	fmt.Println("\n=== GOROUTINE WITH CLOSURE ===")
	messages := []string{"Hello", "World", "Go", "Concurrency"}
	
	var wg2 sync.WaitGroup
	for _, msg := range messages {
		wg2.Add(1)
		go func(message string) {
			defer wg2.Done()
			fmt.Printf("Processing: %s\n", message)
			time.Sleep(100 * time.Millisecond)
		}(msg)
	}
	wg2.Wait()
	
	fmt.Println("\n=== RACE CONDITION EXAMPLE ===")
	var counter1 int
	var wg3 sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			counter1++
		}()
	}
	wg3.Wait()
	fmt.Printf("Counter without mutex: %d (may vary)\n", counter1)
	
	fmt.Println("\n=== MUTEX FOR SAFE ACCESS ===")
	var counter2 int
	var mu sync.Mutex
	var wg4 sync.WaitGroup
	
	for i := 0; i < 1000; i++ {
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			mu.Lock()
			counter2++
			mu.Unlock()
		}()
	}
	wg4.Wait()
	fmt.Printf("Counter with mutex: %d (should be 1000)\n", counter2)
	
	fmt.Println("\n=== GOROUTINE LEAK PREVENTION ===")
	done := make(chan bool)
	
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Goroutine received done signal")
				return
			default:
				fmt.Println("Goroutine working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(500 * time.Millisecond)
	done <- true
	time.Sleep(100 * time.Millisecond)
	
	fmt.Printf("Final number of Goroutines: %d\n", runtime.NumGoroutine())
}
