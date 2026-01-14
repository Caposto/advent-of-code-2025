package main

import (
	"advent-of-code/utils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func readInput(filepath string) []Coordinate {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close() 

	sc := bufio.NewScanner(file)
	input := make([]Coordinate, 0)

	for sc.Scan() {
		arr := strings.Split(sc.Text(), ",")

		x, err := strconv.Atoi(arr[0])
		utils.Check(err)
		y, err := strconv.Atoi(arr[1])
		utils.Check(err)

		input = append(input, Coordinate{x, y})
	}

	return input
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func computeArea(p1, p2 Coordinate) int {
	return (1 + Abs(p1.x - p2.x)) * (1 + Abs(p1.y - p2.y))
}

// Part 1 - brute force: for each red tile tile, compute the area of a rectangle to every other red tile
func maxArea(tiles []Coordinate) int {
	res := 0

	for i := 0; i < len(tiles); i++ {
		for j := 0; j < len(tiles); j++ {
			res = max(res, computeArea(tiles[i], tiles[j]))
		}
	}

	return res
}

// Part 2: Find the red and green tiles and store it all in a list
// Find the corners of this big blob - only care about the red tiles on the edge
// Any other red tile in the middle that gets overalapped will not yield a higher area
// TODO: Taking the L on this one

// RemoveDuplicates removes duplicate elements from a slice.
func RemoveDuplicates[T comparable](sliceList []T) []T {
    allKeys := make(map[T]bool) // Use a map as a set to track seen elements
    list := []T{}               // The slice to store unique elements

    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

// Create a grid with all of the red tiles by using Coordinate Compression
// https://medium.com/algorithms-digest/coordinate-compression-2fff95326fb
func createGrid(tiles []Coordinate) ([][]byte, []Coordinate) {
	allX := make([]int, len(tiles))
	allY := make([]int, len(tiles))

	// Separate X and Y coordinates
	for i, c := range tiles {
		allX[i] = c.x 
		allY[i] = c.y
	}

	// Sort and remove duplicate X and Y coordinates
	slices.Sort(allX)
	slices.Sort(allY)
	uniqX := RemoveDuplicates(allX)
	uniqY := RemoveDuplicates(allY)

	xMap := make(map[int]int, len(uniqX))
	yMap := make(map[int]int, len(uniqY))

	for i, x := range uniqX {
		xMap[x] = i
	}

	for i, y := range uniqY {
		yMap[y] = i
	}

	// Create grid
	grid := make([][]byte, len(uniqY))
	for i := range grid {
		grid[i] = make([]byte, len(uniqX))
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	zPoints := make([]Coordinate, 0, len(tiles))

	// Place Red tiles
	for _, p := range tiles {
		y := yMap[p.y]
		x := xMap[p.x]
		grid[y][x] = '#'
		zPoints = append(zPoints, Coordinate{x, y})
	}

	return grid, zPoints
}

// Connect red tiles to form edges
// TODO: Taking the L on this one: https://github.com/sleekmountaincat/aoc2025/blob/main/src/day9/q2.ts
func rasterize(grid [][]byte, zPoints []Coordinate) [][]byte {
	for i := 0; i < len(zPoints); i++ {
		a := zPoints[i]
		b := zPoints[(i + 1) % len(zPoints)] // fancy trick to wrap around to first point

		if a.x == b.x {

		}
	}
	return grid
}

func main() {
	tiles := readInput("./cmd/day9/input2.txt")

	fmt.Println(maxArea(tiles)) // Part 1 Answer: 4782268188

	grid, _ := createGrid(tiles)
	for _, row := range grid {
		fmt.Println(string(row))
	}
}