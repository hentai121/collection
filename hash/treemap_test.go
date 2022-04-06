package hash

import (
	"fmt"
	"testing"
)

func TestTreeMap(t *testing.T) {
	treeMap := NewTreeMap()
	treeMap.Put(1, "1")
	treeMap.Put(2, 2)
	treeMap.Put(3, 3)
	treeMap.Put(4, 4)
	treeMap.Put(11, 11)
	treeMap.Put(-11, 121)
	fmt.Println(treeMap.Get(1))
	fmt.Println(treeMap.Get(2))
	fmt.Println(treeMap.Get(2) == 2)
	fmt.Println(treeMap.Get(1) == "1")
	for _, key := range treeMap.Keys() {
		fmt.Println(treeMap.Get(key))
	}

}
