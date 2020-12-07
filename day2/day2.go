package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputPath = "./day2/raw-input.txt"

func getLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func processLine(line string) bool {
	// Algorithm
	// 1. Separate by spaces, get 3 words (min-max, letter, word)
	// 2. Separate min-max using -, convert them to Int
	// 3. Using index, get the letter letter[0]
	// 4. Using strings.Counts check the letter is within range

	input := strings.Split(line, " ")
	minMax := strings.Split(input[0], "-")
	letterColon := strings.Split(input[1], "")
	letter := letterColon[0]
	word := input[2]

	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	repetitions := strings.Count(word, letter)

	if repetitions >= min && repetitions <= max {
		return true
	}

	return false
}

// FindErrors reads line by line the input file and returns the incorrect lines
func FindErrors() {
	var numberOfErrors int
	lines, err := getLines(inputPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		hasError := processLine(line)

		if hasError {
			numberOfErrors++
		}
	}

	fmt.Printf("Number of Errors: %v\n", numberOfErrors)
}
