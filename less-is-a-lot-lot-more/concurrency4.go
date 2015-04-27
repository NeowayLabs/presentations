package main

import (
	"fmt"
	"runtime"
	"time"
)

func f(n int) {
	var i int

	for {
		fmt.Println(n, ":", i)
		i++

		runtime.Gosched()
	}
}

func main() {
	go f(0)
	go f(1)

	time.Sleep(time.Second)
}
