package day2

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func makeSlicelist(input []byte) [][]int {

	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	pairs := strings.Split(lines[0], "\n")

	var slice_list [][]int

	for _, pair := range pairs {
		slice := strings.Split(pair, " ")
		var int_slice []int
		for _, s := range slice {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			int_slice = append(int_slice, i)
		}
		slice_list = append(slice_list, int_slice)
	}

	return slice_list

}
func isStrictlySorted(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	isAscending := true
	isDescending := true

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			isAscending = false
		}
		if nums[i] >= nums[i-1] {
			isDescending = false
		}
	}

	return isAscending || isDescending
}

func ifSafe(slice []int) bool {
	// Check if the slice is strictly sorted
	if !isStrictlySorted(slice) {
		return false
	}

	// Check if all adjacent differences are within the safe range
	for i := 0; i < len(slice)-1; i++ {
		if math.Abs(float64(slice[i]-slice[i+1])) > 3 {
			return false
		}
	}

	// If all conditions are met, the slice is safe
	return true
}

func Part1(input []byte) int {
	// Parse input

	slice_list := makeSlicelist(input)

	count_safe := 0

	for _, slice := range slice_list {

		if ifSafe(slice) {
			count_safe++
		}

	}
	fmt.Println(ifSafe(slice_list[0]))
	return count_safe
}

// removeIndex removes the element at the specified index from the slice
func removeIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		panic("index out of range")
	}
	return append(slice[:index], slice[index+1:]...)
}

// toleratingIfSafe attempts to remove one element at a time to find a "safe" slice
func toleratingIfSafe(slice []int) bool {
	if ifSafe(slice) {
		return true
	}

	for i := 0; i < len(slice); i++ {
		tmp := slices.Clone(slice)
		modifiedSlice := removeIndex(tmp, i)

		fmt.Println("Modified Slice:", modifiedSlice)

		if ifSafe(modifiedSlice) {
			return true
		}
	}

	// If no safe slice is found, return false
	return false
}

func Part2(input []byte) int {

	slice_list := makeSlicelist(input)

	count_safe := 0

	for _, slice := range slice_list {

		if toleratingIfSafe(slice) {
			count_safe++
		}

	}
	fmt.Println(slice_list[0])
	fmt.Println(toleratingIfSafe(slice_list[0]))
	return count_safe

}
