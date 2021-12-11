package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type grid [][]byte

func newGrid(lines []string) (g grid) {
	g = make(grid, 10)
	for y, line := range lines {
		g[y] = make([]byte, 10)
		for x, c := range line {
			g[y][x] = byte(c) - '0'
		}
	}

	return
}

func (g grid) String() (output string) {
	for _, line := range g {
		for _, v := range line {
			output += fmt.Sprintf("%d", v)
		}
		output += "\n"
	}

	return
}

func (g grid) allFlash() bool {
	for y, line := range g {
		for x := range line {
			if g[y][x] != 0 {
				return false
			}
		}
	}

	return true
}

func (g grid) step() (flashCount int) {
	// step 1
	for y, line := range g {
		for x := range line {
			g[y][x]++
		}
	}

	// step 2
	hasFlashes := true
	for hasFlashes {
		hasFlashes = false
		for y, line := range g {
			for x := range line {
				if g[y][x] == 10 {
					flashes := g.increaseNeighbours(x, y)
					g[y][x] = 0
					flashCount++
					if !hasFlashes {
						hasFlashes = flashes
					}
				}
			}
		}
	}
	return
}

func doAllSteps(lines []string, steps int) (total int) {
	g := newGrid(lines)

	for i := 0; i < steps; i++ {
		total += g.step()
	}

	return
}

func (g grid) increaseNeighbours(x, y int) bool {
	changed := false
	for _, currY := range []int{y - 1, y, y + 1} {
		for _, currX := range []int{x - 1, x, x + 1} {
			if (currX != x || currY != y) &&
				currX >= 0 && currY >= 0 &&
				currX < len(g[0]) && currY < len(g) &&
				g[currY][currX] > 0 && g[currY][currX] < 10 {
				g[currY][currX]++
				if !changed {
					changed = true
				}
			}
		}
	}

	return changed
}

func syncStep(lines []string) (flashStep int) {
	g := newGrid(lines)

	step := 0
	for {
		step++
		g.step()
		if g.allFlash() {
			return step
		}
	}
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", doAllSteps(lines, 100))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", syncStep(lines))
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
