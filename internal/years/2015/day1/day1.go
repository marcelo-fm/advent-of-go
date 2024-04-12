package day1

import (
	"fmt"
	"os"
)

func partOne(input string) int {}

func partTwo(input string) int {}

func main() {
	filepath := "./input/input.txt"
	if len(os.Args) == 1 {
		switch os.Args[0] {
		case "1":
			fmt.Printf("Part One result: %d", partOne(filepath))
		case "2":
			fmt.Printf("Part Two Result: %d", partTwo(filepath))
		default:
			fmt.Printf("Part One result: %d", partOne(filepath))
			fmt.Printf("Part Two Result: %d", partTwo(filepath))
		}
	}
}
