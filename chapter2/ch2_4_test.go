package chapter2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_partitionV1(t *testing.T) {
	tests := []struct {
		input     []int
		partition int
		expected  []int
	}{
		{
			input:     []int{1, 2, 5, 3, 8, 2, 6, 7},
			partition: 5,
			expected:  []int{1, 2, 3, 2, 8, 5, 6, 7},
		},
		{
			input:     []int{1, 2, 8, 3, 2, 5, 6, 1, 1},
			partition: 5,
			expected:  []int{1, 2, 3, 2, 1, 1, 6, 8, 5},
		},
		{
			input:     []int{1, 2, 8, 3, 2, 5, 6, 1, 1},
			partition: 2,
			expected:  []int{1, 1, 1, 3, 2, 5, 6, 2, 8},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v -> %v (partition %d)", tt.input, tt.expected, tt.partition), func(t *testing.T) {
			input := newLinkedList(tt.input)
			expected := newLinkedList(tt.expected)

			partitionV1(input, tt.partition)

			list := toList(input)
			list2 := toList(expected)

			if !reflect.DeepEqual(list, list2) {
				t.Fatalf("expected %v but got %v", expected, input)
			}
		})
	}
}
