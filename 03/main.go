package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")

	if err != nil {
		panic(err)
	}

	puzzle1(string(file))
	puzzle2(string(file))
}

func puzzle1(input string) {
	var engineArray []string
	var engineParts []string

	reNumbers := regexp.MustCompile("[0-9]+")
	reNumber := regexp.MustCompile("[0-9]")

	engineArray = append(engineArray, strings.Split(input, "\n")...)

	for i, v := range engineArray {
		matches := reNumbers.FindAllString(v, -1)

		newString := v

		for _, match := range matches {
			index := strings.Index(newString, match)
			length := len(match)
			possible := false
			startIndex := index
			endIndex := startIndex + length

			if index > 0 {
				startIndex -= 1
			}

			if endIndex < (len(newString) - 1) {
				endIndex += 1
			}

			indexToStart := i
			indexToEnd := indexToStart

			if i > 0 {
				indexToStart -= 1
			}

			if len(engineArray) >= indexToEnd+3 {
				indexToEnd += 2
			}

			for _, val := range engineArray[indexToStart:indexToEnd] {
				cleanString := reNumber.ReplaceAllString(val, ".")

				for i := startIndex; i < endIndex; i++ {
					if string(string(cleanString)[i]) != "." {
						possible = true
					}
				}
			}

			if possible {
				engineParts = append(engineParts, match)
			}

			newString = strings.Replace(newString, match, randSeq(len(match)), 1)
		}
	}

	sum := 0
	fmt.Println()

	for _, v := range engineParts {
		int, err := strconv.Atoi(v)

		// fmt.Println(int)

		if err != nil {
			panic(err)
		}

		sum += int
	}

	// println()
	println(sum)
}

func puzzle2(input string) {
	println("")
	println("Puzzle2")
	println("")

	var engineArray []string
	var engineParts []int

	// reAsteriks := regexp.MustCompile("[*]")
	engineArray = append(engineArray, strings.Split(input, "\n")...)

	for i, v := range engineArray {
		asteriksIndex := strings.Index(v, "*")

		if asteriksIndex == -1 {
			continue
		}

		// check previous row
		match1, matched1 := checkPrevRow(engineArray, i, asteriksIndex)
		match2, matched2 := checkNextRow(engineArray, i, asteriksIndex)
		match3, matched3 := checkCurrentRow(engineArray, i, asteriksIndex)

		println(i)
		println(match1, matched1)
		println(match2, matched2)
		println(match3, matched3)

		count := 0

		if matched1 {
			count += 1
		}

		if matched2 {
			count += 1
		}

		if matched3 {
			count += 1
		}

		if count > 1 {

			sum := 0

			if match1 != "" {
				num1, err1 := strconv.Atoi(match1)

				if err1 != nil {
					panic(err1)
				}

				sum = num1
			}

			if match2 != "" {
				num2, err2 := strconv.Atoi(match2)

				if err2 != nil {
					panic(err2)
				}

				sum *= num2
			}

			if match3 != "" {
				num3, err3 := strconv.Atoi(match3)

				if err3 != nil {
					panic(err3)
				}

				sum *= num3
			}

			engineParts = append(engineParts, sum)
		}
	}

	sum := 0
	// fmt.Println()

	for _, v := range engineParts {
		println(v)
		sum += v
	}

	println()
	println(sum)

}

func checkPrevRow(engineArray []string, i int, ai int) (string, bool) {
	if i == 0 {
		return "", false
	}

	row := engineArray[i-1]

	for i := ai - 1; i < ai+2; i++ {
		if 48 <= row[i] && row[i] <= 57 {
			return findNumber(row, string(row[i])), true
		}
	}

	return "", false
}

func checkNextRow(engineArray []string, i int, ai int) (string, bool) {
	if i >= len(engineArray)-1 {
		return "", false
	}

	// println(i, len(engineArray), "here")

	row := engineArray[i+1]

	for i := ai - 2; i < ai+2; i++ {
		if 48 <= row[i] && row[i] <= 57 {
			println(i, row, "a")
			return findNumber(row, string(row[i])), true
		}
	}

	return "", false
}

func checkCurrentRow(engineArray []string, i int, ai int) (string, bool) {
	row := engineArray[i]

	for i := ai - 2; i < ai+2; i++ {
		if 48 <= row[i] && row[i] <= 57 {
			return findNumber(row, string(row[i])), true
		}
	}

	return "", false
}

func findNumber(input string, startDigit string) string {
	index := strings.Index(input, startDigit)

	finalNumber := ""

	for i := index + 1; i < len(input); i++ {
		if input[i] < 48 || input[i] > 57 {
			// println(input[i], string(input[i]), "h2")
			break
		}

		if finalNumber == "" {
			finalNumber = fmt.Sprintf("%s%s", startDigit, string(input[i]))
		} else {
			finalNumber = fmt.Sprintf("%s%s", finalNumber, string(input[i]))
		}
	}

	for i := index - 1; i >= 0; i-- {
		if input[i] < 48 || input[i] > 57 {
			break
		}

		if finalNumber == "" {
			finalNumber = fmt.Sprintf("%s%s", string(input[i]), startDigit)
		} else {
			finalNumber = fmt.Sprintf("%s%s", string(input[i]), finalNumber)
		}
	}

	if finalNumber == "" {
		return startDigit
	}

	return finalNumber
}

func randSeq(n int) string {
	b := make([]string, n)
	for i := range b {
		b[i] = "."
	}
	return string(strings.Join(b, ""))
}
