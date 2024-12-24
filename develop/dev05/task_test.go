package main

import (
	"reflect"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		config grepConfig
		output []string
	}{
		{
			name:  "Simple match",
			input: []string{"hello", "world", "hello world"},
			config: grepConfig{
				pattern: "hello",
			},
			output: []string{"hello", "hello world"},
		},
		{
			name:  "Match with ignore case",
			input: []string{"Hello", "world", "HELLO WORLD"},
			config: grepConfig{
				pattern:    "hello",
				ignoreCase: true,
			},
			output: []string{"Hello", "HELLO WORLD"},
		},
		{
			name:  "Invert match",
			input: []string{"hello", "world", "hello world"},
			config: grepConfig{
				pattern: "hello",
				invert:  true,
			},
			output: []string{"world"},
		},
		{
			name:  "Count matches",
			input: []string{"hello", "world", "hello world"},
			config: grepConfig{
				pattern: "hello",
				count:   true,
			},
			output: []string{"2"},
		},
		{
			name:  "Context lines",
			input: []string{"line1", "line2", "line3", "line4", "line5"},
			config: grepConfig{
				pattern: "line3",
				context: 1,
			},
			output: []string{"line2", "line3", "line4"},
		},
		{
			name:  "Line numbers",
			input: []string{"hello", "world", "hello world"},
			config: grepConfig{
				pattern: "hello",
				lineNum: true,
			},
			output: []string{"1:hello", "3:hello world"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := grep(test.input, test.config)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected: %v, got: %v", test.output, result)
			}
		})
	}
}
