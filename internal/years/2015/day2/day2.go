package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func partOne(input []byte) (*int, error) {
	inputStr := string(input)
	dimArr := strings.Split(inputStr, "\n")
	totalArea := 0
	for _, dim := range dimArr {
		var l int
		var w int
		var h int
		_, err := fmt.Sscanf(strings.TrimSpace(dim), "%dx%dx%d", &l, &w, &h)
		if err != nil {
			return nil, err
		}
		d := make([]int, 3)
		d[0] = l
		d[1] = w
		d[2] = h
		for i := 0; i < len(d)-2; i++ {
			if d[i] > d[i+1] {
				d[i+1], d[i] = d[i], d[i+1]
			}
		}
		totalArea = totalArea + ((2 * l * w) + (2 * w * h) + (2 * h * l) + (d[0] * d[1]))
	}
	return &totalArea, nil
}

func main() {
	input, err := os.ReadFile("./input/input.txt")
	if err != nil {
		log.Fatalf("Error when reading input file in %s: %v", input, err)
	}
	result, err := partOne(input)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	fmt.Printf("partOne: %d sqrft", result)
}
