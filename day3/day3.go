package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hawkaii/advent_of_code_2024_go/utils"
)

func Part1(input []byte) int {
	input_string := aoc.ParseInput(input)
	fmt.Println(input_string)

	/*
		mul(911,768)what(805,778)mul(690,737)from()who())select()<~mul(248,530)mul(638,821)mul(218,217)(^why();&mul(684,550)]
	*/
	//find all expression having the form `mul(1, 2)`
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(input_string[0], -1)

	sum := 0
	for _, match := range matches {
		int1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		int2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		product := int1 * int2
		sum += product

	}

	return sum
}

func findMulSum(s string) int {
	//find all expression having the form `mul(1, 2)`
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(s, -1)

	sum := 0
	for _, match := range matches {
		int1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		int2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		product := int1 * int2
		sum += product

	}

	return sum
}

func Part2(input []byte) int {

	input_string := aoc.ParseInput(input)

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	matches := re.FindAllString(input_string[0], -1)

	mulenabled := true
	sum := 0
	for _, match := range matches {
		if match == "do()" {
			mulenabled = true
		} else if match == "don't()" {
			mulenabled = false
		} else if mulenabled {
			var a, b int
			fmt.Sscanf(match, "mul(%d,%d)", &a, &b)
			sum += a * b
		}

	}

	return sum
}
