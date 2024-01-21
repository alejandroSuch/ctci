package chapter1

import "sort"

func isPermutationV1(seq string, seq2 string) bool {
	if len(seq) != len(seq2) {
		return false
	}

	r1 := []rune(seq)
	r2 := []rune(seq2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	for i, _ := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}

	return true
}

func isPermutationV2(seq string, seq2 string) bool {
	if len(seq) != len(seq2) {
		return false
	}

	items := make(map[rune]int)

	for _, r := range seq {
		items[r]++
	}

	for _, r := range seq2 {
		items[r]--
	}

	for _, val := range items {
		if val != 0 {
			return false
		}
	}

	return true
}
