package main

import (
	"fmt"
	"math/rand"
	"time"
)

func search() int {
	// Random delay between 100ms and 2000ms
	// delayMs := 100 + rand.Intn(1901)
	// time.Sleep(time.Duration(delayMs) * time.Millisecond)
	return rand.Int()
}

/*
This method uses `select` in goroutines to wait for either a write on a channel or a read on done channel.
- It receives a num (number of searchers), launches that many goroutines, each running the searcher func.
- The fastest searcher response signals data to result (unbuffered channel).
- After receiving the first result, close the done channel to signal cancellation to others.
*/
func searchData(num int) int {
	done := make(chan struct{})
	result := make(chan int)
	for i := 0; i < num; i++ { // Fixed: Can't range over int; use standard for loop.
		go func(searcher func() int, index int) {
			fmt.Printf("Goroutine %d launched\n", index) // Fixed: Use index (passed param), add clarity.
			select {
			case result <- searcher(): // searcher() runs first (blocks on sleep), then tries to send.
				fmt.Printf("Goroutine %d sent result\n", index) // Clarified print.
			case <-done:
				//select only picks ready cases. If both ready, pseudo-random choice; if one, that's it.
				// Here, send can't be ready (no receiver), because after receiving the main program have moved to next process (close done)
				// so <-done always wins. Prints "canceled", exits.
				fmt.Printf("Goroutine %d canceled via done\n", index) // Clarified print.
			}
		}(search, i) // Pass search func and index.
	}
	r := <-result // Block until first send succeeds.
	fmt.Printf("Main received result: %d\n", r)
	close(done)                 // Signal cancellation to remaining goroutines.
	time.Sleep(2 * time.Second) // Pause to observe prints from canceled goroutines.
	return r
}

func main() {
	searchData(10)
}
