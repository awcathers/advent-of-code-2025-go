package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func Solve(input_file string) int {
	data, err := os.ReadFile(input_file)
	if err != nil {
		log.Fatal(err)
	}

	dial_position := 50
	text := string(data)
	lines := strings.Split(text, "\r\n")
	count := 0
	for _, line := range lines {
		d := line[0:1]
		a, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("Input did not contain a valid number", err)
		}
		if d == "L" {
			dial_position -= a
		} else if line == "R" {
			dial_position += a
		}

		if dial_position == 0 {
			count = count + 1
		}

	}
	return count
}
