package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func makeIntSlice(input []byte) ([]int, []int) {

	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	pairs := strings.Split(lines[0], "\n")

	var list1, list2 []int
	for i, v := range pairs {

		if v == "" {
			continue
		}

		slice := strings.Split(strings.TrimSpace(v), " ")
		num1, err := strconv.Atoi(slice[0])
		if err != nil {
			fmt.Println(i)
			fmt.Println(err)
			return nil, nil
		}
		num2, err := strconv.Atoi(slice[3])
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)

	}

	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2

}

func Part1(input []byte) int {

	list1, list2 := makeIntSlice(input)

	differences := make([]int, len(list1))

	for i := range list1 {
		tmp := list2[i] - list1[i]
		if tmp < 0 {
			tmp = -tmp
		}
		differences[i] = tmp
	}

	sum := sum(differences)

	return sum

}

func sum(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}
	return sum
}

func Part2(input []byte) int {

	list1, list2 := makeIntSlice(input)

	count_list := make([]int, len(list1))

	counts := make(map[int]int)
	for _, v := range list2 {
		counts[v]++
	}
	for i, v := range list1 {
		count_list[i] = v * counts[v]
	}

	return sum(count_list)

}
