package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	//ch := make(chan string,10)

	//ch <- "1"
	go sendData(ch)
	//go getData(ch)
	//time.Sleep(1e9)

	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	//time.Sleep(2e9)
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}
