package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Read file directions into an array
// Transform array into integer: Right = +, Left = -
// Use an integer to track the current position on the dial (starting at 50)
// Itertate over arry of directions, incrementing by certain amount
// Account for the rotation case (i.e 99 + 1 should go to 0, 0 - 1 should go to 99)
// if dial > 99, dial = 0 + dial - 99
// After every direction, if the dial is equal to 0, increment counter

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

func main() {
	directions := readDirections("./cmd/day1/input.txt")
	
	// Integer to track the dial's position starting at 50
	dial := 50	
	password := 0

	for _, d := range directions {
		// Part 1
		// dial += d
		// dial = ((dial % 100) + 100) % 100

		// if dial == 0 {
		// 	password++
		// }

		// Part 2
		dial += d
		// fmt.Println("Dial", dial)
		// fmt.Println("Increment", int(math.Floor(float64(dial) / 100.0)))
		password += int(math.Abs(math.Floor(float64(dial) / 100.0)))
		dial = ((dial % 100) + 100) % 100
		fmt.Println("Dial Position", dial)
	}

	fmt.Println(password)
}