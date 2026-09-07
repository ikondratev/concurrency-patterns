package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	maxJobs = 50
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
			case <-time.After(900 * time.Millisecond):
			}
			select {
			case <-ctx.Done():
				return
			case results <- v * 3:
				fmt.Printf("worker %d passed the job %d\n", id, v)
			}
		}
	}
}

func main() {
	jobs 	:= make(chan int, maxJobs)
	results := make(chan int, maxJobs)
	start 	:= time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	
	var wg sync.WaitGroup
	wg.Add(maxWorkers)

	for w := range maxWorkers {
		go func() {
			defer wg.Done()
			worker(ctx, w, jobs, results)
		}()
	}

	go func() {
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

	fmt.Println("Score:", time.Since(start))
}