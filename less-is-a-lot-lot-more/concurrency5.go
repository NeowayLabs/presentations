package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int, c chan string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("%d: %d", n, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go f(0, c)
	go f(1, c)

	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine say: %q\n", <-c) // Receive expression is just a value.
	}
} //
