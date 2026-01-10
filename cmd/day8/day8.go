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

type Coord struct {
	x int 
	y int
	z int
}

// Use this to help track connections between cicuits, using the index of each coordinate in the input as its label
type Junction struct {
	label int
	distance float64
}

type JunctionConnection struct {
	start int
	end int
	distance float64
}

func readInput(filepath string) []Coord {
	file, err := os.Open(filepath)
	utils.Check(err)
	defer file.Close()

	sc := bufio.NewScanner(file)
	junctions := make([]Coord, 0)

	for sc.Scan() {
		arr := strings.Split(sc.Text(), ",")
		x, err := strconv.Atoi(arr[0])
		utils.Check(err)

		y, err := strconv.Atoi(arr[1])
		utils.Check(err)

		z, err := strconv.Atoi(arr[2])
		utils.Check(err)

		junctions = append(junctions, Coord{x, y, z})
	}

	return junctions
}

// Compute the distance between 2 x,y,z coordinates
func distance(c1, c2 Coord) float64 {
	dx := float64(c1.x - c2.x)
	dy := float64(c1.y - c2.y)
	dz := float64(c1.z - c2.z)

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Compute the distance of each junction to every other junction sorted from closest to farthest
func allDistances(junctions []Coord) map[int][]Junction {
	distances := make(map[int][]Junction, 0)
	for i, c1 := range junctions {
		for j, c2 := range junctions {
			if i != j { // exclude self
				d := distance(c1, c2)
				distances[i] = append(distances[i], Junction{j, d})
			}
		}
		// Sort from closest to farthest
		slices.SortFunc(distances[i], func(a, b Junction) int {
			return cmp.Compare(a.distance, b.distance)
		})
	}
	return distances
}

// Sort all potential connections from shortest to greatest
func allDistancesV2(junctions []Coord) []JunctionConnection {
	distances := make([]JunctionConnection, 0)

	for i, c1 := range junctions {
		for j := i + 1; j < len(junctions); j++ { // only unique undirected pairs
			c2 := junctions[j]
			d := distance(c1, c2)
			distances = append(distances, JunctionConnection{start: i, end: j, distance: d})
		}
	}

	slices.SortFunc(distances, func(a, b JunctionConnection) int {
		return cmp.Compare(a.distance, b.distance)
	})
	return distances
}

// Form the first 1000 connections
// Than use union-find to identify all of the circuits
// Get the size of the 3 largest circuits and multiply

// Create a helper DSU - Disjoint Union struct
type DSU struct {
	parent []int
	size []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &DSU{parent: p, size: s}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	ra := d.Find(a)
	rb := d.Find(b)
	if ra == rb {
		return false // already in the same circuit (they have the same parent)
	}

	// union by size
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}

// Returns all component sizes (one per root).
func (d *DSU) ComponentSizes() []int {
	counts := make(map[int]int)

	for i := range d.parent {
		r := d.Find(i)
		counts[r]++
	}

	sizes := make([]int, 0, len(counts))
	for _, v := range counts {
		sizes = append(sizes, v)
	}
	return sizes
}

func main() {
	junctions := readInput("./cmd/day8/input.txt")
	edges := allDistancesV2(junctions)

	// Part 1, answer: 57564
	k := 1000 
	if k > len(edges) {
		k = len(edges)
	}

	uf := NewDSU(len(junctions))

	for i := 0; i < k; i++ {
		e := edges[i]
		uf.Union(e.start, e.end)
	}

	sizes := uf.ComponentSizes()
	slices.SortFunc(sizes, func(a, b int) int {
		return cmp.Compare(b, a) // descending order
	})

	if len(sizes) < 3 {
    fmt.Println("Not enough circuits to multiply top 3.")
    fmt.Println("Circuit sizes:", sizes)
    return
	}

	result := sizes[0] * sizes[1] * sizes[2]
	fmt.Println("Top 3 circuit sizes:", sizes[0], sizes[1], sizes[2])
	fmt.Println("Product:", result)
}

