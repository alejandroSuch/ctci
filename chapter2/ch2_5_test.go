package chapter2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sumListsV1(t *testing.T) {
	tests := []struct {
		input1   []int
		input2   []int
		expected []int
	}{
		{
			input1:   []int{7, 1, 6},
			input2:   []int{5, 9, 2},
			expected: []int{2, 1, 9},
		},
		{
			input1:   []int{0, 1, 6},
			input2:   []int{0, 9, 3},
			expected: []int{0, 0, 0, 1},
		},
		{
			input1:   []int{9, 9, 9},
			input2:   []int{9, 9, 9},
			expected: []int{8, 9, 9, 1},
		},
		{
			input1:   []int{9, 9},
			input2:   []int{9, 9, 9},
			expected: []int{8, 9, 0, 1},
		},
		{
			input1:   []int{9, 9, 3},
			input2:   []int{9, 9},
			expected: []int{8, 9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v + %v = %v", tt.input1, tt.input2, tt.expected), func(t *testing.T) {
			input1 := newLinkedList(tt.input1)
			input2 := newLinkedList(tt.input2)

			got := sumListsV1(input1, input2)
			gotList := toList(got)

			if !reflect.DeepEqual(gotList, tt.expected) {
				t.Errorf("sumLists() = %v, want %v", gotList, tt.expected)
			}
		})
	}
}

func Test_sumListsFollowUp(t *testing.T) {
	tests := []struct {
		input1   []int
		input2   []int
		expected []int
	}{
		{
			input1:   []int{6, 1, 7},
			input2:   []int{2, 9, 5},
			expected: []int{9, 1, 2},
		},
		{
			input1:   []int{6, 1, 0},
			input2:   []int{3, 9, 0},
			expected: []int{1, 0, 0, 0},
		},
		{
			input1:   []int{9, 9, 9},
			input2:   []int{9, 9, 9},
			expected: []int{1, 9, 9, 8},
		},
		{
			input1:   []int{9, 9},
			input2:   []int{9, 9, 9},
			expected: []int{1, 0, 9, 8},
		},
		{
			input1:   []int{3, 9, 9},
			input2:   []int{9, 9},
			expected: []int{4, 9, 8},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v + %v = %v", tt.input1, tt.input2, tt.expected), func(t *testing.T) {
			input1 := newLinkedList(tt.input1)
			input2 := newLinkedList(tt.input2)

			got := sumListsFollowUp(input1, input2)
			gotList := toList(got)

			if !reflect.DeepEqual(gotList, tt.expected) {
				t.Errorf("sumListsFollowUp() = %v, want %v", gotList, tt.expected)
			}
		})
	}
}
