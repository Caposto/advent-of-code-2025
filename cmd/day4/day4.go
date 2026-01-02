package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
)

// Can assume all strings are the same length and only composed of "@" and "."
// "@" = paper, "." = empty
func readInput(filepath string) []string {
	var input = make([]string, 0)

	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close() 

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		str := scanner.Text()
		input = append(input, str)
	}

	return input
}

type Coordinate struct {
	Row int
	Column int
}

var neighbors = []Coordinate{
    {-1, -1}, // top-left
    {-1,  0}, // top
    {-1,  1}, // top-right
    { 0, -1}, // left
    { 0,  1}, // right
    { 1, -1}, // bottom-left
    { 1,  0}, // bottom
    { 1,  1}, // bottom-right
}

func findMovableRolls(input []string) int {
    result := 0

    for r := 0; r < len(input); r++ {
        for c := 0; c < len(input[r]); c++ {
            if input[r][c] != '@' {
                continue
            }

            count := 0
            for _, d := range neighbors {
                nr := r + d.Row
                nc := c + d.Column
                if nr < 0 || nc < 0 || nr >= len(input) || nc >= len(input[nr]) {
                    continue
                }
                if input[nr][nc] == '@' {
                    count++
                }
            }

            if count < 4 {
                result++
            }
        }
    }

    return result
}

func main() {
	input := readInput("./cmd/day4/input.txt")
	fmt.Println(findMovableRolls(input)) // Part 1 answer: 
}