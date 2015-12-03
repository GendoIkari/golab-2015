package main

import "time"

func main() {
	channel := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 5; i++ {
		<-channel
	}
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 5; i++ {
		<-channel
	}
}
