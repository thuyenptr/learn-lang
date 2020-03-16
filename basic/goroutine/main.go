package main

import (
	"fmt"
	"sync"
)

var (
	counter int64
	mutex sync.Mutex
)

func main() {

	for i:=0; i<100;i++ {
		go func() {
			mutex.Lock()
			for j:=0; j<100;j++ {
				counter++
			}
			mutex.Unlock()
		} ()
	}
	var c string
	fmt.Scanln(&c)
	fmt.Println(counter)
}