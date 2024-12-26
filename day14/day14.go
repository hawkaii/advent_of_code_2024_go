package day14

import (
	"fmt"
	"regexp"
	"strconv"

	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
)

type Point struct {
	x int
	y int
}

func Part1(input []byte) int {

	lines := aoc.ParseInput(input)

	// row := 11
	// column := 7

	row := 101
	column := 103

	mapValues := make(map[Point]int)

	for _, line := range lines {
		values := parseValues(line)
		fmt.Println("Parsed values:", values)

		x, y := posAfterTime(row, column, values["px"], values["py"], values["vx"], values["vy"], 100)
		fmt.Println("Position after time:", x, y)

		if mapValues[Point{x, y}] >= 1 {
			mapValues[Point{x, y}] += 1
		} else {
			mapValues[Point{x, y}] = 1
		}

	}

	fmt.Println("Map values:", mapValues)
	quadrants := countInQuadrants(row, column, mapValues)
	fmt.Println("Quadrants:", quadrants)

	ans := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

	return ans

}

func mapValues(row, column int, lines []string, time int) map[Point]int {
	mapValues := make(map[Point]int)

	for _, line := range lines {
		values := parseValues(line)
		fmt.Println("Parsed values:", values)

		x, y := posAfterTime(row, column, values["px"], values["py"], values["vx"], values["vy"], time)
		fmt.Println("Position after time:", x, y)

		if mapValues[Point{x, y}] >= 1 {
			mapValues[Point{x, y}] += 1
		} else {
			mapValues[Point{x, y}] = 1
		}

	}
	return mapValues
}

func Part2(input []byte) int {

	lines := aoc.ParseInput(input)

	// row := 11
	// column := 7

	row := 101
	column := 103

	mapValues := mapValues(row, column, lines, 7492)

	printMap(row, column, mapValues)

	// for i := 0; i < 100000; i++ {
	// 	clearMap(row, column, mapValues)
	//
	// 	for _, line := range lines {
	// 		values := parseValues(line)
	// 		fmt.Println("Parsed values:", values)
	//
	// 		x, y := posAfterTime(row, column, values["px"], values["py"], values["vx"], values["vy"], i)
	// 		fmt.Println("Position after time:", x, y)
	//
	// 		if mapValues[Point{x, y}] >= 1 {
	// 			mapValues[Point{x, y}] += 1
	// 		} else {
	// 			mapValues[Point{x, y}] = 1
	// 		}
	//
	// 	}
	//
	// 	if isChrismasTree(row, column, mapValues) {
	// 		fmt.Println("Found a chrismas tree at time:", i)
	// 		break
	// 	}
	// }

	// fmt.Println("Map values:", mapValues)

	return 0
}
func clearMap(row int, column int, mapValues map[Point]int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			mapValues[Point{i, j}] = 0
		}
	}
}

func isChrismasTree(row int, column int, mapValues map[Point]int) bool {
	lineLength := 10 // Define the length of the straight line to check for

	// Check for vertical lines
	for x := 0; x < row; x++ {
		count := 0
		for y := 0; y < column; y++ {
			if mapValues[Point{x, y}] >= 1 {
				count++
				if count >= lineLength {
					fmt.Println("Found a vertical line at x:", x)
					return true
				}
			} else {
				count = 0
			}
		}
	}

	return false
}

func printMap(row int, column int, mapValues map[Point]int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if mapValues[Point{i, j}] >= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countInQuadrants(row int, column int, mapValues map[Point]int) []int {
	quadrants := make([]int, 4)
	for point, value := range mapValues {
		if point.x == row/2 || point.y == column/2 {
			// Skip robots exactly in the middle
			fmt.Println("Skipping robot in the middle, point:", point)
			continue
		}
		if point.x < row/2 && point.y < column/2 {
			quadrants[0] += value
		} else if point.x > row/2 && point.y < column/2 {
			quadrants[1] += value
		} else if point.x < row/2 && point.y > column/2 {
			quadrants[2] += value
		} else if point.x > row/2 && point.y > column/2 {
			quadrants[3] += value
		}

	}
	return quadrants

}

func posAfterTime(row, column, px, py, vx, vy, time int) (int, int) {

	xcor := (px + vx*time) % row
	if xcor < 0 {
		xcor += row
	}

	ycor := (py + vy*time) % column
	if ycor < 0 {
		ycor += column
	}
	fmt.Println(xcor, ycor)
	return xcor, ycor

}
func parseValues(input string) map[string]int {
	pattern := `p=(?P<px>-?\d+),(?P<py>-?\d+) v=(?P<vx>-?\d+),(?P<vy>-?\d+)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	if matches == nil {
		return nil
	}

	values := make(map[string]int)

	for _, match := range matches {
		if match[1] != "" {
			px, _ := strconv.Atoi(match[1])
			py, _ := strconv.Atoi(match[2])
			vx, _ := strconv.Atoi(match[3])
			vy, _ := strconv.Atoi(match[4])

			values["px"] = px
			values["py"] = py
			values["vx"] = vx
			values["vy"] = vy
		}

	}
	return values

}
