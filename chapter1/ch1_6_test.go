package chapter1

import (
	"fmt"
	"testing"
)

func Test_basicCompressV1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "aaaa",
			want:  "a4",
		},
		{
			input: "abcd",
			want:  "abcd",
		},
		{
			input: "aaaabbcddd",
			want:  "a4b2cd3",
		},
		{
			input: "",
			want:  "",
		},
		{
			input: "a",
			want:  "a",
		},
		{
			input: " ",
			want:  " ",
		},
		{
			input: "  ",
			want:  "  ",
		},
		{
			input: "   ",
			want:  " 3",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s -> %s", tt.input, tt.want), func(t *testing.T) {
			if got := basicCompressV1(tt.input); got != tt.want {
				t.Errorf("basicCompressV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
