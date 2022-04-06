package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack([]int{1, 2, 3})
	stack.Push(1)
	fmt.Println(stack.Geek())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Geek())
}
