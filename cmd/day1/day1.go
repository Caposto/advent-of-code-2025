package main

import (
	"bufio"
	"fmt"
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
	defer file.Close()
	check(err)

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
	directions := readDirections("./cmd/day1/input2.txt")
	
	// Integer to track the dial's position starting at 50
	dial := 50	
	password := 0

	for _, d := range directions {
		dial = dial + d
		if dial > 99 {
			dial = dial % 100
		}
		if dial < 0 {
			remainder := dial % -99
			dial = 100 + remainder
		}
		if dial == 0 {
			password++
		}
		fmt.Println("Dial Position", dial)
	}

	fmt.Println(password)
}