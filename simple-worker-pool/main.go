package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ( 
	jobsCount = 30
	workersCount = 5
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker: %d, start with job: %d\n", id, j)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker: %d, end with job: %d\n", id, j)
		result <- rand.Intn(j + 100)
	}
}

func main() {
	jobs := make(chan int, jobsCount)
	results := make(chan int, jobsCount)

	for w := range workersCount {
		go worker(w, jobs, results)
	}

	for j := range jobsCount {
		jobs <- j
	}
	close(jobs)

	for range jobsCount {
		<-results
	}
}