package chapter1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_zeroMatrix(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
				{10, 11, 0},
			},
			expected: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{7, 0, 0},
				{0, 0, 0},
			},
		},

		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("zeroMatrix %d", i), func(t *testing.T) {

			zeroMatrix(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Fatalf("Expected: %v, actual: %v\n", tt.expected, tt.input)
			}
		})
	}
}

func Test_zeroMatrixV2(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
				{10, 11, 0},
			},
			expected: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{7, 0, 0},
				{0, 0, 0},
			},
		},

		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("zeroMatrixV2 %d", i), func(t *testing.T) {
			zeroMatrixV2(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Fatalf("Expected: %v, actual: %v\n", tt.expected, tt.input)
			}
		})
	}
}

func Test_zeroMatrixV3(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
				{10, 11, 0},
			},
			expected: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{7, 0, 0},
				{0, 0, 0},
			},
		},

		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("zeroMatrixV3 %d", i), func(t *testing.T) {
			zeroMatrixV3(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Fatalf("Expected: %v, actual: %v\n", tt.expected, tt.input)
			}
		})

	}
}
