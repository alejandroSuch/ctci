package chapter2

func isPalindromeV1(l LinkedList[rune]) bool {
	rl := l.Reverse()

	n1 := l.N
	n2 := rl.N
	for i := 0; i < l.Length; i++ {
		if n1.Data != n2.Data {
			return false
		}

		n1 = n1.Next
		n2 = n2.Next
	}

	return true
}

func isPalindromeV2(l LinkedList[rune]) bool {
	rl := l.Reverse()

	n1 := l.N
	n2 := rl.N
	for i := 0; i < l.Length/2; i++ {
		if n1.Data != n2.Data {
			return false
		}

		n1 = n1.Next
		n2 = n2.Next
	}

	return true
}

func isPalindromeV3(l LinkedList[rune]) bool {
	fast := l.N
	slow := l.N

	var stack *Node[rune]

	for fast != nil && fast.Next != nil {
		aux := NewNode(slow.Data)
		if stack == nil {
			stack = NewNode(aux.Data)
		} else {
			aux.Next = stack
			stack = aux
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Data != stack.Data {
			return false
		}

		slow = slow.Next
		stack = stack.Next
	}

	return true
}
