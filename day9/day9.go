package day9

import (
	"fmt"
	"strconv"

	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
)

var blocks = map[int]int{}
var spaces = map[int]int{}

type Disk struct {
	start int
	digit int
	size  int
}

var disks = []Disk{}

func putDot(file []int, size int) []int {
	for i := 0; i < size; i++ {
		file = append(file, -1)
	}
	return file
}

func parseDisk(input string) []int {

	file := []int{}
	digit := 0
	for i, rune := range input {
		num, _ := strconv.Atoi(string(rune))

		if i%2 == 0 {
			blocks[digit] = num // for part 2
			disks = append(disks, Disk{start: i, digit: digit, size: num})
			for j := 0; j < num; j++ {
				file = append(file, digit)

			}
			digit += 1
		} else {
			spaces[i] = num // for part 2
			disks = append(disks, Disk{start: i, digit: -1, size: num})
			file = putDot(file, num)
		}

	}
	return file

}

func Part1(input []byte) int {
	inputStr := aoc.ParseInput(input)
	fmt.Println(inputStr)
	fmt.Println(parseDisk(inputStr[0]))

	fmt.Println(moveBlocks(parseDisk(inputStr[0])))
	movedFile := moveBlocks(parseDisk(inputStr[0]))

	ans := 0
	for i, block := range movedFile {
		ans += i * block
	}

	return ans

}

func moveBlocks(file []int) []int {
	l := len(file) - 1
	toCut := 0
	for i := 0; i < len(file); i++ {
		if file[i] != -1 {
			continue

		}
		for j := l; j > i; j-- {
			if file[j] != -1 {
				file[i] = file[j]
				file[j] = -1
				toCut = j
				break

			}
		}

	}

	fmt.Println(toCut)

	return file[:toCut]
}

func moveWholeBlocks() []Disk {
	newDisks := []Disk{}

	l := len(disks) - 1
	for i := 0; i < len(disks); i++ {
		if disks[i].digit != -1 {
			newDisks = append(newDisks, disks[i])
			continue
		}
		for j := l; j > i; j-- {
			if disks[j].digit != -1 {

			}
		}

	}

	return newDisks

}
func parseDisk2(input string) {
	digit := 0
	pos := 0
	for i, rune := range input {
		num, _ := strconv.Atoi(string(rune))

		if num == 0 {
			continue
		}

		if i%2 == 0 {
			blocks[digit] = num
			disks = append(disks, Disk{start: pos, digit: digit, size: num})
			pos += num
			digit += 1
		} else {

			spaces[digit] = num
			disks = append(disks, Disk{start: pos, digit: -1, size: num})
			pos += num
		}

	}
}

func Part2(input []byte) int {
	inputStr := aoc.ParseInput(input)
	fmt.Println(inputStr)

	parseDisk2(inputStr[0])

	fmt.Println(disks)

	return 0

}
