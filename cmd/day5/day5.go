package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Add all fresh id ranges to an array
// Add all valid IDs into an array
// Iterate through slice and see if they are present in the range
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

// Part 1
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

// Part 2
// Consolidate overlapping ranges into a lsit of new ranges and return the difference
func considerFresh(fresh []Range) (int, []Range) {
	// Sort fresh ranges first by start value to ensure ranges are properly consolidated
	sort.Slice(fresh, func(i, j int) bool {
		return fresh[i].start < fresh[j].start
	})

	newRanges := []Range{fresh[0]} // contains consolidated ranges

	for i := 1; i < len(fresh); i++ {
		addNewRange := true
		start := fresh[i].start 
		end := fresh[i].end

		// compare each existing range to new ranges to see if it can be consolidated
		for j := 0; j < len(newRanges); j++ {
			nr := newRanges[j]
			// In existing range so break
			if inRange(start, nr.start, nr.end) && inRange(end, nr.start, nr.end) {
				addNewRange = false
				break
			}
			// Start is in an existing range, but end is not. Expand the end
			if inRange(start, nr.start, nr.end) && !inRange(end, nr.start, nr.end) {
				addNewRange = false
				newRanges[j].end = end
				break
			}
			// Start is not in an existing range, but end is. Expand the start
			if !inRange(start, nr.start, nr.end) && inRange(end, nr.start, nr.end) {
				addNewRange = false
				newRanges[j].start = start
				break
			}
		}
		// Range did not overlap with any existing range
		if addNewRange {
			newRanges = append(newRanges, Range{start, end})
		}
	}

	// sum the consolidated ranges up
	result := 0
	for _, r := range newRanges {
		result += r.end - r.start + 1 // ranges are inclusive so add a 1 to account for the ends
	}

	return result, newRanges
}

// helper function returns if a value is within start and end
func inRange(val, start, end int) bool {
	return val >= start && val <= end
}

func main() {
	freshIds, validIds := readInput("./cmd/day5/input.txt")
	// fmt.Println("Fresh: ", freshIds)
	// fmt.Println("Valid: ", validIds)

	fmt.Println("Fresh Count: ", countFresh(freshIds, validIds)) // Part 1 Answer: 563

	validFresh, _ := considerFresh(freshIds)
	// fmt.Println("Valid Fresh: ", newRanges)
	fmt.Println("Valid Fresh:", validFresh) // Part 2 Answer: 338693411431456
}