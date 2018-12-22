package heap

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	heap := NewHeap(func(first, second interface{}) bool {
		return first.(int) < second.(int)
	})

	heap.Add(1)
	heap.Add(2)
	heap.Add(4)
	heap.Add(3)
	heap.Add(6)
	heap.Add(5)
	heap.Add(4)
	heap.Add(1)

	log.Println(heap)

	heap.Remove(2)
	log.Println(heap)
}
