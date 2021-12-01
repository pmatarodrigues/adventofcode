package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read input file
func readInputFile() string {
	res, err := os.ReadFile("input.txt")
	check(err)

	return string(res)
}

func main() {
	input := readInputFile()

	totalIncreased := 0
	previousValue := 0

	// loop through all values
	for index, value := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		// convert values to int
		intValue, _ := strconv.Atoi(string(value))
		if index == 0 {
			previousValue = intValue
			continue
		}

		// if current value is higher than previous value -> increased
		if intValue > previousValue {
			fmt.Println(intValue, previousValue)
			totalIncreased++
		}
		previousValue = intValue
	}

	fmt.Println(totalIncreased)
}
