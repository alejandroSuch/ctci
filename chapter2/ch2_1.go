package chapter2

func removeDuplicatesV1(n *Node[string]) *Node[string] {
	dupes := make(map[string]int)
	it := n

	for it != nil {
		dupes[it.Data] = dupes[it.Data] + 1
		it = it.Next
	}

	for k, v := range dupes {
		if v > 1 {
			for i := 1; i < v; i++ {
				n = n.Delete(k)
			}
		}
	}

	return n
}

func removeDuplicatesV2(n *Node[string]) *Node[string] {
	p1 := n

	for p1 != nil {
		p2 := p1.Next
		for p2 != nil {
			if p2.Data == p1.Data {
				n = n.Delete(p1.Data)
				break
			} else {
				p2 = p2.Next
			}
		}

		p1 = p1.Next

	}

	return n
}

func removeDuplicatesV3(n *Node[string]) *Node[string] {
	for current := n; current != nil; current = current.Next {
		for runner := current; runner.Next != nil; {
			if runner.Next.Data == current.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
	}

	for current := n; current != nil && current.Next != nil; {
		if current.Data == current.Next.Data {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return n
}
