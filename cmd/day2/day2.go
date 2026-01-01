package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read file into an array of tuples: [(11, 22), (95-115)...]

// Invalid ID: any ID which is made only of some sequence of digits repeated TWICE
// Half the digits! Only numbers with an even number of digits are invalid
// Regex?, Convert to string and than compare?, split number in half and subtract (if equals 0 it's invalid)
// Time: O(n), Space: ??

// 11 = [1, 1]. 1 - 1 = 0
// 12 = [1, 2]. 1 - 2 = -1
// 13 = [1, 3]. 1 - 3 = -2

type Range struct {
	start int 
	end int
}

func readProductIds(filepath string) []Range {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var line string

	// There should only be a single line in the input
	if scanner.Scan() {
		line = scanner.Text()
	}

	// Manipulate string into a slice of Range structs
	ranges := []Range{}
	rangeStrings := strings.Split(line, ",")

	for _, rs := range rangeStrings {
		s := strings.Split(rs, "-")
		start, err := strconv.Atoi(s[0])
		utils.Check(err)

		end, err := strconv.Atoi(s[1])
		utils.Check(err)

		r := Range {
			start: start,
			end: end,
		}

		ranges = append(ranges, r)
	}

	return ranges
}

// For each ID in the range, split the in half by digits
// If the difference between the first and second half is 0, the halves are identical
// and the ID is invalid
func findInvalidIds(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			length := len(str)
			mid := length / 2
			// Skip pids with an odd number of digits
			if length % 2 == 1 {
				continue
			} else {
				// divide id in half
				firstHalf := str[:mid]
				secondHalf := str[mid:]

				p1, _ := strconv.Atoi(firstHalf)
				p2, _ := strconv.Atoi(secondHalf)

				if p1 - p2 == 0 {
					fmt.Println("Invalid ID: ", i)
					sum += i
				}
			}
		}
	}
	return sum
}

// Part 2
// Going to have to compare every substring of digits to the rest of the string
// Only have to use substrings up until the middle of the digit: 123456, only have to do 1, 12, 123
// If length of string module length of substring does not equal 0, can skip
// Get length of string, iterate from beginning to midpoint, compare substring to entire string
func findInvalidIdsComplex(ranges []Range) int {
	sum := 0

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			length := len(str)
			mid := length / 2

			for j := 1; j <= mid; j++ {
				if length % j != 0 { // skips substrings that don't evenly fit into string
        	continue
    		}
				substr := str[:j]
				invalid := true

				// Iterate over string using the length of the substring as an iterator
				for k := j; k + j <= length; k += j {
					
					if substr != str[k: k + j] {
						invalid = false
						break 
					}
				}

				if invalid {
					fmt.Println("Invalid ID:", i)
					sum += i
					break // break to prevent duplicates
				}
			}
		}
	}

	return sum
}

func main() {
	ranges := readProductIds("./cmd/day2/input.txt")
	// fmt.Println(findInvalidIds(ranges)) // Answer: 13108371860

	fmt.Println(findInvalidIdsComplex(ranges)) // Answer: 22471660255
}
