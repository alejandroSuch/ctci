package chapter2

import "reflect"

type LinkedList[T interface{}] struct {
	N      *Node[T]
	Length int
}

func (l *LinkedList[T]) Append(data T) {
	if l.N == nil {
		l.N = NewNode(data)
	} else {
		l.N.Append(data)
	}

	l.Length = l.Length + 1
}

func (l *LinkedList[T]) Reverse() LinkedList[T] {
	n := l.N
	if n == nil {
		return NewLinkedList[T](nil)
	}

	var result *Node[T]

	for n != nil {
		aux := NewNode(n.Data)
		aux.Next = result
		result = aux
		n = n.Next
	}

	return NewLinkedList(result)
}

func NewLinkedList[T interface{}](n *Node[T]) LinkedList[T] {
	length := 0
	aux := n
	if aux != nil {
		for aux != nil {
			aux = aux.Next
			length = length + 1
		}
	}
	return LinkedList[T]{
		N:      n,
		Length: length,
	}
}

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
