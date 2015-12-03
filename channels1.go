package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)

	go func() {
		ch1 <- "Hello"
		ch1 <- "World!"
	}()

	go func() {
		for {
			fmt.Println(<-ch1)
		}
	}()

	time.Sleep(1 * time.Second)
}
