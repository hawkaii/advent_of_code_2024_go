package main

import (
	"os"

	"github.com/hawkaii/advent_of_code_2024_go/day1"
)

func main() {
	input, err := os.ReadFile("./inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	ans := day1.Part1(input)
	println(ans)
}
