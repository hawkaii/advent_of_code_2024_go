package aoc

import "strings"

func ParseInput(input []byte) []string {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	return lines
}

func ParseString(input string) []string {
	s := strings.Split(input, " ")
	return s

}

func ParseMaze(lines []string) [][]string {

	maze := make([][]string, len(lines))
	for i, line := range lines {
		maze[i] = strings.Split(line, "")
	}

	return maze
}
