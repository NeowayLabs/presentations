package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(name string, c chan string) {
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("%s: %d", name, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	go f("func1", c)
	go f("func2", c)

	for i := 0; i < 20; i++ {
		fmt.Printf("Goroutine say: %q\n", <-c) // Receive expression is just a value.
	}
} //
