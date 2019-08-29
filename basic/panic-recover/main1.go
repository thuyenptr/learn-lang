package main

import (
	"fmt"
	"log"
)

func foo(i int) {
	defer func() {
		log.Println("defer call")
		if r := recover(); r != nil {
			log.Println("recover in foo ", r)
		}
	}()

	i++
	if i > 4 {
		log.Println("do panic function")
		panic(fmt.Errorf("i = 4"))
	}
}

func main() {
	foo(5)
}
