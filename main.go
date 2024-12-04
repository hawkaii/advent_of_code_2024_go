package main

import (
	"os"

	"github.com/hawkaii/advent_of_code_2024_go/day3"
)

func main() {
	input, err := os.ReadFile("../inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	ans := day3.Part2(input)
	println(ans)
}
