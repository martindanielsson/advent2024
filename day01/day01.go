package day01

import (
	"advent2024/utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	input := utils.OpenFile("day01/part1.txt")

	part1Result := Part1(input)
	fmt.Printf("Day 01, Part 1 Result: %d\n", part1Result)

	part2Result := Part2(input)
	fmt.Printf("Day 01, Part 2 Result: %d\n", part2Result)
}

func Part1(input string) int {
	left, right := ParseInput(input)
	result := 0
	sort.Ints(left)
	sort.Ints(right)

	for index, leftInt := range left {
		result += int(math.Abs(float64(right[index] - leftInt)))
	}

	return result
}

func Part2(input string) int {
	left, right := ParseInput(input)
	result := 0
	sort.Ints(left)
	sort.Ints(right)

	counts := CountAllOccurrences(right)

	for _, leftInt := range left {
		result += leftInt * counts[leftInt]
	}

	return result
}

func ParseInput(input string) ([]int, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var left []int
	var right []int

	for _, line := range lines {
		columns := strings.Fields(line)

		if len(columns) != 2 {
			continue
		}

		leftVal, errLeft := strconv.Atoi(columns[0])
		rightVal, errRight := strconv.Atoi(columns[1])

		if errLeft != nil && errRight != nil {
			log.Fatalf("failed to parse input left: %s right: %s", errLeft, errRight)
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	return left, right
}

func CountAllOccurrences(list []int) map[int]int {
	counts := make(map[int]int)

	for _, value := range list {
		counts[value]++
	}

	return counts
}
