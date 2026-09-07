package main

import (
	"context"
	"slices"
	"testing"
)

func TestGenerator(t *testing.T) {
	ctx := context.Background()
	out := generator(ctx)

	var got []int
	for v := range out {
		got = append(got, v)
	}

	want := []int{1,2,3,4,5,6,7,8,9,10}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want: %v", got, want)
	}
}

func TestDoubler(t *testing.T) {
	ctx := context.Background()
	ch := make(chan int, 3) 
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	out := doubler(ctx, ch)

	var got []int
	for v := range out {
		got = append(got, v)
	}

	want := []int{2,4,6}
	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestPipiline(t *testing.T) {
	ctx := context.Background()
	out := doubler(ctx, generator(ctx))

	var got []int
	for i := range out {
		got = append(got, i)
	}

	want := []int{2,4,6,8,10,12,14,16,18,20}
	if !slices.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}