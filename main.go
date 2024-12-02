package main

import (
	"os"

	"github.com/hawkaii/advent_of_code_2024_go/day2"
)

func main() {
	input, err := os.ReadFile("../inputs/day2.txt")
	if err != nil {
		panic(err)
	}
	ans := day2.Part2(input)
	println(ans)
}
