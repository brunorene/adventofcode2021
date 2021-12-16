package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type priorityQueue struct {
	queue []*vertex
}

func (pq priorityQueue) len() int {
	return len(pq.queue)
}

func (pq *priorityQueue) push(item *vertex) {
	pq.queue = append(pq.queue, item)
	sort.Slice(pq.queue, func(i, j int) bool { return pq.queue[i].distance > pq.queue[j].distance })
}

func (pq *priorityQueue) pop() (item *vertex) {
	item = pq.queue[len(pq.queue)-1]
	pq.queue[len(pq.queue)-1] = nil
	pq.queue = pq.queue[:len(pq.queue)-1]
	return
}

func increase(in string, sum int) (out string) {
	for _, c := range in {
		v, err := strconv.Atoi(string(c))
		check(err)
		v += sum
		if v > 9 {
			v -= 9
		}
		out = fmt.Sprintf("%s%d", out, v)
	}

	return
}

func bigMap(in []string, size int) (out []string) {
	for v := 0; v < size; v++ {
		for _, line := range in {
			out = append(out, increase(line, v))
		}
	}

	for idx, line := range out {
		for h := 1; h < size; h++ {
			out[idx] += increase(line, h)
		}
	}

	return
}

type vertex struct {
	x, y, risk, distance int
}

type riskGrid [][]*vertex

func processInput(lines []string) (grid riskGrid) {
	for y := 0; y < len(lines); y++ {
		grid = append(grid, []*vertex{})
		for x := 0; x < len(lines[y]); x++ {
			val, err := strconv.Atoi(string(lines[y][x]))
			check(err)
			grid[y] = append(grid[y], &vertex{x, y, val, math.MaxInt})
		}
	}

	return
}

func (g riskGrid) upVertex(v vertex) *vertex {
	if v.y == 0 {
		return nil
	}

	return g[v.y-1][v.x]
}

func (g riskGrid) leftVertex(v vertex) *vertex {
	if v.x == 0 {
		return nil
	}

	return g[v.y][v.x-1]
}

func (g riskGrid) downVertex(v vertex) *vertex {
	if v.y == len(g)-1 {
		return nil
	}

	return g[v.y+1][v.x]
}

func (g riskGrid) rightVertex(v vertex) *vertex {
	if v.x == len(g[0])-1 {
		return nil
	}

	return g[v.y][v.x+1]
}

func (g riskGrid) neighours(v vertex) (nb []*vertex) {
	for _, v := range []*vertex{g.upVertex(v), g.downVertex(v), g.rightVertex(v), g.leftVertex(v)} {
		if v != nil {
			nb = append(nb, v)
		}
	}

	return
}

func (g riskGrid) isDestination(v vertex) bool {
	return v.y == len(g)-1 && v.x == len(g[0])-1
}

func (g riskGrid) dijkstra() (distance int) {
	next := priorityQueue{}
	g[0][0].distance = 0
	next.push(g[0][0])

	for next.len() > 0 {
		current := next.pop()

		for _, neighbour := range g.neighours(*current) {
			newDistance := current.distance + neighbour.risk

			if newDistance < neighbour.distance {
				neighbour.distance = newDistance

				if g.isDestination(*neighbour) {
					break
				}

				next.push(neighbour)
			}
		}
	}

	return g[len(g)-1][len(g[0])-1].distance
}

func lowestPath(lines []string) int {
	grid := processInput(lines)

	return grid.dijkstra()
}

func bigLowestPath(lines []string) int {
	big := bigMap(lines, 5)

	grid := processInput(big)

	return grid.dijkstra()
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", lowestPath(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", bigLowestPath(lines))
}

func main() {
	part1()
	part2()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(filename string) (lines []string) {
	_, path, _, _ := runtime.Caller(0)
	dir := strings.ReplaceAll(path, "main.go", "")

	file, err := os.Open(dir + filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return
}
