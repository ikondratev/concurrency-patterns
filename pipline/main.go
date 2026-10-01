package main

import (
	"context"
	"fmt"
	"time"
)

func generator(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range 10 {
			select {
			case <-ctx.Done():
				return
			case out <- i + 1:
			}
		}
	}()

	return out
}

func doubler(ctx context.Context, input <-chan int) <-chan int {
	out := make(chan int) 

	go func() {
		defer close(out)
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
				case <-time.After(300 * time.Millisecond):
				}
				select {
				case <-ctx.Done():
					return
				case out <- v * 2:
				}
			}
		}
	}()

	return out
}

func reader(ctx context.Context, input <-chan int) {
	for i := range input {
		fmt.Println(i)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-input:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}

func main() {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	reader(ctx, doubler(ctx, generator(ctx)))
	
	fmt.Println("Duration:", time.Since(start))
}