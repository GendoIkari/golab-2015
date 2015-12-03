package main

import "math/rand"

func main() {
	sharedInput := 0
	sharedOutput := 0

	go func() {
		for {
			// ...
			sharedInput = rand.Int()
			// ...
		}
	}()

	go func() {
		for {
			// ...
			sharedOutput = sharedInput * 2
			// ...
		}
	}()
}
