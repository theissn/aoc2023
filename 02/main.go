package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input-1")

	if err != nil {
		panic(err)
	}

	puzzle1(string(data))
	puzzle2(string(data))
}

func puzzle1(data string) {
	fmt.Println("Puzzle 1")

	cubeBag := map[string]int{
		"green": 13,
		"red":   12,
		"blue":  14,
	}

	reNumbers := regexp.MustCompile("[0-9]+")
	reLetters := regexp.MustCompile("[a-z]+")

	var possibleGames []int

	for i, v := range strings.Split(data, "\n") {
		currentGame := i + 1
		input := strings.ReplaceAll(v, fmt.Sprintf("Game %d: ", currentGame), "")

		possible := true

		for _, v1 := range strings.Split(input, ";") {
			for _, v2 := range strings.Split(v1, ", ") {
				count, err := strconv.Atoi(strings.TrimSpace(reLetters.ReplaceAllString(v2, "")))

				if err != nil {
					panic(err)
				}

				colour := strings.TrimSpace(reNumbers.ReplaceAllString(v2, ""))

				if cubeBag[colour] < count {
					possible = false
				}
			}
		}

		if possible {
			possibleGames = append(possibleGames, currentGame)
		}
	}

	sum := 0

	for _, v := range possibleGames {
		sum += v
	}

	fmt.Println("Puzzle 1:", sum)
}

func puzzle2(data string) {
	fmt.Println("Puzzle 2")

	reNumbers := regexp.MustCompile("[0-9]+")
	reLetters := regexp.MustCompile("[a-z]+")

	var power []int

	for i, v := range strings.Split(data, "\n") {
		currentGame := i + 1

		input := strings.ReplaceAll(v, fmt.Sprintf("Game %d: ", currentGame), "")

		cubeCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, v1 := range strings.Split(input, ";") {
			for _, v2 := range strings.Split(v1, ", ") {
				count, err := strconv.Atoi(strings.TrimSpace(reLetters.ReplaceAllString(v2, "")))

				if err != nil {
					panic(err)
				}

				colour := strings.TrimSpace(reNumbers.ReplaceAllString(v2, ""))

				if cubeCount[colour] < count {
					cubeCount[colour] = count
				}
			}
		}

		power = append(power, (cubeCount["red"] * cubeCount["green"] * cubeCount["blue"]))
	}

	sum := 0

	for _, v := range power {
		sum += v
	}

	fmt.Println("Puzzle 2:", sum)
}
