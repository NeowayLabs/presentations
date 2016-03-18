package main

import "fmt"

func update(x int, y int, primes [3]int) {
	x = 10
	y = 10
	primes[0] = 10
}

func main() {
	var x, y int
	primes := [3]int{2, 3, 5}

	update(x, y, primes)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(primes)
}
