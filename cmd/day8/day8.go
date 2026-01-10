package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Sort potential connections by distance
// Graph Problem. [Juntion Box] : [list of its connections]
// Connect the first 1000 together
// How to track circuits? A hashmap?
// Djikstra? DFS, BFS? Any algo that could help out?

// Brute Force:
// Find the distance of each circuit to every other circuit
// Sort all of the shortest distances
// Connect them and keep track of connections in a hashmap

type Coord struct {
	x int 
	y int
	z int
}

func readInput(filepath string) []Coord {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	circuits := make([]Coord, 0)

	for sc.Scan() {
		arr := strings.Split(sc.Text(), ",")
		x, err := strconv.Atoi(arr[0])
		utils.Check(err)

		y, err := strconv.Atoi(arr[1])
		utils.Check(err)

		z, err := strconv.Atoi(arr[2])
		utils.Check(err)

		circuits = append(circuits, Coord{x, y, z})
	}

	return circuits
}

func main() {
	circuits := readInput("./cmd/day8/input2.txt")
	fmt.Println(circuits)
}