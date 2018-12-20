package linked_list

import (
	"errors"
	"fmt"
)

type Element struct {
	next, prev *Element
	Value      interface{}
}

type LinkedList struct {
	head, tail *Element
	len        int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{len: 0}
}

func (list *LinkedList) Len() int {
	return list.len
}

func (list *LinkedList) Append(value interface{}) {
	element := Element{Value: value}
	if list.head == nil {
		list.head = &element
		list.tail = &element
	} else {
		list.tail.next = &element
		element.prev = list.tail
		list.tail = &element
	}

	list.len++
}

func (list *LinkedList) Prepend(value interface{}) {
	element := Element{Value: value}
	if list.head == nil {
		list.head = &element
		list.tail = &element
	} else {
		list.head.prev = &element
		element.next = list.head
		list.head = &element
	}

	list.len++
}

func (list *LinkedList) Get(index int) interface{} {
	if index >= list.len || index < 0 {
		panic(errors.New("out of range"))
	}

	temp := list.head

	for temp != nil {
		if index == 0 {
			return temp.Value
		}

		temp = temp.next
		index--
	}

	panic(errors.New("uncaught error"))
}

func (list *LinkedList) String() string {
	values := make([]interface{}, 0, list.len+2)
	values = append(values, "[")
	temp := list.head
	for temp != nil {
		values = append(values, temp.Value)
		temp = temp.next
	}

	values = append(values, "]")
	return fmt.Sprint(values...)
}

func (list *LinkedList) Delete(index int) interface{} {
	if index >= list.len || index < 0 {
		panic(errors.New("out of range"))
	}

	// delete alone node :(
	if list.len == 0 {
		value := list.head.Value

		list.head = nil
		list.tail = nil
		list.len = 0

		return value
	}

	//delete first node
	if index == 0 {
		value := list.head.Value

		list.head = list.head.next
		list.head.prev = nil
		list.len--

		return value
	}

	//delete last node
	if index == list.len-1 {
		value := list.tail.Value

		list.tail = list.tail.prev
		list.tail.next = nil
		list.len--

		return value
	}

	// delete middle node
	temp := list.head
	for temp != nil {
		if index == 0 {

			temp.prev.next = temp.next
			temp.next.prev = temp.prev
			list.len--
			return temp.Value
		}

		temp = temp.next
		index--
	}

	panic(errors.New("uncaught error"))
}

func (list *LinkedList) Find(value interface{}) int {
	temp := list.head

	index := 0
	for temp != nil {
		if temp.Value == value {
			return index
		}

		temp = temp.next
		index++
	}

	return -1
}

func (list *LinkedList) ForEach(fn func(value interface{})) {
	temp := list.head

	for temp != nil {
		fn(temp.Value)

		temp = temp.next
	}
}

func (list *LinkedList) Reverse() {
	list.head, list.tail = list.tail, list.head

	temp := list.head

	for temp != nil {
		temp.prev, temp.next = temp.next, temp.prev

		temp = temp.next
	}
}
