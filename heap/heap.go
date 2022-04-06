package heap

type MinHeap struct {
	Heap []int
}

func NewMinHeap(heap []int) *MinHeap {
	m := &MinHeap{heap}
	m.build()
	return m
}

func (mh *MinHeap) Push(value int) {
	mh.Heap = append(mh.Heap, value)
	mh.siftUp(len(mh.Heap) - 1)
}

func (mh *MinHeap) Pop() int {
	if len(mh.Heap) == 0 {
		return 0
	}
	res := mh.Heap[0]
	mh.Heap[0] = mh.Heap[len(mh.Heap)-1]
	mh.Heap = mh.Heap[:len(mh.Heap)-1]
	mh.siftDown(0, len(mh.Heap)>>1-1)
	return res
}

func (mh *MinHeap) build() {
	last := len(mh.Heap)>>1 - 1
	for i := last; i >= 0; i-- {
		left, right := 2*i+1, 2*i+2
		if mh.Heap[i] > mh.Heap[left] {
			swap(i, left, mh.Heap)
			if left <= last {
				mh.siftDown(left, last)
			}
		}
		// left不会越界,防止right越界
		if right < len(mh.Heap) && mh.Heap[i] > mh.Heap[right] {
			swap(i, right, mh.Heap)
			if right <= last {
				mh.siftDown(right, last)
			}
		}
	}
}

func (mh *MinHeap) siftUp(pos int) {
	for pos > 0 {
		parent := (pos - 1) / 2
		if mh.Heap[pos] > mh.Heap[parent] {
			break
		}
		swap(pos, parent, mh.Heap)
		pos = parent
	}
}

func (mh *MinHeap) siftDown(pos, end int) {
	l, r := 2*pos+1, 2*pos+2
	if mh.Heap[pos] > mh.Heap[l] {
		swap(pos, l, mh.Heap)
		if l <= end {
			mh.siftDown(l, end)
		}
	}
	// left不会越界,防止right越界
	if r < len(mh.Heap) && mh.Heap[pos] > mh.Heap[r] {
		swap(pos, r, mh.Heap)
		if r <= end {
			mh.siftDown(r, end)
		}
	}
}

func swap(x, y int, heap []int) {
	heap[x], heap[y] = heap[y], heap[x]
}
