package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const(
	maxJobs = 30
	maxWorkers = 5
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-jobs:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case <-time.After(500 * time.Millisecond):
			}
			select {
			case <-ctx.Done():
				return
			case results <- v * 3:
				fmt.Printf("Worker %d passed job %d\n", id, v)
			}

		}
	}
}

func main() {
	jobs := make(chan int, maxJobs)
	results := make(chan int, maxJobs)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(maxWorkers)
	for w := range maxWorkers {
		go func() {
			defer wg.Done()
			
			worker(ctx, w, jobs, results)
		}()
	}

	go func () {
		wg.Wait()
		close(results)
	}()

l:	
	for j := range maxJobs {
		select {
		case <-ctx.Done():
			break l
		case jobs <- j:
		}
	}
	close(jobs)

	for r := range results {
		fmt.Println("got:", r)
	}
}