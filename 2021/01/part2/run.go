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
	res, err := os.ReadFile("../input.txt")
	check(err)

	return string(res)
}

// sum all values in list
func sumListValues(list []string) int {
	sum := 0
	for _, value := range list {
		// convert string to int
		intValue, _ := strconv.Atoi(string(value))
		sum += intValue
	}

	return sum
}

func main() {
	input := readInputFile()
	inputList := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	var listOfSavedValues []int

	totalIncreased := 0
	previousValue := 0

	// get section sum
	for index, _ := range inputList {
		if index > len(inputList)-3 {
			continue
		}

		listOfSavedValues = append(listOfSavedValues, sumListValues(inputList[index:index+3]))
	}

	// loop through all values
	for index, value := range listOfSavedValues {
		if index == 0 {
			previousValue = value
			continue
		}

		// if current value is higher than previous value -> increased
		if value > previousValue {
			totalIncreased++
		}
		previousValue = value
	}

	fmt.Println(totalIncreased)

}
