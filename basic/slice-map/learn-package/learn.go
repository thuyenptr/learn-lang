package learn

import (
	"fmt"
	"os"
)

func Learn() {
	fmt.Println("Learning")

	sum, prod := function(2, 5)

	fmt.Println("sum: ", sum, " prod: ", prod)

	fmt.Println("Flow control")

	learnFlowControl()

	fmt.Println(fact(5))

	f := square()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	list := anonymousfunction()

	for _, f := range list {
		f()
	}

	deferfunction()
}

func function(a, b int) (sum, prod int) {
	sum = a + b
	prod = a * b
	return
}

func learnFlowControl() {
	if true {
		fmt.Println("true")
	}

	if false {

	} else {
		fmt.Println("else statement")
	}

	x := 0

	switch x {
	case 0:
		fmt.Println("x equal to 0")
	case 1:
		fmt.Println("x equal to 1")
	case 42:
		fmt.Println("x equal to 42")

	case 43:

	default:
		fmt.Println("default reached")
	}
}

func fact(n uint) uint {
	var rs, i uint = 1, 1
	for {
		rs *= i
		i++
		if i > n {
			break
		}
	}
	return rs
}

func square() func() int {
	var x int // default la 0
	return func() int {
		x++
		return x * x
	}
}

func anonymousfunction() []func() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	var listfunc []func()
	for _, v := range a {
		tmp := v
		listfunc = append(listfunc, func() {
			fmt.Println("In function: ", tmp)
		})
	}
	return listfunc
}

func createFile(fname string) (file *os.File) {
	fmt.Println("Create file")
	file, err := os.Create(fname)
	if err != nil {
		panic(err)
	}

	return
}

func closeFile(file *os.File) {
	fmt.Println("Close file")
	_ = file.Close()
}

func writingFile(file *os.File) {
	fmt.Println("Writing to file")
	_,_ = fmt.Fprintln(file, "Hello")
}

func deferfunction() {
	file := createFile("defer.txt")
	defer closeFile(file)
	writingFile(file)
}
