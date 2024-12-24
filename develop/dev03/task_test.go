package main

import (
	"os"
	"strings"
	"testing"
)

func createTempFile(content string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("", "sortutil_test")
	if err != nil {
		return nil, err
	}
	_, err = tmpFile.WriteString(content)
	if err != nil {
		return nil, err
	}
	return tmpFile, nil
}

func TestSortFile(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		config     SortConfig
		expected   string
		shouldFail bool
	}{
		{
			name:  "Lexicographic sort",
			input: "banana\napple\ncherry\n",
			config: SortConfig{
				Column: 1,
			},
			expected: "apple\nbanana\ncherry\n",
		},
		{
			name:  "Numeric sort",
			input: "10\n2\n30\n",
			config: SortConfig{
				Numeric: true,
			},
			expected: "2\n10\n30\n",
		},
		{
			name:  "Reverse order",
			input: "apple\nbanana\ncherry\n",
			config: SortConfig{
				Reverse: true,
			},
			expected: "cherry\nbanana\napple\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile, err := createTempFile(tt.input)
			if err != nil {
				t.Fatalf("Failed to create input file: %v", err)
			}
			defer os.Remove(inputFile.Name())

			tt.config.InputFile = inputFile.Name()

			lines, err := sortFile(tt.config)
			if err != nil {
				if !tt.shouldFail {
					t.Errorf("Unexpected error: %v", err)
				}
				return
			}

			output := strings.Join(lines, "\n") + "\n"
			if output != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, output)
			}
		})
	}
}
