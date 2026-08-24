package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workersCount = 5
	jobsCount = 30
)

func main() {
	jobs := make([]int, 0, jobsCount)
	for j := range jobsCount {
		jobs = append(jobs, j)
	}

	jobsCh := make(chan int, jobsCount)
	go func() {
		defer close(jobsCh)
		for _, j := range jobs {
			jobsCh <- j
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(workersCount)
	for w := range workersCount {
		go func() {
			defer wg.Done()
			for j := range jobsCh {
				fmt.Printf("Woker: %d, start processing job: %d\n", w, j)
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Woker: %d, stop processing job: %d\n", w, j)
			}
		}()
	}

	wg.Wait()
}