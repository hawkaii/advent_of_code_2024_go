package aoc

import "strings"

func ParseInput(input []byte) []string {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	return lines
}
