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

func TestStack_Geek(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stack
		want    int
		wantErr bool
	}{
		{"test1", &Stack{[]int{1, 2, 3}}, 1, false},
		{"test1", &Stack{[]int{3, 4, 5}}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Geek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Geek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.Geek() = %v, want %v", got, tt.want)
			}
		})
	}
}
