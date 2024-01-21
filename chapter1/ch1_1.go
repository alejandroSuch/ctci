package chapter1

import "sort"

func isUniqueV1(seq string) bool {
	occurrences := make(map[rune]int)

	for _, it := range seq {
		occurrences[it]++

		if occurrences[it] > 1 {
			return false
		}
	}

	return true
}

func isUniqueV2(seq string) bool {
	idx := 0
	occurrences := make([]rune, len(seq))

	for _, r1 := range seq {
		for _, r2 := range occurrences[0 : idx+1] {
			if r1 == r2 {
				return false
			}
		}

		occurrences[idx] = r1
		idx++
	}

	return true
}

func isUniqueV3(seq string) bool {
	runes := []rune(seq)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			return false
		}
	}

	return true
}
