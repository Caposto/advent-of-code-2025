package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Iterate through input. Add all fresh id ranges to an array
// Add all valid IDs into a slice
// Iterate through slice and see if they are present in the map
// T: O(n), S: O(n)

type Range struct {
	start int
	end int
}

func readInput(filepath string) ([]Range, []int) {
	fresh := make([]Range, 0)
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

		fresh = append(fresh, Range{start, end})
	}

	// Process second half of input for valid ids
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		utils.Check(err)
		validIds = append(validIds, id)
	}

	return fresh, validIds
}

func countFresh(fresh []Range, valid []int) int {
	result := 0

	for _, id := range valid {
		for _, r := range fresh {
			if id >= r.start && id <= r.end {
				result++
				break // break the loop otherwise it might double count due to overlapping ranges
			}
		}
	}

	return result
}

func main() {
	freshIds, validIds := readInput("./cmd/day5/input.txt")
	// fmt.Println("Fresh: ", freshIds)
	// fmt.Println("Valid: ", validIds)

	fmt.Println("Fresh Count: ", countFresh(freshIds, validIds))
}