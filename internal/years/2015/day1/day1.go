package main

import (
	"fmt"
	"log"
	"os"
)

func partOne(input string) int {
	inputBuffer, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error when reading input file in %s: %v", input, err)
	}
	floor := 0
	for _, char := range inputBuffer {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}

func partTwo(input string) int {
	inputBuffer, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error when reading input file in %s: %v", input, err)
	}
	floor := 0
	for i, char := range inputBuffer {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return 0
}

func main() {
	filepath := "./input/input.txt"
	if len(os.Args) == 1 {
		switch os.Args[0] {
		case "1":
			fmt.Printf("Part One result: %d", partOne(filepath))
		case "2":
			fmt.Printf("Part Two Result: %d", partTwo(filepath))
		}
	} else {
		fmt.Printf("Part One result: %d\n", partOne(filepath))
		fmt.Printf("Part Two Result: %d", partTwo(filepath))
	}
}
