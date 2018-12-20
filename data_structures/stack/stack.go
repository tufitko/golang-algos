package stack

import "github.com/tufitko/golang-algos/data_structures/linked_list"

type Stack struct {
	list *linked_list.LinkedList
}

func NewStack() *Stack {
	return &Stack{list: linked_list.NewLinkedList()}
}

func (stack *Stack) IsEmpty() bool {
	if stack == nil || stack.list == nil || stack.list.Len() == 0 {
		return true
	}
	return false
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	return stack.list.Get(0)
}

func (stack *Stack) Push(value interface{}) {
	stack.list.Prepend(value)
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic("stack is empty")
	}

	return stack.list.Delete(0)
}

func (stack *Stack) String() string {
	return stack.list.String()
}
