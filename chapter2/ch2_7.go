package chapter2

func findIntersectionV1(n1 *Node[int], n2 *Node[int]) *Node[int] {
	t1, s1 := getTailAndSize(n1)
	t2, s2 := getTailAndSize(n2)

	h1 := n1
	h2 := n2

	if t1 != t2 {
		return nil
	}

	for {
		if s1 > s2 {
			h1 = h1.Next
			s1 = s1 - 1
		} else if s2 > s1 {
			h2 = h2.Next
			s2 = s2 - 1
		} else {
			break
		}
	}

	for h1 != nil {
		if h1 == h2 {
			return h1
		}

		h1 = h1.Next
		h2 = h2.Next
	}

	return nil
}

func getTailAndSize(n *Node[int]) (*Node[int], int) {
	if n == nil {
		return nil, 0
	}

	size := 1
	for n.Next != nil {
		n = n.Next
		size = size + 1
	}

	return n, size
}
