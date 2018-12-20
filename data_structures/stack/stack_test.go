package stack

import (
	"log"
	"testing"
)

func TestSeimple(t *testing.T) {
	stack := NewStack()

	for i := 0; i < 4; i++ {
		stack.Push(i)
		log.Println(stack)
	}

	log.Println(stack.Peek())

	for !stack.IsEmpty() {
		log.Println(stack.Pop())
	}
}
