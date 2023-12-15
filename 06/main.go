package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sample, _ := os.ReadFile("sample")
	input, _ := os.ReadFile("input")

	puzzle(sample)
	puzzle(input)

	println()

	puzzle2(sample)
	puzzle2(input)
}

func puzzle(input []byte) {
	data := strings.Split(string(input), "\n")

	time := strings.Replace(data[0], "Time: ", "", -1)
	dist := strings.Replace(data[1], "Distance: ", "", -1)

	numbersRe := regexp.MustCompile(`\d+`)

	times := numbersRe.FindAllString(time, -1)
	dists := numbersRe.FindAllString(dist, -1)

	fmt.Println(calculatePossibilities(times, dists))
}

func puzzle2(input []byte) {
	data := strings.Split(string(input), "\n")

	time := strings.Replace(data[0], "Time: ", "", -1)
	dist := strings.Replace(data[1], "Distance: ", "", -1)

	numbersRe := regexp.MustCompile(`\d+`)

	times := numbersRe.FindAllString(time, -1)
	dists := numbersRe.FindAllString(dist, -1)

	var singleTime []string
	singleTime = append(singleTime, strings.Join(times, ""))

	var singleDist []string
	singleDist = append(singleDist, strings.Join(dists, ""))

	fmt.Println(calculatePossibilities(singleTime, singleDist))
}

func calculatePossibilities(times, dists []string) int {
	possibilities := 0

	for i, t := range times {
		poss := 0

		tme, _ := strconv.Atoi(t)
		dst, _ := strconv.Atoi(dists[i])

		for ii := 1; ii < tme; ii++ {
			distanceMoved := ii * (tme - ii)

			if distanceMoved > dst {
				poss++
			}

		}

		if possibilities == 0 {
			possibilities = poss
			continue
		}

		possibilities *= poss

	}

	return possibilities
}
