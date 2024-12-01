package main

import (
	"advent2024/day01"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run main.go [dayXX | all]")
	}

	day := strings.ToLower(os.Args[1])

	switch day {
	case "day01":
		day01.Run()
	case "all":
		day01.Run()
		// Add additional days as they are implemented
		// day02.Run()
	default:
		log.Fatalf("Unknown day: %s", day)
	}
}
