package chapter2

import "reflect"

type Node[T interface{}] struct {
	Next *Node[T]
	Data T
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{
		Data: val,
	}
}

func (n *Node[T]) Append(data T) {
	cur := n

	for ; cur.Next != nil; cur = cur.Next {
	}

	cur.Next = NewNode[T](data)
}

func (n *Node[T]) Delete(data T) *Node[T] {
	if reflect.DeepEqual(n.Data, data) {
		return n.Next
	}

	for cur := n; cur.Next != nil; cur = cur.Next {
		if cur.Next != nil && reflect.DeepEqual(cur.Next.Data, data) {
			cur.Next = cur.Next.Next
			break
		}

	}
	return n
}
