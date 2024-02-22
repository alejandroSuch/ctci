package chapter2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_deleteMiddleNodeV1(t *testing.T) {
	tests := []struct {
		list     []int
		index    int
		expected []int
	}{
		{
			index:    1,
			list:     []int{1, 2, 3},
			expected: []int{1, 3},
		},
		{
			list:     []int{1, 2, 3, 4},
			index:    1,
			expected: []int{1, 3, 4},
		},
		{
			list:     []int{1, 2, 3, 4},
			index:    2,
			expected: []int{1, 2, 4},
		},
		{
			list:     []int{1, 2},
			index:    0,
			expected: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v - %d - %v", tt.list, tt.index, tt.expected), func(t *testing.T) {
			list := newLinkedList(tt.list)
			expected := newLinkedList(tt.expected)
			n := list

			for i := 0; i < tt.index; i++ {
				n = n.Next
			}

			deleteMiddleNodeV1(n)

			if !reflect.DeepEqual(list, expected) {
				t.Fatalf("Expected %v but got %v", expected, list)
			}
		})
	}
}
