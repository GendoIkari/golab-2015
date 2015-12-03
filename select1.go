package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producerSlow(output chan string) {
	for {
		output <- "Message!"
		time.Sleep(1 * time.Second)
	}
}

func producerFast(output chan int) {
	for {
		output <- rand.Int()
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go producerSlow(ch1)
	go producerFast(ch2)

	for {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
}
