package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Turn input into a slice of strings
func readBanks(filepath string) []string {
	file, err := os.Open(filepath)
	check(err)
	defer file.Close()

	banks := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, line)
	}

	return banks
}

// Part 1
// For every bank, iterate through each combination of digits from left to right
// Time: O(n^2)
func bruteForceJoltage(banks []string) int {
	totalJoltage := 0
	for _, bank := range banks {
		maxJoltage := 0
		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				var str = string(bank[i]) + string(bank[j])
				joltage, err := strconv.Atoi(str)

				if err != nil {
					panic(err)
				}

				maxJoltage = max(maxJoltage, joltage)
			}
		}
		fmt.Println(maxJoltage)
		totalJoltage += maxJoltage
	}
	return totalJoltage
}

// Part 2
// Hint: Any 12 digit number that starts with a 9 is larger than all 12 digit numbers that start with 8
// Hint: Monotonic Stack
// Hint: Delete n-12 digits to maximize the result

// Assumptions: 
// the strings are at least 12 digits but can be much longer
// the digits are 1-9 no 0s

type Stack[T any] []T

func (s *Stack[T]) Push(item  T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if s.IsEmpty() {
        var zero T // Return zero value for the type if empty
        return zero, false
    }
    index := len(*s) - 1
    item := (*s)[index]
    *s = (*s)[:index] // Slice off the top element
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if s.IsEmpty() {
        var zero T
        return zero, false
    }
    index := len(*s) - 1
    return (*s)[index], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(*s) == 0
}

// Part 2: choose exactly 12 digits from each bank (in order) to form the largest possible 12-digit number.
// Greedy monotonic stack: drop = len(bank) - 12 digits are allowed to be removed.
func complexJoltage(banks []string) uint64 {
	var totalJoltage uint64 = 0
	const need = 12

	for _, bank := range banks {
		var joltage Stack[int] // use a monotonic stack to maintain the result
		drop := len(bank) - need // number of digits that can be popped from the Stack while ensuring you will have at least 12 digits by the end

		for i := 0; i < len(bank); i++ {
			d := int(bank[i] - '0') // use byte math to get digit

			for drop > 0 && !joltage.IsEmpty() {
				top, _ := joltage.Peek()
				if top >= d {
					break
				}
				joltage.Pop()
				drop--
			}

			joltage.Push(d)
		}
		
		// If not all drops are used, cut off back digits until you have the first 12
		if len(joltage) > need {
			joltage = joltage[:need]
		}

		// Convert stack to a single value
		var val uint64 = 0
		for i := 0; i < need; i++ {
			val = val*10 + uint64(joltage[i])
		}
		totalJoltage += val
	}

	return totalJoltage
}

func main() {
	banks := readBanks("./cmd/day3/input.txt")
	fmt.Println(bruteForceJoltage(banks)) // part 1 answer: 17158

	fmt.Println(complexJoltage(banks)) // part 2 answer: 170449335646486
}