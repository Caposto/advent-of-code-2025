package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
)

// Get the input as a list of strings and the length of each row
func readInput(filepath string) ([]string, int) {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	input := make([]string, 0)

	for sc.Scan() {
		input = append(input, sc.Text())
	}

	return input, len(input[0])
}

// Part 1
func countSplits(manifold []string, rowLength int) int {
	splits := 0
	tachyons := make(map[int]bool, 0) // track indices of tachyons

	// Find the starting index in the first row marked by an "S"
	for i, char := range manifold[0] {
		if char == 'S' {
			tachyons[i] = true
		}
	}

	// Iterate starting from the 3rd row (2nd index)
	// update the splits to be j - 1, j + 1 where j is the index of the "^"
	for i := 2; i < len(manifold); i++ {
		for j := 0; j < rowLength; j++ {
			// If a split is encountered by a current tachyon ray
			if manifold[i][j] == '^' && tachyons[j] {
				tachyons[j] = false // remove previous
				splits++
				if j != 0 {
					tachyons[j - 1] = true
				} 
				if j != rowLength - 1 {
					tachyons[j + 1] = true
				}
			}
		}
	}

	return splits
}

// Recursive brute force solution
// Funs for smaller inputs but takes too long for longer inputs
// T: Exponential, O(2^n) where n is the number of splits
func quantum(manifold []string, row int, column int) int {
	// Base case - tachyon reaches the bottom of the manifold
	if row == len(manifold) - 1 {
		return 1
	}
	// When you encounter a split make a recursive call in either direction
	if manifold[row][column] == '^' {
		return quantum(manifold, row, column - 1) + quantum(manifold, row, column + 1)
	}
	// Progress downwards with the current tachyon path
	return quantum(manifold, row + 1, column)
}

type Coord struct {
	row int
	col int
}

var previous map[Coord]int

// Memoization
// T: O(n * m) where n is the number of rows, m is the number of columns
func quantumMemo(manifold []string, row int, column int) int {
	key := Coord{row, column}

	// if r, c in the cache, return the cached result
	if val, ok := previous[key]; ok {
		return val
	}
	
	var result int

	// Base Case
	if row == len(manifold) - 1 {
		result = 1
	} else if manifold[row][column] == '^' {
		result = quantumMemo(manifold, row, column-1) + quantumMemo(manifold, row, column + 1)
	} else {
		result = quantumMemo(manifold, row + 1, column)
	}

	previous[key] = result
	return result
}

func main() {
	manifold, length := readInput("./cmd/day7/input.txt")

	// fmt.Println("Length: ", length)
	// fmt.Println(manifold)

	fmt.Println(countSplits(manifold, length)) // Part 1 Answer: 1555

	
	var start int
	for i, ch := range manifold[0] {
		if ch == 'S' {
			start = i
		}
	}
	// Part 2, brute force
	// fmt.Println(quantum(manifold, 0, start))

	// Part 2, memoized
	previous = make(map[Coord]int, 0)
	fmt.Println(quantumMemo(manifold, 0, start)) // Part 2 Answer: 12895232295789
}