package chapter1

import "testing"

func Test_urlify(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello world",
			want:  "hello%20world",
		},
		{
			input: "helloworld",
			want:  "helloworld",
		},
		{
			input: "helloworld ",
			want:  "helloworld%20",
		},
		{
			input: " helloworld ",
			want:  "%20helloworld%20",
		},
		{
			input: " helloworld",
			want:  "%20helloworld",
		},
		{
			input: " hello   world ",
			want:  "%20hello%20%20%20world%20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := urlify(tt.input); got != tt.want {
				t.Errorf("urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}
