package chapter1

import (
	"fmt"
	"testing"
)

func Test_oneAwayV1(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		want   bool
	}{
		{
			input1: "hola",
			input2: "hola",
			want:   true,
		},
		{
			input1: "hola mundo",
			input2: "hola",
			want:   false,
		},
		{
			input1: "hola",
			input2: "hola mundo",
			want:   false,
		},
		{
			input1: "hola",
			input2: "bola",
			want:   true,
		},
		{
			input1: "hola",
			input2: "hole",
			want:   true,
		},
		{
			input1: "hola",
			input2: "halo",
			want:   false,
		},
		{
			input1: "hola",
			input2: "ola",
			want:   true,
		},
		{
			input1: "ola",
			input2: "hola",
			want:   true,
		},
		{
			input1: "ola",
			input2: "alo",
			want:   false,
		},
		{
			input1: "alo",
			input2: "ola",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("'%s' - '%s' ==> %t", tt.input1, tt.input2, tt.want), func(t *testing.T) {
			if got := oneAwayV1(tt.input1, tt.input2); got != tt.want {
				t.Errorf("oneAwayV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
