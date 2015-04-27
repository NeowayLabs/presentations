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
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Nanosecond * amt)
		i++
	}
}
func main() {
	f(0)
	f(1)

	//	time.Sleep(time.Second)
}
