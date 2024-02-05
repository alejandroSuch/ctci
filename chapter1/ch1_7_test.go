package chapter1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
			},
			[][]int{
				[]int{13, 9, 5, 1},
				[]int{14, 10, 6, 2},
				[]int{15, 11, 7, 3},
				[]int{16, 12, 8, 4},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			rotateMatrix(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Fatalf("Expected: %v, actual: %v\n", tt.expected, tt.input)
			}
		})
	}
}
