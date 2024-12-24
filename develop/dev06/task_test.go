package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		config cutConfig
		output []string
	}{
		{
			name:  "Basic usage",
			input: []string{"col1\tcol2\tcol3", "val1\tval2\tval3"},
			config: cutConfig{
				fields:    []int{1, 3},
				delimiter: "\t",
			},
			output: []string{"col1\tcol3", "val1\tval3"},
		},
		{
			name:  "Custom delimiter",
			input: []string{"a,b,c", "d,e,f"},
			config: cutConfig{
				fields:    []int{2},
				delimiter: ",",
			},
			output: []string{"b", "e"},
		},
		{
			name:  "Separated only",
			input: []string{"a\tb\tc", "d"},
			config: cutConfig{
				fields:    []int{2},
				delimiter: "\t",
				separated: true,
			},
			output: []string{"b"},
		},
		{
			name:  "Invalid field number",
			input: []string{"a\tb\tc"},
			config: cutConfig{
				fields:    []int{4},
				delimiter: "\t",
			},
			output: []string{""},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := cut(test.input, test.config)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected: %v, got: %v", test.output, result)
			}
		})
	}
}
