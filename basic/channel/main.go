package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	share := make(chan int)
	wg.Add(2)
	go play("bill", share)
	share <- 1
	go play("gates", share)


	wg.Wait()
}

func play(name string, share chan int) {
	defer wg.Done()

	for {
		value, ok := <- share
		if !ok {
			log.Println(name, " finished")
			return
		}

		log.Println(name, " get value ", value)

		value++
		share <- value
		//time.Sleep(2 * time.Second)
	}
}
