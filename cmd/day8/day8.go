package main

import (
	"advent-of-code/utils"
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
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

// Use this to help track connections between cicuits, using the index of each coordinate in the input as its label
type Circuit struct {
	label int
	distance float64
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

// Compute the distance between 2 x,y,z coordinates
func distance(c1, c2 Coord) float64 {
	dx := float64(c1.x - c2.x)
	dy := float64(c1.y - c2.y)
	dz := float64(c1.z - c2.z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Compute the distance of each circuit to every other circuit sorted from closest to farthest
func allDistances(circuits []Coord) map[int][]Circuit {
	distances := make(map[int][]Circuit, 0)
	for i, c1 := range circuits {
		for j, c2 := range circuits {
			d := distance(c1, c2)
			distances[i] = append(distances[i], Circuit{j, d})
		}
		// Sort from closest to farthest
		slices.SortFunc(distances[i], func(a, b Circuit) int {
			return cmp.Compare(a.distance, b.distance)
		})
	}
	return distances
}

func main() {
	circuits := readInput("./cmd/day8/input2.txt")
	d := allDistances(circuits)
	// fmt.Println(circuits)

	fmt.Println(d)
}

