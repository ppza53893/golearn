package main

import (
	"fmt"
	"time"
)

func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Println("funcA")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Println("main")
		time.Sleep(100 * time.Millisecond)
	}
}
