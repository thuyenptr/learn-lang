package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float32
	perimeter() float32
}

type Square struct {
	edge float32
}

type Rect struct {
	width, height float32
}

func (sq Square) area() float32 {
	return float32(math.Pow(float64(sq.edge), 2.0))
}

func (sq Square) perimeter() float32 {
	return 4*sq.edge
}

func (rect Rect) area() float32 {
	return rect.width*rect.height
}

func (rect Rect) perimeter() float32 {
	return 2*(rect.height+rect.width)
}

func main() {
	rect := Rect{5,4}
	sq := Square{5}

	shape := Shape(rect)
	fmt.Println(shape.area())

	shape = Shape(sq)
	fmt.Println(shape.area())
}
