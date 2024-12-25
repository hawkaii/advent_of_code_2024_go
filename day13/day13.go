package day13

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
)

func ParseValues(input string) map[string]int {

	pattern := `Button (?P<Button>[A-Z]): X\+(\d+), Y\+(?P<Y>\d+)|Prize: X=(?P<PX>\d+), Y=(?P<PY>\d+)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)
	if matches == nil {
		return nil
	}

	values := make(map[string]int)
	offset := 10000000000000
	for _, match := range matches {
		if match[1] != "" {
			button := match[1]
			x, _ := strconv.Atoi(match[2])
			y, _ := strconv.Atoi(match[3])
			values[button+"X"] = x
			values[button+"Y"] = y
		} else if match[4] != "" {
			x, _ := strconv.Atoi(match[4])
			y, _ := strconv.Atoi(match[5])
			values["PX"] = x + offset
			values["PY"] = y + offset
		}

	}
	return values

}

func calculateToken(input []string) []int {
	var tokens []int
	for _, line := range input {
		values := ParseValues(line)
		a1, b1, c1 := values["AX"], values["BX"], values["PX"]
		a2, b2, c2 := values["AY"], values["BY"], values["PY"]

		x, y, err := solveLinearSystem(a1, b1, c1, a2, b2, c2)

		if err != nil {
			fmt.Println(err)

		}

		token := x*3 + y
		tokens = append(tokens, token)

	}

	return tokens
}

func Part1(input []byte) int {
	lines := aoc.ParseInput2(input)

	tokens := calculateToken(lines)
	ans := 0
	for _, token := range tokens {
		ans += token
	}

	return ans

}

func solveLinearSystem(a1, b1, c1, a2, b2, c2 int) (int, int, error) {
	// Calculate the determinant
	det := a1*b2 - a2*b1
	if det == 0 {
		return 0, 0, errors.New("no unique solution: determinant is zero")
	}

	// Calculate potential solutions using Cramer's Rule
	xNumerator := c1*b2 - c2*b1
	yNumerator := a1*c2 - a2*c1

	// Ensure x and y are integers
	if xNumerator%det != 0 || yNumerator%det != 0 {
		return 0, 0, nil // Return (0, 0) if the solutions are not integers
	}

	x := xNumerator / det
	y := yNumerator / det

	fmt.Printf("x: %d, y: %d\n", x, y)
	return x, y, nil
}
