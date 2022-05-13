package heap

type MinHeap struct {
	data []int
	size int
}

func (this *MinHeap) AddAll(arr []int) {
	if len(this.data) == 0 {
		if len(arr) <= this.size {
			this.data = append(this.data, arr...)
			arr = arr[0:0]
		} else {
			this.data = append(this.data, arr[:this.size]...)
			arr = arr[this.size:]
		}
	}
	this.adjustAll()
	for _, v := range arr {
		this.add(v)
	}
}

func (this *MinHeap) add(val int) {
	if this.isFull() {
		this.data[0] = val
		this.adjustDown(0)
	} else {
		this.data = append(this.data, val)
		this.adjustUp(this.len())
	}
}

func (this *MinHeap) pop() int {
	res := this.data[0]
	this.swap(0, this.len()-1)
	this.data = this.data[:this.len()-1]
	this.adjustDown(0)
	return res
}

func (this *MinHeap) adjustAll() {
	for i := this.len() - 1; i > 0; i-- {
		if this.data[i] < this.data[(i-1)/2] {
			this.swap(i, (i-1)/2)
			this.adjustDown(i)
		}
	}
}

func (this *MinHeap) adjustUp(index int) {
	if index < 0 {
		return
	}
	if this.data[index] < this.data[(index-1)/2] {
		this.swap(index, (index-1)/2)
		this.adjustUp((index - 1) / 2)
	}
}

func (this *MinHeap) adjustDown(index int) {
	if 2*index+1 >= this.len() {
		return
	}
	son := 2*index + 1
	if 2*index+2 < this.len() && this.data[2*index+2] < this.data[son] {
		son = 2*index + 2
	}
	if this.data[index] > this.data[son] {
		this.swap(index, son)
	}
	this.adjustDown(son)
}

func (this *MinHeap) len() int {
	return len(this.data)
}

func (this *MinHeap) isFull() bool {
	return len(this.data) == this.size
}

func (this *MinHeap) swap(s1, s2 int) {
	this.data[s1], this.data[s2] = this.data[s2], this.data[s1]
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{make([]int, 0, size), size}
}
