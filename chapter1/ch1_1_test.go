package chapter1

import "testing"

type isUniqueTest struct {
	seq      string
	expected bool
}

type isUniqueTests []isUniqueTest

var isUniqueTestTable isUniqueTests = isUniqueTests{
	{
		seq:      "unique",
		expected: false,
	},
	{
		seq:      "umbrella",
		expected: false,
	},
	{
		seq:      "alex",
		expected: true,
	},
}

func TestIsUniqueV1(t *testing.T) {
	for _, test := range isUniqueTestTable {
		t.Run(test.seq, func(t *testing.T) {
			result := isUniqueV1(test.seq)

			if result != test.expected {
				t.Fatalf("isUniqueV1(\"%s\") expected %t but got %t", test.seq, test.expected, result)
			}
		})
	}
}

func TestIsUniqueV2(t *testing.T) {
	for _, test := range isUniqueTestTable {
		t.Run(test.seq, func(t *testing.T) {
			result := isUniqueV2(test.seq)

			if result != test.expected {
				t.Fatalf("isUniqueV1(\"%s\") expected %t but got %t", test.seq, test.expected, result)
			}
		})
	}
}

func TestIsUniqueV3(t *testing.T) {
	for _, test := range isUniqueTestTable {
		t.Run(test.seq, func(t *testing.T) {
			result := isUniqueV3(test.seq)

			if result != test.expected {
				t.Fatalf("isUniqueV1(\"%s\") expected %t but got %t", test.seq, test.expected, result)
			}
		})
	}
}
