package chapter1

import (
	"fmt"
	"testing"
)

type args struct {
	seq  string
	seq2 string
}

var tests = []struct {
	args args
	want bool
}{
	{
		args: args{
			seq:  "amor",
			seq2: "roma",
		},
		want: true,
	},
	{
		args: args{
			seq:  "a",
			seq2: "a",
		},
		want: true,
	},
	{
		args: args{
			seq:  "a",
			seq2: "b",
		},
		want: false,
	},
	{
		args: args{
			seq:  "amor",
			seq2: "marro",
		},
		want: false,
	},
}

func Test_isPermutationV1(t *testing.T) {
	for _, tt := range tests {
		name := fmt.Sprintf("%s - %s", tt.args.seq, tt.args.seq2)
		t.Run(name, func(t *testing.T) {
			if got := isPermutationV1(tt.args.seq, tt.args.seq2); got != tt.want {
				t.Errorf("isPermutationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPermutationV2(t *testing.T) {
	for _, tt := range tests {
		name := fmt.Sprintf("%s - %s", tt.args.seq, tt.args.seq2)
		t.Run(name, func(t *testing.T) {
			if got := isPermutationV2(tt.args.seq, tt.args.seq2); got != tt.want {
				t.Errorf("isPermutationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
