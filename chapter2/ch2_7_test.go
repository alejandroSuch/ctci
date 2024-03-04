package chapter2

import (
	"testing"
)

func Test_findIntersectionV1(t *testing.T) {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n2.Next = n1
	n7 := NewNode(7)
	n7.Next = n2
	n6 := NewNode(7)
	n6.Next = n7
	n4 := NewNode(4)
	n4.Next = n6
	n9 := NewNode(9)
	n9.Next = n7
	n5 := NewNode(5)
	n5.Next = n9
	n8 := NewNode(1)
	n8.Next = n9
	n3 := NewNode(3)
	n3.Next = n8

	n10 := NewNode(10)
	n11 := NewNode(11)
	n11.Next = n10
	n12 := NewNode(12)
	n12.Next = n11
	n13 := NewNode(13)
	n13.Next = n12
	n14 := NewNode(14)
	n14.Next = n13
	n15 := NewNode(15)
	n15.Next = n14
	n16 := NewNode(16)
	n16.Next = n15
	n17 := NewNode(17)
	n17.Next = n16

	type args struct {
		n1 *Node[int]
		n2 *Node[int]
	}
	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		{
			name: "intersecting 2 lists with different size that intersect at n7",
			args: args{n1: n3, n2: n4},
			want: n7,
		},
		{
			name: "intersecting 2 lists with different size (changing order) that intersect at n7",
			args: args{n1: n4, n2: n3},
			want: n7,
		},
		{
			name: "intersecting 2 lists with the same size that intersect at n7",
			args: args{n1: n5, n2: n4},
			want: n7,
		},
		{
			name: "intersecting 2 lists with the same size that don't intersect",
			args: args{n1: n5, n2: n15},
			want: nil,
		},
		{
			name: "intersecting 2 lists with different size that don't intersect",
			args: args{n1: n5, n2: n17},
			want: nil,
		},
		{
			name: "intersecting 2 lists with different size (changing order) that don't intersect",
			args: args{n1: n17, n2: n5},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findIntersectionV1(tt.args.n1, tt.args.n2)
			if got != tt.want {
				t.Errorf("findIntersectionV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
