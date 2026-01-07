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

// Part 1
func cephalopodMathSimple(nums [][]int, ops []string) int {
	result := 0

	for i := range ops {
		prob := 0
		if ops[i] == "*" {
			prob = 1
		}
		for j := range nums {
			switch ops[i] {
				case "+":
					prob += nums[j][i]
				case "*":
					prob *= nums[j][i]
			}
		}
		result += prob
	}

	return result
}

// TODO: Taking the L, I know I have to preserve the original spacing of the input
// Part 2 - Preserve original spacing of input
// Get max length of each column to know how many iterations to run
// Iterate through each column from top to bottom
// If index exists, tack that current digit on to the string
// place new numbers in a matrix from top to bottom and run simple math on them
func cephalopodMathComplex(nums [][]int, ops []string) int {
	result := 0
	lengths := make([]int, len(ops))

	rows := len(nums)
	cols := len(nums[0])
	strs := make([][]string, rows) 
	for i := 0; i < rows; i++ {
		strs[i] = make([]string, cols)
	}

	// Get the "length" of a number (count its digits) and find the max per column
	// This is how many numbers will be created in the new column
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			str := strconv.Itoa(nums[i][j])
			lengths[j] = max(lengths[j], len(str))
			strs[i][j] = str
		}
	}

	for j := 0; j < cols; j++ {
		n := lengths[j] - 1
		tmp := 0
		if ops[j] == "*" {
			tmp = 1
		} 
		for n >= 0 {
			newStr := ""
			for i := 0; i < rows; i++ {
				if n < len(strs[i][j]) {
					newStr += string(strs[i][j][n])
				}
			}
			fmt.Println(newStr)
			newNum, _ := strconv.Atoi(newStr)
			if ops[j] == "*" {
				tmp *= newNum
			} else {
				tmp += newNum
			}
			result += tmp
			n--
		}
	}

	return result
}


func main() {
	nums, ops := readInput("./cmd/day6/input2.txt")
	// fmt.Println(nums)
	// fmt.Println(ops)
	fmt.Println(cephalopodMathSimple(nums, ops)) // Part 1 Answer: 4309240495780

	fmt.Println(cephalopodMathComplex(nums, ops))
}