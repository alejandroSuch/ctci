package chapter2

func partitionV1(n *Node[int], partition int) {
	var pivot *Node[int] = nil

	for n != nil {
		if n.Data < partition {
			if pivot != nil {
				aux := pivot.Data
				pivot.Data = n.Data
				n.Data = aux
				pivot = pivot.Next
			}
		} else if pivot == nil {
			pivot = n
		}

		n = n.Next
	}
}
