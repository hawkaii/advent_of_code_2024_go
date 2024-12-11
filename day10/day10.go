package day10

import (
	"fmt"
	"strconv"

	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
)

var positions = [][]int{}

func findZero(maze [][]string) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == "0" {
				positions = append(positions, []int{i, j})

			}
		}
	}

}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func checkBound(x, y int, maze [][]string) bool {
	return x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0])
}

func tillTrail(maze [][]string, num string, pos []int, visited [][]bool) int {
	count := 0
	for _, dir := range directions {
		x, y := pos[0]+dir[0], pos[1]+dir[1]
		if checkBound(x, y, maze) && !visited[x][y] {

			nextNum := maze[x][y]
			intNum, _ := strconv.Atoi(num)
			intNextNum, _ := strconv.Atoi(nextNum)

			fmt.Println(intNum, intNextNum)

			if intNum+1 == intNextNum {
				if intNextNum == 9 {
					fmt.Println("found 9")
					return 1
				}

				visited[x][y] = true
				count += tillTrail(maze, nextNum, []int{x, y}, visited)
				visited[x][y] = false
			}

		}
	}
	return count

}

func findTrails(maze [][]string) int {
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}
	count := 0
	for _, pos := range positions {
		x, y := pos[0], pos[1]
		visited[x][y] = true
		trails := tillTrail(maze, "0", []int{x, y}, visited)
		fmt.Println(trails)
		count += trails
	}
	return count

}
func Part1(input []byte) int {
	inputStr := aoc.ParseInput(input)
	fmt.Println(inputStr)

	maze := aoc.ParseMaze(inputStr)
	fmt.Println(maze)
	findZero(maze)
	ans := findTrails(maze)
	return ans
}
