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

func main() {
	machines := readInput("./cmd/day10/input2.txt")
	fmt.Println(machines)
}