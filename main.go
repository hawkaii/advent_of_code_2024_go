package main

import (
	"os"
	// "github.com/hawkaii/advent_of_code_2024_go/day9"
	"github.com/hawkaii/advent_of_code_2024_go/day14"
)

func main() {
	// input, err := os.ReadFile("../inputs/eg14.txt")
	input, err := os.ReadFile("../inputs/day14.txt")

	if err != nil {
		panic(err)
	}
	ans := day14.Part2(input)
	println(ans)

}
