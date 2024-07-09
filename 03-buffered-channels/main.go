package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	// A channel deadlock occurs when all goroutines are blocked, waiting for something to happen, but nothing can progress. This often happens due to improper use of channels. Here are some key points:
	// Basic concept:
	// Channels are used for communication between goroutines.
	//	A deadlock happens when goroutines are stuck, unable to proceed.

	// To prevent we can use buffered channels.

	channel <- "First message"
	channel <- "Second message"
	//go func() {
	//	channel <- "First message"
	//}()

	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

// Channels are first in first out queue.
// can we loop through channel?
// Yes... but... there are many things that we need to take care
