package Exercise1

import (
	"fmt"
	"sync"
)

func main(){
	// Create an unbuffered channel.
	share := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch two goroutines.
	go func() {
		goroutine("Bill", share)
		wg.Done()
	}()

	go func() {
		goroutine("Joan", share)
		wg.Done()
	}()

	// Start the sharing.
	share <- 1

	// Wait for the program to finish.
	wg.Wait()
}

func goroutine(name string, share chan int){
	for {
		value,ok := <- share
		if !ok {
			// If the channel was closed, return.
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}
		// Display the value.
		fmt.Printf("Goroutine %s Inc %d\n", name, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s Down\n", name)
			return
		}

		// Increment the value and send it
		// over the channel.
		share <- (value + 1)
	}
}
