package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Iterate through input. Add all fresh ids to a map for constant time lookup map[int]bool
// Add all valid IDs into a slice
// Iterate through slice and see if they are present in the map
// T: O(n), S: O(n)

func readInput(filepath string) (map[int]bool, []int) {
	fresh := make(map[int]bool, 0)
	validIds := make([]int, 0)

	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	// How can I split the text file into 2 sections by the empty line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Process first half of input for fresh ids
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}

		split := strings.Split(scanner.Text(), "-")
		start, err := strconv.Atoi(split[0])
		utils.Check(err)
		end, err := strconv.Atoi(split[1])
		utils.Check(err)

		for i := start; i <= end; i++ {
			fresh[i] = true
		}
	}

	// Process second half of input for valid ids
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		utils.Check(err)
		validIds = append(validIds, id)
	}

	return fresh, validIds
}

func main() {
	freshIds, validIds := readInput("./cmd/day5/input2.txt")
	fmt.Println("Fresh: ", freshIds)
	fmt.Println("Valid: ", validIds)
}