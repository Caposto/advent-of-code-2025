package main

import (
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readProductIds(filepath string) []Range {
	file, err := os.Open(filepath)
	check(err)
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
		check(err)

		end, err := strconv.Atoi(s[1])
		check(err)

		r := Range {
			start: start,
			end: end,
		}

		ranges = append(ranges, r)
	}

	return ranges
}

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

func main() {
	ranges := readProductIds("./cmd/day2/input.txt")
	fmt.Println(findInvalidIds(ranges))
}
