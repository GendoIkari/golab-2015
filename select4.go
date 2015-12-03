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
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case num := <-ch2:
			fmt.Println(num)
		default:
			fmt.Println("Missing data!")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
