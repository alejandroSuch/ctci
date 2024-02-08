package chapter1

import (
	"strings"
	"unicode/utf8"
)

func isStringRotationV1(input1 string, input2 string) bool {
	_, pos := utf8.DecodeRuneInString(input2)

	if len(input1) <= 1 {
		return true
	}

	if len(input1) != len(input2) {
		return false
	}

	for i := pos; i < len(input2); {
		_, inc := utf8.DecodeRuneInString(input2[0:])

		sub1 := input2[0:i]
		sub2 := input2[i:]

		if strings.Contains(input1, sub1) && strings.Contains(input1, sub2) {
			return true
		}

		i = i + inc
	}

	return false
}

func isStringRotationV2(input1 string, input2 string) bool {
	result := strings.Contains(input2+input2, input1)
	return result
}
