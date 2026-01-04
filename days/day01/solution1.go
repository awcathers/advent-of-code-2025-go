package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(Solve("input1"))
}

func Solve(input_file string) int {
	data, err := os.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	dial_position := 50
	text := string(data)
	lines := strings.Split(text, "\n")
	count := 0
	for _, line := range lines {
		d := line[0:1]
		a, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Input did not contain a valid number", err)
		}
		if d == "L" {
			dial_position -= a
		} else if d == "R" {
			dial_position += a
		} else {
			log.Fatal("Invalid direction in input file")
		}

		if dial_position < 0 {
			dial_position += 100
		} else {
			dial_position = dial_position % 99
		}

		if dial_position == 0 {
			count = count + 1
		}

	}
	return count
}
