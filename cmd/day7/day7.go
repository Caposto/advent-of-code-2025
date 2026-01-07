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

func main() {
	manifold, length := readInput("./cmd/day7/input.txt")

	// fmt.Println("Length: ", length)
	// fmt.Println(manifold)

	fmt.Println(countSplits(manifold, length)) // Part 1 Answer: 1555
}