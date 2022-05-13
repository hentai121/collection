package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap(16)
	heap.AddAll([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 0, 22, 11, 44, 67, 33, 10})
	for heap.len() > 0 {
		t.Logf("heap.pop(): %v\n", heap.pop())
	}
}
