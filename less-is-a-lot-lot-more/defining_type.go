package main

import "log"

type Point struct {
	x int
	y int
}

func (self *Point) getLocation() (int, int) {
	return self.x, self.y
}
func (p *Point) Move(x int, y int) {
	p.x = x
	p.y = y
}
func main() {
	myPoint := Point{x: 1, y: 2}
	x, y := myPoint.getLocation()
	log.Printf("x[%d] y[%d]", x, y)
	myPoint.Move(5, 7)
	x, y = myPoint.getLocation()
	log.Printf("x[%d] y[%d]", x, y)
}
