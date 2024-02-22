package chapter2

func deleteMiddleNodeV1(n *Node[int]) {
	if n.Next.Next == nil {
		n.Data = n.Next.Data
		n.Next = nil
	} else {
		n.Data = n.Next.Data
		deleteMiddleNodeV1(n.Next)
	}
}
