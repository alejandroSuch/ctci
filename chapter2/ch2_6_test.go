package chapter2

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "arra",
			want:  true,
		},
		{
			input: "arera",
			want:  true,
		},
		{
			input: "arrea",
			want:  false,
		},
		{
			input: "amor a roma",
			want:  true,
		},
		{
			input: "dabalearrozalazorraellabad",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("isPalindromeV1(%s) => %t", tt.input, tt.want), func(t *testing.T) {
			var l LinkedList[rune]
			for _, r := range tt.input {
				l.Append(r)
			}

			if got := isPalindromeV1(l); got != tt.want {
				t.Errorf("isPalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("isPalindromeV2(%s) => %t", tt.input, tt.want), func(t *testing.T) {
			var l LinkedList[rune]
			for _, r := range tt.input {
				l.Append(r)
			}

			if got := isPalindromeV2(l); got != tt.want {
				t.Errorf("isPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("isPalindromeV3(%s) => %t", tt.input, tt.want), func(t *testing.T) {
			var l LinkedList[rune]
			for _, r := range tt.input {
				l.Append(r)
			}

			if got := isPalindromeV3(l); got != tt.want {
				t.Errorf("isPalindromeV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
