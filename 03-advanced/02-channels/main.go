// Channels: komunikasi antar goroutines dengan type-safe message passing
package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string, name string) {
	for i := 1; i <= 3; i++ {
		message := fmt.Sprintf("Message %d from %s", i, name)
		ch <- message
		fmt.Printf("Sent: %s\n", message)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func receiver(ch <-chan string) {
	for message := range ch {
		fmt.Printf("Received: %s\n", message)
	}
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}

func consumer(id int, ch <-chan int) {
	for value := range ch {
		fmt.Printf("Consumer %d consumed: %d\n", id, value)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	fmt.Println("=== BASIC CHANNEL ===")
	ch := make(chan string)
	
	go func() {
		ch <- "Hello from goroutine"
	}()
	
	message := <-ch
	fmt.Printf("Received: %s\n", message)
	
	fmt.Println("\n=== BUFFERED CHANNEL ===")
	bufferedCh := make(chan int, 3)
	
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	
	fmt.Printf("Channel length: %d, capacity: %d\n", len(bufferedCh), cap(bufferedCh))
	
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	fmt.Printf("Received: %d\n", <-bufferedCh)
	
	fmt.Println("\n=== CHANNEL DIRECTIONS ===")
	ch2 := make(chan string, 2)
	
	go sender(ch2, "Sender1")
	go receiver(ch2)
	
	time.Sleep(2 * time.Second)
	
	fmt.Println("\n=== RANGE OVER CHANNEL ===")
	fibCh := make(chan int)
	go fibonacci(10, fibCh)
	
	fmt.Print("Fibonacci sequence: ")
	for num := range fibCh {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
	
	fmt.Println("\n=== SELECT STATEMENT ===")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Message from ch2"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}
	
	fmt.Println("\n=== SELECT WITH TIMEOUT ===")
	slowCh := make(chan string)
	
	go func() {
		time.Sleep(2 * time.Second)
		slowCh <- "Slow message"
	}()
	
	select {
	case msg := <-slowCh:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout occurred")
	}
	
	fmt.Println("\n=== SELECT WITH DEFAULT ===")
	nonBlockingCh := make(chan string)
	
	select {
	case msg := <-nonBlockingCh:
		fmt.Printf("Received: %s\n", msg)
	default:
		fmt.Println("No message available")
	}
	
	select {
	case nonBlockingCh <- "test":
		fmt.Println("Sent message")
	default:
		fmt.Println("Channel not ready for sending")
	}
	
	fmt.Println("\n=== PRODUCER-CONSUMER PATTERN ===")
	prodCh := make(chan int, 2)
	
	go producer(prodCh)
	go consumer(1, prodCh)
	go consumer(2, prodCh)
	
	time.Sleep(3 * time.Second)
	
	fmt.Println("\n=== CHANNEL CLOSING ===")
	closeCh := make(chan int, 3)
	closeCh <- 1
	closeCh <- 2
	closeCh <- 3
	close(closeCh)
	
	for {
		value, ok := <-closeCh
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("Received: %d\n", value)
	}
	
	fmt.Println("\n=== FAN-OUT FAN-IN PATTERN ===")
	input := make(chan int)
	
	c1 := make(chan int)
	c2 := make(chan int)
	
	go func() {
		for n := range input {
			c1 <- n * n
		}
		close(c1)
	}()
	
	go func() {
		for n := range input {
			c2 <- n * n * n
		}
		close(c2)
	}()
	
	go func() {
		for i := 1; i <= 3; i++ {
			input <- i
		}
		close(input)
	}()
	
	for i := 0; i < 6; i++ {
		select {
		case square := <-c1:
			fmt.Printf("Square: %d\n", square)
		case cube := <-c2:
			fmt.Printf("Cube: %d\n", cube)
		}
	}
	
	fmt.Println("\n=== QUIT CHANNEL PATTERN ===")
	quit := make(chan bool)
	work := make(chan int)
	
	go func() {
		for {
			select {
			case job := <-work:
				fmt.Printf("Processing job: %d\n", job)
			case <-quit:
				fmt.Println("Worker stopping...")
				return
			}
		}
	}()
	
	for i := 1; i <= 3; i++ {
		work <- i
		time.Sleep(100 * time.Millisecond)
	}
	
	quit <- true
	time.Sleep(100 * time.Millisecond)
}
