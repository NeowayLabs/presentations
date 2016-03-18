package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Getter interface {
	Get() uint64
}

type Binary uint64

func (i Binary) Get() uint64 {
	return uint64(i)
}

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func main() {
	b := Binary(123)
	s := Stringer(b)
	s = Binary(999)
	fmt.Println(s)
	fmt.Println(b)
	g := Getter(b)
	g = Binary(999)
	fmt.Println(g)
	fmt.Println(b)
}
