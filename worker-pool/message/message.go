package message

import (
	"math/rand"
)

const maxChunkSize = 10
var messageCounter = 0

type Message struct {
	Id 	  		  int
	MessagePassed int
	Title 		  string
}

func GetMessages() []Message {
	count := rand.Intn(maxChunkSize)
	messages := make([]Message, 0, count)

	for range count {
		messageCounter++
		messages = append(messages, Message{Id: messageCounter})
	}

	return messages
}