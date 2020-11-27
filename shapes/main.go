package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLen float64
}

func main() {

}

func (s square) getArea() float64 {
	return s.sideLen * s.sideLen
}

func (t triangle) getArea() float64 {
	return t.height * t.base * .5
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
