package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var cubeLimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func findMatches(pattern, s string) ([]string, [][]int, [][]string) {
	r := regexp.MustCompile(pattern)
	foundStrings := r.FindAllString(s, -1)
	stringIndices := r.FindAllStringIndex(s, -1)
	submatches := r.FindAllStringSubmatch(s, -1)
	return foundStrings, stringIndices, submatches
}

func processColorCubeCount(color, line string) (isValidGame bool, maxColorCount int) {
	isValidGame = true
	maxColorCount = -1

	_, _, colorSubmatches := findMatches("(\\d+) "+color, line)

	for _, submatch := range colorSubmatches {
		colorCount, _ := strconv.Atoi(submatch[1])

		if isValidGame && colorCount > cubeLimit[color] {
			isValidGame = false
		}

		if colorCount > maxColorCount {
			maxColorCount = colorCount
		}
	}

	return isValidGame, maxColorCount
}

func processGameLine(line string) (index int, isValidGame bool, power int) {
	_, _, gSubmatches := findMatches("Game (\\d+)", line)
	index, _ = strconv.Atoi(gSubmatches[0][1])

	validRedGame, redMax := processColorCubeCount("red", line)
	validGreenGame, greenMax := processColorCubeCount("green", line)
	validBlueGame, blueMax := processColorCubeCount("blue", line)

	isValidGame = validRedGame && validBlueGame && validGreenGame

	power = redMax * greenMax * blueMax

	return index, isValidGame, power
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	sumPartOne := 0
	sumPartTwo := 0

	for scanner.Scan() {
		line := scanner.Text()
		index, isValidGame, power := processGameLine(line)

		if isValidGame {
			sumPartOne += index
		}

		sumPartTwo += power
	}

	fmt.Println("Part 1: Sum", sumPartOne)
	fmt.Println("Part 2: Sum", sumPartTwo)
}
