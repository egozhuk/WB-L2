package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestShellCommands(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:  "pwd command",
			input: "pwd\n",
			expected: func() string {
				dir, _ := os.Getwd()
				return dir + "\n"
			}(),
		},
		{
			name:     "cd and pwd command",
			input:    "cd /tmp\npwd\n",
			expected: "/tmp\n",
		},
		{
			name:     "echo command",
			input:    "echo Hello, World!\n",
			expected: "Hello, World!\n",
		},
		{
			name:     "ps command",
			input:    "ps\n",
			expected: "", // Проверим только, что команда возвращает результат
		},
		{
			name:     "pipe command",
			input:    "echo Hello, World! | grep World\n",
			expected: "Hello, World!\n",
		},
		{
			name:     "quit command",
			input:    "\\quit\n",
			expected: "Exiting Simple Shell.\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "task.go")
			cmd.Stdin = strings.NewReader(test.input)

			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out

			err := cmd.Run()
			if err != nil {
				t.Fatalf("Command failed: %v\nOutput: %s", err, out.String())
			}

			if test.expected != "" {
				output := out.String()
				if !strings.Contains(output, test.expected) {
					t.Errorf("Expected output:\n%q\nGot output:\n%q", test.expected, output)
				}
			}
		})
	}
}
