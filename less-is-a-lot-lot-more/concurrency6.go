package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	var i int

	for {
		fmt.Println(n, ":", i)
		i++
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go f(0)
	go f(1)

	time.Sleep(time.Second)
}
