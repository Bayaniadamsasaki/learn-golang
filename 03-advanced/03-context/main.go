// Context: manajemen lifecycle, cancelation, dan deadline untuk goroutines
package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context, name string) {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: Task cancelled at step %d. Reason: %v\n", name, i, ctx.Err())
			return
		default:
			fmt.Printf("%s: Step %d\n", name, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("%s: Task completed successfully\n", name)
}

func fetchData(ctx context.Context, id int) (string, error) {
	select {
	case <-time.After(2 * time.Second):
		return fmt.Sprintf("Data for ID %d", id), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func processWithTimeout(ctx context.Context, duration time.Duration) {
	timeoutCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Process completed")
	case <-timeoutCtx.Done():
		fmt.Printf("Process timed out: %v\n", timeoutCtx.Err())
	}
}

func processWithDeadline(ctx context.Context, deadline time.Time) {
	deadlineCtx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()
	
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Deadline process completed")
	case <-deadlineCtx.Done():
		fmt.Printf("Deadline process cancelled: %v\n", deadlineCtx.Err())
	}
}

func processWithValue(ctx context.Context) {
	userID := ctx.Value("userID")
	requestID := ctx.Value("requestID")
	
	fmt.Printf("Processing request %v for user %v\n", requestID, userID)
	
	if userID == nil {
		fmt.Println("No user ID found in context")
		return
	}
	
	if uid, ok := userID.(string); ok {
		fmt.Printf("User ID as string: %s\n", uid)
	}
}

type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

func middlewareExample(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, userIDKey, "user123")
	ctx = context.WithValue(ctx, requestIDKey, "req456")
	return ctx
}

func main() {
	fmt.Println("=== BASIC CONTEXT ===")
	ctx := context.Background()
	go longRunningTask(ctx, "Task1")
	time.Sleep(3 * time.Second)
	
	fmt.Println("\n=== CONTEXT WITH CANCEL ===")
	cancelCtx, cancel := context.WithCancel(context.Background())
	
	go longRunningTask(cancelCtx, "CancellableTask")
	
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling task...")
	cancel()
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n=== CONTEXT WITH TIMEOUT ===")
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer timeoutCancel()
	
	data, err := fetchData(timeoutCtx, 1)
	if err != nil {
		fmt.Printf("Fetch failed: %v\n", err)
	} else {
		fmt.Printf("Fetched: %s\n", data)
	}
	
	fmt.Println("\n=== CONTEXT WITH DEADLINE ===")
	deadline := time.Now().Add(1500 * time.Millisecond)
	deadlineCtx, deadlineCancel := context.WithDeadline(context.Background(), deadline)
	defer deadlineCancel()
	
	data2, err2 := fetchData(deadlineCtx, 2)
	if err2 != nil {
		fmt.Printf("Fetch with deadline failed: %v\n", err2)
	} else {
		fmt.Printf("Fetched with deadline: %s\n", data2)
	}
	
	fmt.Println("\n=== CONTEXT WITH VALUES ===")
	valueCtx := context.WithValue(context.Background(), "userID", "john123")
	valueCtx = context.WithValue(valueCtx, "requestID", "req789")
	
	processWithValue(valueCtx)
	
	fmt.Println("\n=== TYPED CONTEXT KEYS ===")
	typedCtx := middlewareExample(context.Background())
	
	if userID := typedCtx.Value(userIDKey); userID != nil {
		fmt.Printf("Typed User ID: %s\n", userID)
	}
	
	if requestID := typedCtx.Value(requestIDKey); requestID != nil {
		fmt.Printf("Typed Request ID: %s\n", requestID)
	}
	
	fmt.Println("\n=== COMBINING CONTEXTS ===")
	parentCtx := context.Background()
	
	timeoutCtx2, timeoutCancel2 := context.WithTimeout(parentCtx, 2*time.Second)
	defer timeoutCancel2()
	
	valueCtx2 := context.WithValue(timeoutCtx2, userIDKey, "user456")
	
	cancelCtx2, cancel2 := context.WithCancel(valueCtx2)
	defer cancel2()
	
	go func() {
		time.Sleep(1 * time.Second)
		cancel2()
	}()
	
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Combined context completed")
	case <-cancelCtx2.Done():
		fmt.Printf("Combined context cancelled: %v\n", cancelCtx2.Err())
		if userID := cancelCtx2.Value(userIDKey); userID != nil {
			fmt.Printf("User ID from cancelled context: %s\n", userID)
		}
	}
	
	fmt.Println("\n=== TIMEOUT VS DEADLINE ===")
	fmt.Println("Testing timeout:")
	processWithTimeout(context.Background(), 1*time.Second)
	
	fmt.Println("Testing deadline:")
	processWithDeadline(context.Background(), time.Now().Add(1*time.Second))
	
	fmt.Println("\n=== CONTEXT PROPAGATION ===")
	rootCtx := context.Background()
	parentCtx2, parentCancel := context.WithCancel(rootCtx)
	defer parentCancel()
	
	childCtx, childCancel := context.WithTimeout(parentCtx2, 5*time.Second)
	defer childCancel()
	
	go longRunningTask(childCtx, "ChildTask")
	
	time.Sleep(1 * time.Second)
	fmt.Println("Cancelling parent context...")
	parentCancel()
	time.Sleep(1 * time.Second)
}
