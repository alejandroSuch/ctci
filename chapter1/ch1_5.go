package chapter1

import "unicode/utf8"

func oneAwayV1(input1 string, input2 string) bool {
	var w1, w2 int
	var r1, r2 rune

	changes := false
	input1Len := utf8.RuneCountInString(input1)
	input2Len := utf8.RuneCountInString(input2)

	if input1Len != input2Len && input2Len-1 != input1Len && input1Len-1 != input2Len {
		return false
	}

	for i, j := 0, 0; i < input1Len && j < input2Len; {
		r1, w1 = utf8.DecodeRuneInString(input1[i:])
		r2, w2 = utf8.DecodeRuneInString(input2[j:])

		if r1 != r2 {
			if changes {
				return false
			}

			changes = true

			if input1Len == input2Len {
				i, j = i+w1, j+w2
			} else if input1Len > input2Len {
				i = i + w1
			} else {
				j = j + w2
			}
		} else {
			i, j = i+w1, j+w2
		}
	}

	return true
}
