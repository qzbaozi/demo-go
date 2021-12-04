package main

import (
	"fmt"
	"time"
)

func main() {
	ticker()

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

}

func ticker() {
	ticker := time.NewTicker(1e9)

	defer ticker.Stop()
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println(" .")
			time.Sleep(5e7)
		}
	}
}
