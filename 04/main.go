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

	// scrathCardMap := make(map[int]int)

	// for i, _ := range strings.Split(string(file), "\n") {
	// 	scrathCardMap[i+1] = 1
	// }

	cards := strings.Split(string(file), "\n")

	sum := 0

	countsOfCards := make([]int, len(cards))
	for i := 0; i < len(cards); i++ {
		countsOfCards[i] = 1
	}

	for cardI, card := range strings.Split(string(file), "\n") {
		var goodNumbers []int
		var winningNumbers []int

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

		for io := 1; io <= len(winningNumbers); io++ {
			countsOfCards[cardI+io] += countsOfCards[cardI]

			fmt.Println(countsOfCards, countsOfCards[cardI+io], countsOfCards[cardI], cardI)
		}

		sum += countsOfCards[cardI]
	}

	fmt.Println("Sum part 1:", pointSum)

	fmt.Println("Sum part 2:", sum)

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
