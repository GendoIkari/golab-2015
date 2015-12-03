package main

import "fmt"

func main() {

	data := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			data <- i
		}
		close(data)
	}()

	sum := 0

	for i := range data {
		sum += i
	}

	fmt.Println(sum)
}
