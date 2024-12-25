package day12

import (
	"fmt"
	"sync"

	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
)

type Stack struct {
	items []int
	mux   sync.Mutex
}

func (s *Stack) Push(i int) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.items[len(s.items)-1]
}

// process each cell in the maze
func Part1(input []byte) int {
	lines := aoc.ParseInput(input)

	maze := aoc.ParseMaze(lines)
	fmt.Println(maze)

	result := calculateFenceCost(maze)

	// Display results
	for plantType, regions := range result {
		fmt.Printf("Plant Type %c:\n", plantType)
		for _, region := range regions {
			fmt.Printf("  Area: %d, Perimeter: %d\n", region.Area, region.Perimeter)
		}
	}

	return 0

}

type Region struct {
	Area      int
	Perimeter int
}

func calculateFenceCost(matrix [][]string) map[string][]Region {
	rows := len(matrix)
	cols := len(matrix[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	regions := make(map[string][]Region)

	// Flood-fill using a stack
	var dfs func(int, int, string) (int, int)
	dfs = func(x, y int, plantType string) (int, int) {
		stack := [][2]int{{x, y}}
		area := 0
		perimeter := 0

		for len(stack) > 0 {
			// Pop from the stack
			cell := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i, j := cell[0], cell[1]

			// Skip already visited cells
			if visited[i][j] {
				continue
			}

			visited[i][j] = true
			area++

			// Check all four directions
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols {
					if matrix[ni][nj] == plantType {
						if !visited[ni][nj] {
							stack = append(stack, [2]int{ni, nj})
						}
					} else {
						perimeter++
					}
				} else {
					perimeter++
				}
			}
		}

		return area, perimeter
	}

	// Process the matrix
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] {
				plantType := matrix[i][j]
				area, perimeter := dfs(i, j, plantType)
				regions[plantType] = append(regions[plantType], Region{Area: area, Perimeter: perimeter})
			}
		}
	}

	return regions
}
