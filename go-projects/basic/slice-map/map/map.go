package mapdemo

import "fmt"

func Show() {
	fmt.Println("hello")

	var student map[int]string
	fmt.Println(student)

	student = make(map[int]string)

	student[1] = "Mai"
	student[2] = "Lan"
	student[3] = "Cuc"
	student[4] = "Truc"

	fmt.Println(student)

	delete(student, 4)

	fmt.Println(student)
}
