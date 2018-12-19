package linked_list

import (
	"fmt"
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	list := NewLinkedList()
	log.Println(list)
	list.Append(1)
	list.Append(2)
	log.Println(list)
	list.Prepend(3)
	log.Println(list)
	log.Println(list.Find(3))
	log.Println(list)

	list.ForEach(func(value interface{}) {
		fmt.Println("- ", value)
	})
}
