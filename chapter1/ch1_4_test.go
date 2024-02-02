package chapter1

import "testing"

func Test_isPalindromePermutationV1(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "tactcoa", want: true},
		{input: "tactoao", want: true},
		{input: "tactcoaoaoctcat", want: true},
		{input: "tacttoaob", want: false},
		{input: "a", want: true},
		{input: "ab", want: false},
		{input: "", want: true},
		{input: "amanaplanacanalpanama", want: true},
		{input: "amanalpancaaanplanama", want: true},
		{input: "amanaplanacanalpanamab", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isPalindromePermutationV1(tt.input); got != tt.want {
				t.Errorf("isPalindromePermutationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromePermutationV2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "tactcoa", want: true},
		{input: "tactoao", want: true},
		{input: "tactcoaoaoctcat", want: true},
		{input: "tacttoaob", want: false},
		{input: "a", want: true},
		{input: "ab", want: false},
		{input: "", want: true},
		{input: "amanaplanacanalpanama", want: true},
		{input: "amanalpancaaanplanama", want: true},
		{input: "amanaplanacanalpanamab", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := isPalindromePermutationV2(tt.input); got != tt.want {
				t.Errorf("isPalindromePermutationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
