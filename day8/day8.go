package day8

import (
	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
)

// checkBoundaries ensures the coordinates are within the grid bounds
func checkBoundaries(maze [][]string, pos []int) bool {
	return pos[0] >= 0 && pos[0] < len(maze) && pos[1] >= 0 && pos[1] < len(maze[0])
}

// calclulateDistance returns the difference vector between two points
func calclulateDistance(a, b []int) []int {
	return []int{b[0] - a[0], b[1] - a[1]}
}

func putAntinodes(maze [][]string, antennas map[string][][]int) {
	for _, antennaGroup := range antennas {
		for i := 0; i < len(antennaGroup)-1; i++ {
			for j := i + 1; j < len(antennaGroup); j++ {
				// Calculate the direction vector (difference between antennas)
				diff := calclulateDistance(antennaGroup[i], antennaGroup[j])

				// Place antinodes upwards from the first antenna
				upNode := []int{antennaGroup[i][0] + diff[0], antennaGroup[i][1] + diff[1]}
				if !checkBoundaries(maze, upNode) {
					maze[antennaGroup[i][0]][antennaGroup[i][1]] = "#"
				}

				// Place antinodes downwards from the second antenna
				downNode := []int{antennaGroup[j][0] - diff[0], antennaGroup[j][1] - diff[1]}
				if checkBoundaries(maze, downNode) {
					maze[antennaGroup[j][0]][antennaGroup[j][1]] = "#"
				}
			}
		}
	}
}

// putEnhancedAntinodes places the antinodes on the grid based on antenna positions
func putEnhancedAntinodes(maze [][]string, antennas map[string][][]int) {
	for _, antennaGroup := range antennas {
		for i := 0; i < len(antennaGroup)-1; i++ {
			for j := i + 1; j < len(antennaGroup); j++ {
				// Calculate the direction vector (difference between antennas)
				diff := calclulateDistance(antennaGroup[i], antennaGroup[j])

				// Place antinodes upwards from the first antenna
				upNode := []int{antennaGroup[i][0] + diff[0], antennaGroup[i][1] + diff[1]}
				maze[antennaGroup[i][0]][antennaGroup[i][1]] = "#" // Mark the first antenna itself as an antinode
				for checkBoundaries(maze, upNode) {
					maze[upNode[0]][upNode[1]] = "#"
					upNode[0] += diff[0]
					upNode[1] += diff[1]
				}

				// Place antinodes downwards from the second antenna
				downNode := []int{antennaGroup[j][0] - diff[0], antennaGroup[j][1] - diff[1]}
				maze[antennaGroup[j][0]][antennaGroup[j][1]] = "#" // Mark the second antenna itself as an antinode
				for checkBoundaries(maze, downNode) {
					maze[downNode[0]][downNode[1]] = "#"
					downNode[0] -= diff[0]
					downNode[1] -= diff[1]
				}
			}
		}
	}
}

// countAntinodes counts the number of antinodes on the grid
func countAntinodes(maze [][]string) int {
	count := 0
	for _, row := range maze {
		for _, cell := range row {
			if cell == "#" {
				count++
			}
		}
	}
	return count
}

func Part1(input []byte) int {
	input_string := aoc.ParseInput(input)

	// Parse the input
	maze, antennas := aoc.ParseMaze(input_string)

	// Place the antinodes on the grid
	putAntinodes(maze, antennas)

	// Count the number of antinodes
	return countAntinodes(maze)
}
