package wp

import "fmt"

const wokersCount = 10

type Worker struct {
	id 			  int
	messagePassed int
}

var wokers = []*Worker {
	{id: 1},
	{id: 2},
	{id: 3},
	{id: 4},
	{id: 5},
	{id: 6},
	{id: 7},
	{id: 8},
	{id: 9},
	{id: 10},
}

type Pool[Data any] struct {
	pool 	chan *Worker
	handler func(int, Data)
}

func New[Data any](handler func(int, Data)) *Pool[Data] {
	return &Pool[Data] {
		pool: 	 make(chan *Worker, wokersCount),
		handler: handler,
	}
}

func (p *Pool[Data]) Create() {
	for _, w := range wokers {
		p.pool <-w
	}
}

func (p *Pool[Data]) Handle(data Data) {
	w := <-p.pool
	go func() {
		p.handler(w.id, data)
		w.messagePassed++
		p.pool <- w
	}()
}

func (p *Pool[Data]) Wait() {
	for range len(wokers) {
		<-p.pool
	}
}

func (p *Pool[Data]) Stats() {
	fmt.Println("_____RESULT_____")
	for _, w := range wokers {
		fmt.Printf("Worker: %d, message passed: %d\n", w.id, w.messagePassed )
	}
}