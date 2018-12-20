package queue

import "github.com/tufitko/golang-algos/data_structures/linked_list"

type Queue struct {
	list *linked_list.LinkedList
}

func NewQueue() *Queue {
	return &Queue{list: linked_list.NewLinkedList()}
}

func (queue *Queue) IsEmpty() bool {
	if queue == nil || queue.list == nil || queue.list.Len() == 0 {
		return true
	}
	return false
}

func (queue *Queue) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	return queue.list.Get(0)
}

func (queue *Queue) Enqueue(value interface{}) {
	queue.list.Append(value)
}

func (queue *Queue) Dequeue() interface{} {
	if queue.IsEmpty() {
		panic("queue is empty")
	}

	return queue.list.Delete(0)
}

func (queue *Queue) String() string {
	return queue.list.String()
}
