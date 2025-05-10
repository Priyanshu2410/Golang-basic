package main

import (
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic Goroutine
func basicGoroutine() {
	// Goroutine is a lightweight thread managed by the Go runtime
	go func() {
		fmt.Println("Hello from goroutine!")
	}()
	time.Sleep(time.Millisecond) // Give goroutine time to execute
}

// Example 2: Multiple Goroutines with WaitGroup
func multipleGoroutines() {
	var wg sync.WaitGroup

	// Launch 3 goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment counter
		go func(id int) {
			defer wg.Done() // Decrement counter when done
			fmt.Printf("Goroutine %d executing\n", id)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
}

// Example 3: Channel Communication
func channelCommunication() {
	// Create a channel for communication
	ch := make(chan string)

	// Sender goroutine
	go func() {
		ch <- "Message from goroutine" // Send data
	}()

	msg := <-ch // Receive data
	fmt.Println(msg)
}

// Example 4: Buffered Channels
func bufferedChannels() {
	// Create buffered channel with capacity 2
	ch := make(chan int, 2)

	// Send values without blocking
	ch <- 1
	ch <- 2

	fmt.Println(<-ch) // Read values
	fmt.Println(<-ch)
}

// Example 5: Select Statement
func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Select helps handle multiple channel operations
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(time.Second):
		fmt.Println("Timeout!")
	}
}

// Example 6: Mutex for Shared Resource
func mutexExample() {
	var mutex sync.Mutex
	counter := 0

	// Launch 5 goroutines
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock() // Lock before accessing shared resource
			counter++
			mutex.Unlock() // Unlock after modification
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)
}

func main() {
	fmt.Println("=== Basic Goroutine ===")
	basicGoroutine()

	fmt.Println("\n=== Multiple Goroutines ===")
	multipleGoroutines()

	fmt.Println("\n=== Channel Communication ===")
	channelCommunication()

	fmt.Println("\n=== Buffered Channels ===")
	bufferedChannels()

	fmt.Println("\n=== Select Statement ===")
	selectStatement()

	fmt.Println("\n=== Mutex Example ===")
	mutexExample()
}
