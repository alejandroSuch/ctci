package chapter1

import "strconv"

func basicCompressV1(input string) string {
	if input == "" {
		return input
	}

	var current rune
	var occurrences uint64

	position := 0
	runes := []rune(input)
	inputLen := len(input)
	output := make([]rune, inputLen)

	var onChange = func() {
		output[position] = current
		position++
		if occurrences > 1 {
			formatted := strconv.FormatUint(occurrences, 10)
			for _, it := range formatted {
				output[position] = it
				position++
			}
		}
		occurrences = 0
	}

	for _, r := range runes {
		if r != current && occurrences != 0 {
			onChange()
		}
		current = r
		occurrences++
	}

	onChange()

	result := string(output[:position])
	if len(result) == inputLen {
		return input
	}

	return result
}
