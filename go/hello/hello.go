package main

import (
	"fmt"
	"github.com/learn-go/hello/stringutil"
)

func main() {
	fmt.Print(stringutil.Reverse("hello"))
	i := 5
	var index int
	fmt.Println(i)
	index = 4
	fmt.Println(index)

	var s string
	s = "hello"

	var uni = []rune(s)

	for i := 0; i < len(uni); i = i + 1 {
		fmt.Println(uni[i])
	}
}
