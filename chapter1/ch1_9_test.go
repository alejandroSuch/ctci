package chapter1

import (
	"fmt"
	"testing"
)

func Test_isStringRotation(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		want   bool
	}{
		{
			input1: "hello",
			input2: "helol",
			want:   false,
		},
		{
			input1: "hello world",
			input2: "orldhello w",
			want:   true,
		},
		{
			input1: "",
			input2: "",
			want:   true,
		},
		{
			input1: "a",
			input2: "a",
			want:   true,
		},
		{
			input1: "ab",
			input2: "ba",
			want:   true,
		},
		{
			input1: "ab",
			input2: "a",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("isStringRotationV1: %s - %s --> %t", tt.input1, tt.input2, tt.want), func(t *testing.T) {
			if got := isStringRotationV1(tt.input1, tt.input2); got != tt.want {
				t.Errorf("isStringRotationV1() = %v, want %v", got, tt.want)
			}
		})

		t.Run(fmt.Sprintf("isStringRotationV2: %s - %s --> %t", tt.input1, tt.input2, tt.want), func(t *testing.T) {
			if got := isStringRotationV2(tt.input1, tt.input2); got != tt.want {
				t.Errorf("isStringRotationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
