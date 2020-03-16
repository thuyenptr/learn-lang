package main

import (
	"./common"
	"fmt"
)

func main() {
	var st common.Student
	st.ID = 5
	st.Name = "bill"
	st.Age = 18
	fmt.Println(st)
	st1 := common.Student{1, "john", 20}
	fmt.Println(st1)
}
