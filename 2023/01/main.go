package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var numericDictionary = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func convertToInt(char string) int {
	return numericDictionary[char]
}

func extractDigit(input string, regexp *regexp.Regexp) int {
	var matches []string

	for i := 0; i <= len(input); i++ {
		match := (regexp.FindAllString(input[i:], -1))
		matches = append(matches, match...)
	}

	firstDigit := convertToInt(matches[0])
	lastDigit := convertToInt(matches[len(matches)-1])

	return firstDigit*10 + lastDigit
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	sumPartOne := 0
	sumPartTwo := 0

	regexPartOne, _ := regexp.Compile("\\d")
	regexPartTwo, _ := regexp.Compile("(\\d|one|two|three|four|five|six|seven|eight|nine)")

	for scanner.Scan() {
		line := scanner.Text()

		sumPartOne += extractDigit(line, regexPartOne)
		sumPartTwo += extractDigit(line, regexPartTwo)
	}

	fmt.Println("Part 1: Sum", sumPartOne)
	fmt.Println("Part 2: Sum", sumPartTwo)
}
