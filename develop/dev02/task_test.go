package main

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
		{"a3\\", "", true}, // некорректная строка с незавершенной escape-последовательностью
	}

	for _, test := range tests {
		result, err := Unpack(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Input: %q, unexpected error state: %v", test.input, err)
		}
		if result != test.expected {
			t.Errorf("Input: %q, expected: %q, got: %q", test.input, test.expected, result)
		}
	}
}
