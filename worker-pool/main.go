package main

import (
	"context"
	"fmt"
	"time"

	"github.com/kilila/worker-pool/message"
	"github.com/kilila/worker-pool/wp"
)

func processMessage(workerId int, msg message.Message) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Worker %d processed message %d\n", workerId, msg.Id)
}


type IPool interface {
	Create()
	Handle(message.Message)
	Wait()
	Stats()
}

func main() {
	var pool IPool
	pool = wp.New(processMessage)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

l:
	for {
		select {
		case <- ctx.Done():
			break l
		default:
		}

		messages := message.GetMessages()
		pool.Create()

		for _, m := range messages {
			pool.Handle(m)
		}

		pool.Wait()
	}

	pool.Stats()
}