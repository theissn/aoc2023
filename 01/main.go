package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`[^0-9 ]+`)

func main() {
	fmt.Println("2023/12/01")

	file, err := os.ReadFile("input-1")

	if err != nil {
		panic(err)
	}

	input := string(file)

	puzzle1(input)
	puzzle2(input)
}

func puzzle1(input string) {
	fmt.Println("Puzzle 1 start")
	sum := 0

	for _, v := range strings.Split(input, "\n") {
		sum += convertNum(getFirstAndLastNumber(v))
	}

	fmt.Println("Sum puzzle 1:", sum)
}

func puzzle2(input string) {
	fmt.Println("Puzzle 2 start")
	sum := 0
	numberStringMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for _, data := range strings.Split(input, "\n") {
		orig := data

		for item, v := range numberStringMap {
			data = strings.ReplaceAll(data, item, fmt.Sprintf("%s%s%s", item, v, item))
		}

		fmt.Println(data, getFirstAndLastNumber(data), orig)
		sum += convertNum(getFirstAndLastNumber(data))
	}

	fmt.Println("Sum puzzle 2:", sum)
}

func convertNum(num string) int {
	if num == "" {
		panic("asd")
	}

	i, err := strconv.Atoi(num)

	if err != nil {
		panic(err)
	}

	return i
}

func getFirstAndLastNumber(data string) string {
	data = re.ReplaceAllString(data, "")

	firstNum := string(data[0])
	lastNum := data[len(data)-1:]

	return fmt.Sprintf("%s%s", firstNum, lastNum)
}
