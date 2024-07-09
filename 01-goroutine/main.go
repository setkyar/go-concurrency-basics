package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja)
	}

	// The reason why we have to sleep here is giving time for go routine to execute first
	// Main function execution: When you start goroutines, they run concurrently with the main function.
	// However, the main function doesn't wait for goroutines to complete by default.
	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(time.Second)
}
