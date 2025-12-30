package main

import (
	"bufio"
	"fmt"
	"os"
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

// Assumptions:
// Each bank is on its own line in the input file
// Only positive numbers

// Looking for 2 digits from left to right that yield the greatest value
// thinking a two pointer algorithm
// left = first digit, right = second digit
// while l < r and r < length
// if right + 1 >= max, increment right
// if left + 1 >= max, increment left and increment right if l == r

func main() {
	banks := readBanks("./cmd/day3/input2.txt")
	fmt.Println(banks)
}