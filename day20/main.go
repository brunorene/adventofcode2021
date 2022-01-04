package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type coords struct {
	x, y int
}

type grid struct {
	algorithm              [512]bool
	cells                  map[coords]bool
	minX, maxX, minY, maxY int
}

func (g grid) swapGrid(frameValue bool) (newG grid) {
	newG.cells = make(map[coords]bool)
	newG.maxX = 0
	newG.maxY = 0
	newG.minX = math.MaxInt
	newG.minY = math.MaxInt
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			var algIndexString string
			for _, offsetY := range []int{-1, 0, 1} {
				for _, offsetX := range []int{-1, 0, 1} {
					if g.cells[coords{x + offsetX, y + offsetY}] {
						algIndexString += "1"
					} else {
						algIndexString += "0"
					}

				}
			}
			algIndex, err := strconv.ParseInt(algIndexString, 2, 16)
			check(err)

			newG.algorithm = g.algorithm

			if newG.algorithm[algIndex] {
				if x < newG.minX {
					newG.minX = x
				}
				if y < newG.minY {
					newG.minY = y
				}
				if x > newG.maxX {
					newG.maxX = x
				}
				if y > newG.maxY {
					newG.maxY = y
				}
				newG.cells[coords{x, y}] = true
			}
		}
	}
	newG.minX--
	newG.minY--
	newG.maxX++
	newG.maxY++
	if frameValue {
		for x := newG.minX - 1; x <= newG.maxX+1; x++ {
			newG.cells[coords{x, newG.minY}] = true
			newG.cells[coords{x, newG.minY - 1}] = true
			newG.cells[coords{x, newG.maxY}] = true
			newG.cells[coords{x, newG.maxY + 1}] = true
		}
		for y := newG.minY - 1; y <= newG.maxY+1; y++ {
			newG.cells[coords{newG.minX, y}] = true
			newG.cells[coords{newG.minX - 1, y}] = true
			newG.cells[coords{newG.maxX, y}] = true
			newG.cells[coords{newG.maxX + 1, y}] = true
		}
	}

	return
}

func (g grid) String() (str string) {
	for y := g.minY - 2; y < g.maxY+2; y++ {
		for x := g.minX - 2; x < g.maxX+2; x++ {
			if g.cells[coords{x, y}] {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}

	return
}

func processInput(lines []string) (g grid) {
	for idx, c := range lines[0] {
		g.algorithm[idx] = c == '#'
	}

	g.cells = make(map[coords]bool)
	g.minX = -1
	g.minY = -1
	g.maxY = len(lines[2:])
	g.maxX = len(lines[2])

	for y, line := range lines[2:] {
		for x, c := range line {
			if c == '#' {
				g.cells[coords{x, y}] = true
			}
		}
	}

	return
}

func part1() {
	lines := readInput("input.txt")

	g := processInput(lines)

	g1 := g.swapGrid(true)
	fmt.Println(g1)
	fmt.Println()

	g2 := g1.swapGrid(false)

	fmt.Println(g2)
	fmt.Println(len(g2.cells))
}

func part2() {
	lines := readInput("input.txt")

	g := processInput(lines)
	fmt.Println(g)

	for i := 0; i < 25; i++ {
		g1 := g.swapGrid(true)
		fmt.Println(g1)
		g2 := g1.swapGrid(false)
		fmt.Println(g2)
		g = g2
	}

	fmt.Println(len(g.cells))
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
