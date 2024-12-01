package utils

import (
	"log"
	"os"
)

func OpenFile(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("failed to read file %s: %v", filename, err)
	}

	return string(data)
}
