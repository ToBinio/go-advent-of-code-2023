package io

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil
	}

	return lines
}
