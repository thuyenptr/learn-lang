package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)


func sleepRandom(funcName string, ch chan int) {
	defer func() {
		fmt.Println("sleepRandom complete")
	}()

	fmt.Println("function call sleep ", funcName)

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	sleepTime := r.Intn(100) + 100

	fmt.Println("start sleep ", sleepTime, " ms")
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	fmt.Println("wake up")

	if ch != nil {
		ch <- sleepTime
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	sleeptimeChan := make(chan int)

	go sleepRandom("sleepRandomContext", sleeptimeChan)

	select {
		case <- ctx.Done():
			fmt.Println("sleepRandomContext return")
		case sleeptime:= <- sleeptimeChan:
			fmt.Println("sleepRandomContext sleep for ", sleeptime, " ms")
	}
}

func doWorkContext(ctx context.Context) {

	timeoutContext, cancelFunction := context.WithTimeout(ctx, time.Duration(300)* time.Millisecond)

	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	check := make(chan bool)

	go sleepRandomContext(timeoutContext, check)

	select {
		case <- ctx.Done():
			fmt.Println("doWorkContext return")
		case complete := <- check:
			fmt.Println("doWorkContext return ", complete)
	}
}

func main() {
	ctx := context.Background()

	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	go func() {
		sleepRandom("Main", nil)
		cancelFunction()
	}()

	doWorkContext(ctxWithCancel)
}
