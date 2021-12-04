package main

import (
	"fmt"
	"time"
)

func main() {
	suck2(pump2())
	time.Sleep(1e9)
}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck2(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
