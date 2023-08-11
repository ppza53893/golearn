package main

import (
	"fmt"
	"time"
)

func funcACh(chA chan<- string) {
	time.Sleep(100 * time.Millisecond)
	chA <- "funcA"
}

func main() {
	chA := make(chan string) // create channel
	defer close(chA)         // close channel
	go funcACh(chA)          // call goroutine
	msg := <-chA             // receive message from channel
	fmt.Println(msg)         // print message
}
