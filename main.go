package main

import (
	"os"
	// "github.com/hawkaii/advent_of_code_2024_go/day9"
	"github.com/hawkaii/advent_of_code_2024_go/day11"
)

func main() {
	// input, err := os.ReadFile("../inputs/eg9.txt")
	input, err := os.ReadFile("../inputs/day11.txt")

	if err != nil {
		panic(err)
	}
	ans := day11.Part2(input)
	println(ans)

}
