package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortConfig struct {
	InputFile      string
	OutputFile     string
	Column         int
	Numeric        bool
	Reverse        bool
	Unique         bool
	IgnoreTrailing bool
}

func sortFile(config SortConfig) ([]string, error) {
	lines, err := readLines(config.InputFile)
	if err != nil {
		return nil, err
	}

	if config.IgnoreTrailing {
		for i := range lines {
			lines[i] = strings.TrimRight(lines[i], " ")
		}
	}

	if config.Unique {
		uniqueMap := make(map[string]struct{})
		var uniqueLines []string
		for _, line := range lines {
			if _, exists := uniqueMap[line]; !exists {
				uniqueMap[line] = struct{}{}
				uniqueLines = append(uniqueLines, line)
			}
		}
		lines = uniqueLines
	}

	sort.SliceStable(lines, func(i, j int) bool {
		keyI := getColumn(lines[i], config.Column)
		keyJ := getColumn(lines[j], config.Column)

		if config.Numeric {
			valI, errI := strconv.ParseFloat(keyI, 64)
			valJ, errJ := strconv.ParseFloat(keyJ, 64)
			if errI == nil && errJ == nil {
				if config.Reverse {
					return valI > valJ
				}
				return valI < valJ
			}
		}

		if config.Reverse {
			return keyI > keyJ
		}
		return keyI < keyJ
	})

	return lines, nil
}

func getColumn(line string, column int) string {
	fields := strings.Fields(line)
	if column-1 >= 0 && column-1 < len(fields) {
		return fields[column-1]
	}
	return line
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func main() {
	config := SortConfig{}
	flag.StringVar(&config.InputFile, "i", "", "Input file path")
	flag.StringVar(&config.OutputFile, "o", "", "Output file path")
	flag.IntVar(&config.Column, "k", 1, "Column number for sorting (1-based)")
	flag.BoolVar(&config.Numeric, "n", false, "Sort by numeric value")
	flag.BoolVar(&config.Reverse, "r", false, "Sort in reverse order")
	flag.BoolVar(&config.Unique, "u", false, "Remove duplicate lines")
	flag.BoolVar(&config.IgnoreTrailing, "b", false, "Ignore trailing spaces")
	flag.Parse()

	if config.InputFile == "" {
		fmt.Fprintln(os.Stderr, "Input file is required")
		os.Exit(1)
	}

	lines, err := sortFile(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sorting file: %v\n", err)
		os.Exit(1)
	}

	if config.OutputFile != "" {
		err = writeLines(lines, config.OutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
			os.Exit(1)
		}
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
