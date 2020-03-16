package main

import (
	"log"
	"sync"
	"time"
)

const (
	numTask = 10
	numWorker = 4
)

var (
	wg sync.WaitGroup
)

func main() {
	tasks := make(chan int, 5)
	wg.Add(numWorker)

	for i:=0; i<numWorker; i++ {
		go dowork(tasks, i)
	}

	for j:=0; j<numTask; j++ {
		tasks <- j
	}

	close(tasks)
	wg.Wait()
}

func dowork(tasks chan int, workerId int) {
	defer wg.Done()
	for {
		task, ok := <-tasks

		if !ok {
			log.Println("Complete work")
			return
		}

		log.Println(workerId, " do task ", task)
		time.Sleep(2*time.Second)
		log.Println(workerId, " finish task ", task)
	}

}


