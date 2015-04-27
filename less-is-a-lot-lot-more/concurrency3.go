package main

import (
	"fmt"
	"time"
)

func f(n int) {
	var i int

	for {
		fmt.Println(n, ":", i)
		i++
	}
}

func main() {
	go f(0)
	go f(1)

	time.Sleep(time.Second)
}
