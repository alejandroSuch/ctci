package chapter2

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := NewNode(3)

	if n == nil {
		t.Errorf("n cannot be nil")
	}

	if n.Next != nil {
		t.Errorf("n.Next should be nil")
	}

	if n.Data != 3 {
		t.Errorf("n.Data must be 3")
	}
}

func TestNode_Append(t *testing.T) {
	n := NewNode(3)
	n.Append(2)
	n.Append(1)

	if n.Next.Data != 2 {
		t.Errorf("n.Next.Data should be 2")
	}

	if n.Next.Next.Data != 1 {
		t.Errorf("n.Next.Next.Data should be 1")
	}

	if n.Next.Next.Next != nil {
		t.Errorf("n.Next.Next.Next should be 1")
	}
}

func TestNode_Delete(t *testing.T) {
	type testCase struct {
		name         string
		input        []int
		itemToDelete int
		expected     []int
	}

	tests := []testCase{
		{
			name:         "remove an unexisting item",
			input:        []int{1, 2, 3},
			itemToDelete: 8,
			expected:     []int{1, 2, 3},
		},
		{
			name:         "remove head",
			input:        []int{1, 2, 3},
			itemToDelete: 1,
			expected:     []int{2, 3},
		},
		{
			name:         "remove tail",
			input:        []int{1, 2, 3},
			itemToDelete: 3,
			expected:     []int{1, 2},
		},
		{
			name:         "remove middle",
			input:        []int{1, 2, 3},
			itemToDelete: 2,
			expected:     []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := newLinkedList(tt.input)
			want := newLinkedList(tt.expected)

			got := input.Delete(tt.itemToDelete)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Delete() = %v, want %v", got, want)
			}
		})
	}
}

func newLinkedList[T any](inputs []T) *Node[T] {
	result := NewNode(inputs[0])
	for _, it := range inputs[1:] {
		result.Append(it)
	}

	return result
}

func toList[T any](n *Node[T]) []T {
	var result []T

	for n != nil {
		result = append(result, n.Data)
		n = n.Next
	}

	return result
}
