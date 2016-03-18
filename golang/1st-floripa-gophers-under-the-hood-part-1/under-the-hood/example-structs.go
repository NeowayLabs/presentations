package main

import "fmt"

type Point struct{ X, Y int }
type Rect1 struct{ Min, Max Point }
type Rect2 struct{ Min, Max *Point }

func notUpdate(r Rect1) {
	r.Min.X = 10
}
func update(r Rect2) {
	r.Min.X = 10
}

func main() {
	r1 := Rect1{Point{0, 0}, Point{10, 10}}
	r2 := Rect2{&Point{0, 0}, &Point{10, 10}}

	notUpdate(r1)
	fmt.Println(r1)
	update(r2)
	fmt.Println(r2)
	fmt.Println(*r2.Min)
}
