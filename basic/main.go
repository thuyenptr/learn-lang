//package main
//import "./public"
//func main() {
//	//var arr1 = [3]int{1, 3, 5}
//	//arr2 := [...]int{1, 2, 3, 4, 5}
//	//fmt.Println(arr2)
//	//fmt.Println(arr1)
//	//
//	//var arr3 [5]int
//	//
//	//i := 0
//	//
//	//for i < 5 {
//	//	fmt.Println("Enter: ")
//	//	_, _ = fmt.Scan(&arr3[i])
//	//	fmt.Println(arr3[i])
//	//	i++
//	//}
//
//	rect := public.Rect{Width: 4, Height: 5}
//	sq := public.Square{Edge: 5}
//
//	shape := public.Shape(rect)
//
//	shape = public.Shape(sq)
//
//	println(shape.Area())
//
//}


package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	fmt.Println("start fibonacci function")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}


func main() {
	//c := make(chan int)
	//quit := make(chan int)
	//
	//go func() {
	//	fmt.Println("start anonymous func")
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//
	//fibonacci(c, quit)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("3s out")
	case <-time.After(3 * time.Second):
		fmt.Println("3s out 2")
	}
}
