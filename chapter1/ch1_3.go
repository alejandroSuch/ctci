package chapter1

func urlify(input string) string {
	space := rune(' ')
	spaceCount := 0

	for _, r := range input {
		if r == space {
			spaceCount++
		}
	}

	i := 0
	runes := make([]rune, len(input)+(2*spaceCount))
	for _, r := range input {
		if r == space {
			runes[i] = rune('%')
			runes[i+1] = rune('2')
			runes[i+2] = rune('0')
			i += 3
		} else {
			runes[i] = r
			i++
		}
	}

	return string(runes)
}
