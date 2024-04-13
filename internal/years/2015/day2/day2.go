package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne(input []byte) int {
	inputStr := string(input)
	dimArr := strings.Split(inputStr, "\n")
	totalArea := 0
	for _, dim := range dimArr {
		var lStr string
		var wStr string
		var hStr string
		i, err := fmt.Sscanf(dim, "%dx%dx%d", lStr, wStr, hStr)
		if err != nil || i != 3 {
			log.Println(err)
			return -1
		}
		l, err := strconv.Atoi(lStr)
		w, err := strconv.Atoi(wStr)
		h, err := strconv.Atoi(hStr)
		d := make([]int, 3)
		d[0] = l
		d[1] = w
		d[2] = h
		for i := 0; i < len(d)-2; i++ {
			if d[i] > d[i+1] {
				d[i+1], d[i] = d[i], d[i+1]
			}
		}
		if err != nil {
			log.Println(err)
			return -1
		}
		totalArea = totalArea + ((2 * l * w) + (2 * w * h) + (2 * h * l) + (d[0] * d[1]))
	}
	return totalArea
}

func partTwo(input []byte) int { return 0 }

func main() {
	input, err := os.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error when reading input file in %s: %v", input, err)
	}
	if len(os.Args) == 1 {
		switch os.Args[0] {
		case "1":
			fmt.Printf("partOne: %d sqrft", partOne(input))
		case "2":
			fmt.Printf("partTwo: %d sqrft", partTwo(input))
		}
	}
}
