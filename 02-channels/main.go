package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println("time elapsed: ", time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Tommy"
	go attackTarget(evilNinja, smokeSignal)
	fmt.Println(<-smokeSignal)
}

func attackTarget(target string, attacked chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Throwing ninja stars at", target)
	attacked <- true
}
