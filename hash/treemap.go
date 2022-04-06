package hash

import "sort"

type TreeMap struct {
	entity map[int]interface{}
	keys   []int
}

func NewTreeMap() *TreeMap {
	return &TreeMap{make(map[int]interface{}), make([]int, 0)}
}

func (t *TreeMap) Put(key int, val interface{}) {
	t.entity[key] = val
	t.keys = append(t.keys, key)
	sort.Ints(t.keys)
}

func (t TreeMap) Exist(key int) bool {
	if _, ok := t.entity[key]; ok {
		return true
	}
	return false
}

func (t TreeMap) Get(key int) interface{} {
	return t.entity[key]
}

func (t TreeMap) Keys() []int {
	return t.keys
}
