package chapter2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "a", "b", "c", "a", "b"},
			expected: []string{"c", "a", "b"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test removeDuplicatesV1 - %d", i), func(t *testing.T) {
			input := newLinkedList(tt.input)
			expected := newLinkedList(tt.expected)
			output := removeDuplicatesV1(input)

			if !reflect.DeepEqual(output, expected) {
				t.Fatalf("removeDuplicatesV1 expected %v but got %v", output, expected)
			}
		})

		t.Run(fmt.Sprintf("test removeDuplicatesV2 - %d", i), func(t *testing.T) {
			input := newLinkedList(tt.input)
			expected := newLinkedList(tt.expected)
			output := removeDuplicatesV2(input)

			if !reflect.DeepEqual(output, expected) {
				t.Fatalf("removeDuplicatesV2 expected %v but got %v", output, expected)
			}
		})
	}
}

func Test_removeDuplicatesV3(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			input:    []string{"a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
		{
			input:    []string{"a", "a", "b", "c", "a", "b"},
			expected: []string{"a", "b", "c"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test removeDuplicatesV3 - %d", i), func(t *testing.T) {
			input := newLinkedList(tt.input)
			expected := newLinkedList(tt.expected)
			output := removeDuplicatesV3(input)

			if !reflect.DeepEqual(output, expected) {
				t.Fatalf("removeDuplicatesV3 expected %v but got %v", output, expected)
			}
		})
	}
}
