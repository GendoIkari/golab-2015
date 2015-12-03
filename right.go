package main

import "math/rand"

func main() {
	inputs := make(chan int)
	outputs := make(chan int)

	go func() {
		for {
			// ...
			inputs <- rand.Int()
			// ...
		}
	}()

	go func() {
		for {
			// ...
			input := <-inputs
			outputs <- input * 2
			// ...
		}
	}()
}
