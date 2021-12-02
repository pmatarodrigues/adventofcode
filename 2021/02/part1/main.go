package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

    - forward X increases the horizontal position by X units.
    - down X increases the depth by X units.
    - up X decreases the depth by X units.
	! Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

    - forward 5 adds 5 to your horizontal position, a total of 5.
    - down 5 adds 5 to your depth, resulting in a value of 5.
    - forward 8 adds 8 to your horizontal position, a total of 13.
    - up 3 decreases your depth by 3, resulting in a value of 2.
    - down 8 adds 8 to your depth, resulting in a value of 10.
    - forward 2 adds 2 to your horizontal position, a total of 15.

*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read input file
func readInputFile() []string {
	res, err := os.ReadFile("../input.txt")
	check(err)

	arrayLines := strings.Split(strings.TrimSuffix(string(res), "\n"), "\n")

	return arrayLines
}

func main() {
	depth, horizontal := 0, 0

	// read input file
	submarineActions := readInputFile()

	for _, step := range submarineActions {
		stepArray := strings.Split(step, " ")

		var action string = stepArray[0]
		amount, err := strconv.Atoi(stepArray[1])
		check(err)

		switch action {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}

	}

	fmt.Println("Final position: ", horizontal, "x", depth)
	fmt.Print("Result: ", horizontal*depth)

}
