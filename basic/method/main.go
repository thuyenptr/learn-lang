package main

import (
	"fmt"
	"image/color"
)

type Student struct {
	name string
	id int
}

func (std Student) show() {
	fmt.Printf("student with id %d and name %s", std.id, std.name)
}

type Rect struct {
	width, heigh float32
}

func (rect Rect) Area() float32{
	return rect.heigh*rect.width
}

type RectWidthColor struct {
	Rect
	color color.RGBA
}

func main() {
	student := Student{"bill", 10}
	student.show()

	rect := new(RectWidthColor)
	rect.width = 5
	rect.heigh = 5
	fmt.Println("\n", rect.Area())
}
