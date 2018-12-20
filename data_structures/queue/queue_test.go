package queue

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	queue := NewQueue()

	for i := 0; i < 4; i++ {
		log.Println(queue)
		queue.Enqueue(i)
	}
	log.Println(queue.Peek())

	for !queue.IsEmpty() {
		log.Println(queue.Dequeue())
	}
}
