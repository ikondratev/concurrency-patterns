package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func writter(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range 10 {
			select {
			case <-ctx.Done():
				return
			case ch <- i + 1:
			}
		}
	}()
	return ch
}

func doubler(ctx context.Context, input <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return 
			case v, ok := <- input:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case <-time.After(time.Duration(rand.Intn(300)) * time.Millisecond):
				}
				select {
				case <-ctx.Done():
					return
				case ch <- v*2:
				}
			}
		}
	}()

	return ch
}

func reader(ctx context.Context, input <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <- input:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	reader(ctx, doubler(ctx, writter(ctx)))
}