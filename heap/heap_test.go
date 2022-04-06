package heap

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	h := []int{110, 100, 90, 40, 80, 20, 60, 10, 30, 50}
	heap := NewMinHeap(h)
	fmt.Println(heap.Heap)
	heap.Push(15)
	fmt.Println(heap.Heap)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Heap)
}
