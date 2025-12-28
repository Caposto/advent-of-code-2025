package day1

import (
	"bufio"
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

// Read directions from text file and convert to integer values
// R or "Right" means + so remove the "R"
// L or "Left" means - so replace "L" with a "-"
func readDirections(filepath string) []int {
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var directions []int

	for fileScanner.Scan() {
		text := fileScanner.Text()
		text = strings.ReplaceAll(text, "R", "")
		text = strings.ReplaceAll(text, "L", "-")
		number, err := strconv.Atoi(text)
		check(err)
		directions = append(directions, number)
	}

	return directions
}

// Define global variables to be accessbile by different methods
var dial int = 50
var password int = 0

// Day 1 password
func handleSimplePassword (d int) {
	dial += d
	dial = ((dial % 100) + 100) % 100

	if dial == 0 {
		password++
	}
}

// Day 2 password
func handle0x434C49434BPassword (d int) {
	direction := 1
	if d < 0 {
			direction = -1
	}

	for i := 0; i != d; i += direction {
			dial += direction
			dial = ((dial % 100) + 100) % 100
			if dial == 0 {
					password++
			}
	}
}

// Part 1
func FindSimplePassword() int {
	directions := readDirections("./cmd/day1/input.txt")

	for _, d := range directions {
		// Part 1
		handleSimplePassword(d)
		fmt.Println("Dial Position", dial)
	}

	return password
}

// Part 2
func FindComplexPassword() int {
	directions := readDirections("./cmd/day1/input.txt")

	for _, d := range directions {

		// Part 2
		handle0x434C49434BPassword(d)

		fmt.Println("Dial Position", dial)
	}

	return password
}