package main

import "fmt"

func update(x int, y int, primes []int) {
	x = 10
	y = 10
	primes[0] = 10
}

func main() {
	var x, y int
	primes := [3]int{2, 3, 5}

	t := primes[0:len(primes)]

	update(x, y, t)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(primes)
}
