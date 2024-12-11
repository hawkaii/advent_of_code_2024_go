package day11

import (
	"fmt"
	aoc "github.com/hawkaii/advent_of_code_2024_go/utils"
	"strconv"
	"sync"
)

type Queue[T any] struct {
	mu    sync.Mutex
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func stringSliceToQueue(s []string) *Queue[int] {
	queue := &Queue[int]{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		queue.Enqueue(i)
	}
	return queue
}

func (q *Queue[T]) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}

func (q *Queue[T]) Print() {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, v := range q.items {
		fmt.Println(v)
	}
}

func splitStringToInt(s string) (int, int) {
	l := len(s)
	mid := l / 2
	first := s[:mid]
	second := s[mid:]

	firstInt, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}
	secondInt, err := strconv.Atoi(second)
	if err != nil {
		panic(err)
	}

	return firstInt, secondInt
}

// blinkQueue processes the queue based on specific rules
func blinkQueue(queue *Queue[int], n int) *Queue[int] {
	for i := 0; i < n; i++ {
		l := queue.Size()
		for j := 0; j < l; j++ {
			a, ok := queue.Dequeue()
			if !ok {
				panic("queue underflow")
			}

			// Process the dequeued element based on conditions
			switch {
			case a == 0:
				queue.Enqueue(1)
			case len(strconv.Itoa(a))%2 == 0:
				first, second := splitStringToInt(strconv.Itoa(a))
				queue.Enqueue(first)
				queue.Enqueue(second)
			default:
				queue.Enqueue(a * 2024)
			}
		}
	}
	return queue
}

func blinkFrequencyMap(initial []int, n int) map[int]int {
	// Initialize frequency map
	freq := make(map[int]int)
	for _, v := range initial {
		freq[v]++
	}

	// Process for n iterations
	for i := 0; i < n; i++ {
		newFreq := make(map[int]int)
		for num, count := range freq {
			switch {
			case num == 0:
				newFreq[1] += count
			case len(strconv.Itoa(num))%2 == 0:
				first, second := splitStringToInt(strconv.Itoa(num))
				newFreq[first] += count
				newFreq[second] += count
			default:
				newFreq[num*2024] += count
			}
		}
		freq = newFreq
	}

	return freq
}

func Part1(input []byte) int {
	inputStr := aoc.ParseInput(input)
	fmt.Println(inputStr)
	strSlice := aoc.ParseString(inputStr[0])
	fmt.Println(strSlice)
	queue := stringSliceToQueue(strSlice)
	queue = blinkQueue(queue, 75)

	queue.Print()

	return queue.Size()
}

func Part2(input []byte) int {
	inputStr := aoc.ParseInput(input)
	strSlice := aoc.ParseString(inputStr[0])
	initial := make([]int, len(strSlice))
	for i, v := range strSlice {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		initial[i] = val
	}

	freq := blinkFrequencyMap(initial, 85)

	// Count total elements after 75 iterations
	totalCount := 0
	for _, count := range freq {
		totalCount += count
	}

	return totalCount
}

