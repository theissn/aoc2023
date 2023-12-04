package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input")
	pointSum := 0

	for index, card := range strings.Split(string(file), "\n") {
		var goodNumbers []int

		re := regexp.MustCompile(`Card\s+\d+:\s+`)

		cleanString := re.ReplaceAllString(card, "")

		numbers := strings.Split(cleanString, " | ")

		for _, v := range strings.Split(numbers[0], " ") {
			num, err := strconv.Atoi(v)

			if err != nil {
				continue
			}

			goodNumbers = append(goodNumbers, num)
		}

		var winningNumbers []int

		for _, v := range strings.Split(numbers[1], " ") {
			num, err := strconv.Atoi(v)

			if err != nil {
				continue
			}

			for _, v1 := range goodNumbers {
				if num == v1 {
					winningNumbers = append(winningNumbers, v1)
				}
			}
		}

		pointSum += getPoints(winningNumbers)
	}

	fmt.Println("Sum part 1:", pointSum)

}

func calcPoints(matches int) int {
	return int(math.Pow(2, float64(matches-1)))
}

func getPoints(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) <= 2 {
		return len(numbers)
	}

	return calcPoints(len(numbers))
}
