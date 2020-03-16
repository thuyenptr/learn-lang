package main

import (
	"fmt"
)



func main() {


	c1 := make(chan int)
	c2 := make(chan int)


	go func() {
		//c1 <- 5
		fmt.Println("length: ", cap(c1))
	}()

	go func() {
		//c2 <- 10
		fmt.Println("length: ", cap(c2))
	}()

	select {
	case <-c1:
		fmt.Println("channel c1")
	case <-c2:
		fmt.Println("channel c2")
	default:
		fmt.Println("default")
	}

}
