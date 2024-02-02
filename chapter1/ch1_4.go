package chapter1

func isPalindromePermutationV1(input string) bool {
	inputLen := 0
	letterCounter := make(map[rune]int)

	for _, it := range input {
		letterCounter[it]++
		inputLen++
	}

	var seenOdd = false
	isOddLen := inputLen%2 == 1
	for _, v := range letterCounter {
		if v%2 == 1 {
			if !isOddLen || seenOdd {
				return false
			} else {
				seenOdd = true
			}
		}
	}

	return true
}

func isPalindromePermutationV2(input string) bool {
	var bitVector uint64

	for _, r := range input {
		bitVector = bitVector ^ 1<<(r-'a') // XOR to revert a bit status in the bit vector
	}

	return bitVector == 0 || bitVector&(bitVector-1) == 0
}
