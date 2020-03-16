package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg sync.WaitGroup) {
	fmt.Println("Waiting for signal")
	select {
	case <-ctx.Done():
		fmt.Println("Worker stop")
		wg.Done()
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, wg)

	time.Sleep(2 * time.Second)
	ctx.Done()
	time.Sleep(1 * time.Second)
}
