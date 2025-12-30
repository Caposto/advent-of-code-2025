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

// I don't think this is right
// // Part 1
// // An optimized two pointer version for part 1
// // Time: O(n)
// func optimizedJoltage(banks []string) int {
// 	totalJoltage := 0

// 	for _, bank := range banks {
// 		var l, r int = 0, 1
		
// 		maxVoltage := 0
// 		for r < len(bank) {
// 			// Convert left digit string to integer for pointer comparison
// 			ld, _ := strconv.Atoi(string(bank[l]))

// 			// Concatenate and convert digits to capture current max
// 			var str = string(bank[l]) + string(bank[r])
// 			currentVoltage, _ := strconv.Atoi(str)
// 			maxVoltage = max(maxVoltage, currentVoltage)

// 			if r + 1 == len(bank) {
// 				break
// 			}
// 			next := bank[r + 1]
// 			nextd, _ := strconv.Atoi(string(next))
			
// 			if nextd >= ld {
// 				l = r
// 				r++
// 			} else {
// 				r++
// 			}
// 		}
// 		fmt.Println(maxVoltage)
// 		totalJoltage += maxVoltage
// 	}

// 	return totalJoltage
// }

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
	fmt.Println(bruteForceJoltage(banks))
	fmt.Println(optimizedJoltage(banks))
}