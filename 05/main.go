package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 5")
	println()

	fmt.Println("Sample")
	sample, _ := os.ReadFile("sample")
	puzzle(string(sample))

	println()

	fmt.Println("Input")
	input, _ := os.ReadFile("input")
	puzzle(string(input))
}

func puzzle(input string) {
	lines := strings.Split(input, "\n\n")
	seedString := strings.Replace(lines[0], "seeds: ", "", 1)

	var seeds []int

	for _, v := range strings.Split(seedString, " ") {
		num, _ := strconv.Atoi(v)
		seeds = append(seeds, num)
	}

	partOneValue := seedMath(lines, seeds)
	fmt.Println("Part 1", partOneValue)

	// partTwoValue := SeedMathPart2(lines, seeds)
	partTwoValue := "fuck if i know, i don't own enough ram to do that"
	fmt.Println("Part 2", partTwoValue)
}

func seedMath(lines []string, seeds []int) int {
	val := 0

	for _, v := range lines[1:] {
		section := strings.Split(v, "\n")
		var newSeeds []int

		for _, seed := range seeds {
			found := false

			for _, v := range section[1:] {
				numbers := strings.Split(v, " ")
				destNum, _ := strconv.Atoi(numbers[0])
				sourceNum, _ := strconv.Atoi(numbers[1])
				countNum, _ := strconv.Atoi(numbers[2])

				if seed >= sourceNum && seed <= sourceNum+countNum && !found {
					newSeeds = append(newSeeds, destNum+(seed-sourceNum))

					found = true
				}
			}

			if !found {
				newSeeds = append(newSeeds, seed)
			}
		}

		seeds = newSeeds
		seedsForSorting := slices.Clone(seeds)
		sort.Ints(seedsForSorting)
		val = seedsForSorting[0]
	}

	return val
}

type rage struct {
	x, y int
}

func SeedMathPart2(lines []string, seeds []int) int {
	val := 0

	var ranges []rage

	for i, v := range seeds {
		if (i % 2) != 0 {
			continue
		}

		start := v
		end := v + seeds[i+1]
		ranges = append(ranges, rage{start, end})
	}

	for _, v := range ranges {
		for _, v1 := range lines[1:] {
			section := strings.Split(v1, "\n")

			fmt.Println(v)
			fmt.Println(section)
		}
	}

	return val
}
