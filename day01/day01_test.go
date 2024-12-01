package day01

import (
	"advent2024/utils"
	"testing"
)

func loadInput(filename string) string {
	return utils.OpenFile(filename)
}

func TestPart1(t *testing.T) {
	input := loadInput("part1.txt")

	result := Part1(input)

	expected := 2066446

	if result != expected {
		t.Errorf("Part 1 result: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := loadInput("part1.txt")

	result := Part2(input)

	expected := 24931009

	if result != expected {
		t.Errorf("Part 2 result: expected %d, got %d", expected, result)
	}
}

func TestSample(t *testing.T) {
	input := loadInput("sample.txt")

	part1Result := Part1(input)
	part2Result := Part2(input)

	expectedPart1 := 11
	expectedPart2 := 31

	if part1Result != expectedPart1 {
		t.Errorf("Sample - Part 1 result: expected %d, got %d", expectedPart1, part1Result)
	}

	if part2Result != expectedPart2 {
		t.Errorf("Sample - Part 2 result: expected %d, got %d", expectedPart2, part2Result)
	}
}
