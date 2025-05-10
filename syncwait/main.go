package main

import (
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic WaitGroup usage
func basicWaitGroupExample() {
	var wg sync.WaitGroup

	// Launch 3 goroutines and wait for them to complete
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter
		go func(id int) {
			defer wg.Done() // Decrement counter when goroutine completes
			fmt.Printf("Goroutine %d starting\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Println("All goroutines completed")
}

// Example 2: WaitGroup with channels
func waitGroupWithChannels() {
	var wg sync.WaitGroup
	results := make(chan int, 3)

	// Launch workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Simulate work
			time.Sleep(time.Duration(id) * time.Second)
			results <- id * 2
		}(i)
	}

	// Close channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}
}

// Example 3: WaitGroup with error handling
func waitGroupWithErrors() {
	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	// Launch workers that might return errors
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if id == 2 {
				errChan <- fmt.Errorf("error in worker %d", id)
				return
			}
			fmt.Printf("Worker %d completed successfully\n", id)
		}(i)
	}

	// Wait for all workers and check errors
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Handle errors
	for err := range errChan {
		if err != nil {
			fmt.Printf("Error occurred: %v\n", err)
		}
	}
}

func main() {
	fmt.Println("=== Basic WaitGroup Example ===")
	basicWaitGroupExample()

	fmt.Println("\n=== WaitGroup with Channels Example ===")
	waitGroupWithChannels()

	fmt.Println("\n=== WaitGroup with Error Handling Example ===")
	waitGroupWithErrors()
}
