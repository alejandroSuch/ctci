package chapter2

func sumListsV1(n1 *Node[int], n2 *Node[int]) *Node[int] {
	carry := 0
	result := NewNode(0)
	ptr := result

	for n1 != nil || n2 != nil {
		var s1 int
		var s2 int

		if n1 == nil {
			s1 = 0
		} else {
			s1 = n1.Data
			n1 = n1.Next
		}

		if n2 == nil {
			s2 = 0
		} else {
			s2 = n2.Data
			n2 = n2.Next
		}

		sum := s1 + s2 + carry

		if sum < 10 {
			ptr.Data = sum
			carry = 0
		} else {
			ptr.Data = sum % 10
			carry = 1
		}

		if n1 != nil || n2 != nil {
			ptr.Next = NewNode(0)
			ptr = ptr.Next
		} else if carry > 0 {
			ptr.Next = NewNode(carry)
			ptr = ptr.Next
		}
	}

	return result
}

func sumListsFollowUp(n1 *Node[int], n2 *Node[int]) *Node[int] {
	rn1 := reverse(n1)
	rn2 := reverse(n2)

	result := sumListsV1(rn1, rn2)

	return reverse(result)
}

func reverse(n *Node[int]) *Node[int] {
	if n == nil {
		return nil
	}

	var result *Node[int]

	for n != nil {
		aux := NewNode(n.Data)
		aux.Next = result
		result = aux
		n = n.Next
	}

	return result
}
