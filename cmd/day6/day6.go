package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filepath string) ([][]int, []string) {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	nums := make([][]int, 0)
	ops := make([]string, 0)

	scanner := bufio.NewScanner(file)

	var prevLine *string // use this as a buffer to handle last line differently

	for scanner.Scan() {
		line := scanner.Text()

		// If we already have a previous line, process it normally
		if prevLine != nil {
			row := make([]int, 0)
			for _, s := range strings.Fields(*prevLine) {
				n, err := strconv.Atoi(s)
				utils.Check(err)
				row = append(row, n)
			}
			nums = append(nums, row)
		}

		prevLine = &line
	}

	// handle last line of operations
	if prevLine != nil {
		last := *prevLine
		ops = strings.Fields(last)
	}

	return nums, ops
}

// Initial thought, put entire input into a matrix [][]int
// iterate through bouttom row for the sign and than iterate though column for values
func main() {
	nums, ops := readInput("./cmd/day6/input2.txt")
	fmt.Println(nums)
	fmt.Println(ops)
}