package main

import "fmt"

func main() {
	a := 5
	b := &a
	a++

	fmt.Println(*b)

	var p = new(int)

	*p = 100

	fmt.Println(*p)
}
