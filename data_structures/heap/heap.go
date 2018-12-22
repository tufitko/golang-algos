package heap

import "fmt"

type Heap struct {
	heapContainer []interface{}
	less          func(first, second interface{}) bool
}

func NewHeap(less func(first, second interface{}) bool) *Heap {
	return &Heap{less: less}
}

func (heap *Heap) GetLeftChildIndex(parentIndex int) int {
	return (2 * parentIndex) + 1
}

func (heap *Heap) GetRightChildIndex(parentIndex int) int {
	return (2 * parentIndex) + 2
}

func (heap *Heap) GetParentIndex(childIndex int) int {
	return childIndex / 2
}

func (heap *Heap) HasParent(childIndex int) bool {
	return childIndex > 0
}

func (heap *Heap) HasLeftChild(parentIndex int) bool {
	return heap.GetLeftChildIndex(parentIndex) < len(heap.heapContainer)
}

func (heap *Heap) HasRightChild(parentIndex int) bool {
	return heap.GetRightChildIndex(parentIndex) < len(heap.heapContainer)
}

func (heap *Heap) LeftChild(parentIndex int) interface{} {
	if !heap.HasLeftChild(parentIndex) {
		panic("out of range")
	}

	return heap.heapContainer[heap.GetLeftChildIndex(parentIndex)]
}

func (heap *Heap) RightChild(parentIndex int) interface{} {
	if !heap.HasRightChild(parentIndex) {
		panic("out of range")
	}

	return heap.heapContainer[heap.GetRightChildIndex(parentIndex)]
}

func (heap *Heap) Parent(childIndex int) interface{} {
	if !heap.HasParent(childIndex) {
		panic("out of range")
	}

	return heap.heapContainer[heap.GetParentIndex(childIndex)]
}

func (heap *Heap) Swap(firstIndex, secondIndex int) {
	heap.heapContainer[firstIndex], heap.heapContainer[secondIndex] =
		heap.heapContainer[secondIndex], heap.heapContainer[firstIndex]
}

func (heap *Heap) Peek() interface{} {
	if len(heap.heapContainer) == 0 {
		return nil
	}
	return heap.heapContainer[0]
}

func (heap *Heap) Poll() interface{} {
	item := heap.heapContainer[0]
	heap.heapContainer[0] = heap.heapContainer[len(heap.heapContainer)-1]
	heap.heapContainer = heap.heapContainer[:len(heap.heapContainer)-1]

	heap.heapDown(0)
	return item
}

func (heap *Heap) Add(item interface{}) {
	heap.heapContainer = append(heap.heapContainer, item)
	heap.heapUp(len(heap.heapContainer) - 1)
}

func (heap *Heap) Len() int {
	return len(heap.heapContainer)
}

func (heap *Heap) Remove(index int) {
	if heap.Len() == 0 {
		return
	}
	if len(heap.heapContainer) == index+1 {
		heap.heapContainer = heap.heapContainer[:len(heap.heapContainer)-1]
		return
	}
	deletedItem := heap.heapContainer[index]
	heap.heapContainer[index] = heap.heapContainer[len(heap.heapContainer)-1]
	heap.heapContainer = heap.heapContainer[:len(heap.heapContainer)-1]

	if heap.less(deletedItem, heap.heapContainer[index]) {
		heap.heapDown(index)
	} else {
		heap.heapUp(index)
	}
}

func (heap *Heap) heapUp(index int) {
	for heap.HasParent(index) && heap.less(heap.heapContainer[index], heap.Parent(index)) {
		heap.Swap(index, heap.GetParentIndex(index))
		index = heap.GetParentIndex(index)
	}
}

func (heap *Heap) heapDown(index int) {
	for heap.HasLeftChild(index) {
		nextIndex := 0
		if heap.HasRightChild(index) && !heap.less(heap.LeftChild(index), heap.RightChild(index)) {
			nextIndex = heap.GetRightChildIndex(index)
		} else {
			nextIndex = heap.GetLeftChildIndex(index)
		}

		if !heap.less(heap.heapContainer[index], heap.heapContainer[nextIndex]) {
			break
		}

		heap.Swap(index, nextIndex)
		index = nextIndex
	}
}

func (heap *Heap) String() string {
	return fmt.Sprint(heap.heapContainer)
}
