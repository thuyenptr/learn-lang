package public

import "math"

type Shape interface {
	Area() float32
	perimeter() float32
}

type Square struct {
	Edge float32
}

type Rect struct {
	Width, Height float32
}

func (sq Square) Area() float32 {
	return float32(math.Pow(float64(sq.Edge), 2.0))
}

func (sq Square) perimeter() float32 {
	return 4*sq.Edge
}

func (rect Rect) Area() float32 {
	return rect.Width*rect.Height
}

func (rect Rect) perimeter() float32 {
	return 2*(rect.Height+rect.Width)
}