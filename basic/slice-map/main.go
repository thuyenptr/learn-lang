package main

import (
	"./learn-package"
	mmap "./map"
	"fmt"
)

func main() {
	/*
		var arr1 []int
		var arr2 []int = []int{1, 2, 3, 4, 5}
		fmt.Println(arr1)
		fmt.Println(arr2)

		arr3 := make([]int, 3)
		arr4 := []int{1: 1, 3: 2}
		arr5 := make([]int, 4, 5)
		arr5[3] = 9
	*/
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println("a: ", a)
	a1 := a[1:]
	fmt.Println("a1: ", a1)

	a2 := a[:len(a)-1]
	fmt.Println("a2: ", a2)
	a3 := a[:2]
	fmt.Println("a3: ", a3)
	a3 = append(a3, a[3:]...)
	fmt.Println("a3: ", a3)
	fmt.Println("a: ", a)
	fmt.Println("a2: ", a2)
	fmt.Println("a1: ", a1)

	mmap.Show()

	learn.Learn()
}
