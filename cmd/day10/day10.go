package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Try the minimum number of buttons first
// First wave - try every button once
// Second wave - try every combination of two buttons
// every combo of three. If you find a solution at one stage, stop and add that to the total

type Machine struct {
	EndState []string
	Buttons [][]int
	Joltage []int
}

func readInput(filepath string) []Machine {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	machines := make([]Machine, 0)
	
	for sc.Scan() {
		arr := strings.Split(sc.Text(), " ")

		// Save End state as an array of strings
		endState := arr[0]
		end := len(endState) - 1
		trimmedEndState := endState[1:end]
		EndState := make([]string, 0)
		for _, str := range trimmedEndState {
			EndState = append(EndState, string(str))
		}

		// Get Buttons
		buttons := arr[1:len(arr)-1]
		Buttons := make([][]int, 0)
		for _, button := range buttons {
			buttonTrimmed := button[1:len(button)-1]
			buttonArr := strings.Split(buttonTrimmed, ",")
			buttonInt := make([]int, len(buttonArr))
			for i, str := range buttonArr {
				num, _ := strconv.Atoi(str)
				buttonInt[i] = num
			}
			Buttons = append(Buttons, buttonInt)
		}

		// Get Joltage
		joltage := arr[len(arr) - 1]
		end = len(joltage) - 1
		trimmedJoltage := joltage[1:end]
		joltageArr := strings.Split(trimmedJoltage, ",")
		Joltage := make([]int, len(joltageArr))
		for i, str := range joltageArr {
			num, _ := strconv.Atoi(str)
			Joltage[i] = num
		}

		machines = append(machines, Machine{EndState: EndState, Buttons: Buttons, Joltage: Joltage})
	}

	return machines
}

// Helper function for converting EndState to a bitstring where "." = 0 and "#" = 1
// .##. = 0110
func diagramToMask(diagram []string) uint64 {
	var mask uint64
	for i, c := range diagram {
		if c == "#" {
			mask |= 1 << i // x |= y is shorthand for x = x | y
		}
	}
	return mask
}

// Helper function for converting a button to a bitstring, where each number in the button correlates to a "1"
// (1, 3) on a 4 character endstate would equal "0101"
func buttonToMask(button []int) uint64 {
	var mask uint64
	for _, idx := range button {
		mask |= 1 << idx
	}
	return mask
}

// Parent function for calling DFS
func minPresses(target uint64, buttons []uint64) int {
	n := len(buttons)

	// try subsets of size k = 0..n
	for k := 0; k <= n; k++ {
		if dfs(0, 0, 0, k, target, buttons) {
			return k
		}
	}
	return -1 // should never happen if solvable
}

// Is BFS better
// TODO: Implement in BFS
func dfs(
	start int,
	pressed int,
	state uint64,
	limit int,
	target uint64,
	buttons []uint64,
) bool {
	if pressed == limit {
		return state == target
	}

	for i := start; i < len(buttons); i++ {
		if dfs(i+1, pressed+1, state^buttons[i], limit, target, buttons) {
			return true
		}
	}
	return false
}


// Hints: https://www.reddit.com/r/adventofcode/comments/1pl60n7/2025_day10_part_1_what_is_the_intuition_here_folks/
func MinimumMatches(machines []Machine) int {
	total := 0

	for _, m := range machines {
		target := diagramToMask(m.EndState)

		buttonMasks := make([]uint64, len(m.Buttons))
		for i, b := range m.Buttons {
			buttonMasks[i] = buttonToMask(b)
		}

		total += minPresses(target, buttonMasks)
	}

	return total
}


func main() {
	machines := readInput("./cmd/day10/input.txt")
	fmt.Println(MinimumMatches(machines)) // Part 1 Answer: 535
}