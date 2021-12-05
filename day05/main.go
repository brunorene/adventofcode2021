package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var lineMatcher = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

type position struct {
	x, y, value int
}

type floor struct {
	points map[string]*position
	minY   int
	minX   int
	maxY   int
	maxX   int
}

func NewFloor() *floor {
	return &floor{
		points: make(map[string]*position),
		minY:   math.MaxInt,
		minX:   math.MaxInt,
	}
}

func (f *floor) String() (output string) {
	for y := 0; y <= f.maxY; y++ {
		for x := 0; x <= f.maxX; x++ {
			if f.get(x, y) == nil {
				output += " ."
			} else {
				output += fmt.Sprintf("%2d", f.get(x, y).value)
			}
		}
		output += "\n"
	}
	return output
}

func (f *floor) get(x, y int) *position {
	return f.points[fmt.Sprintf("%d,%d", x, y)]
}

func (f *floor) set(x, y int) {
	pos, exists := f.points[fmt.Sprintf("%d,%d", x, y)]
	if !exists {
		pos = &position{
			x:     x,
			y:     y,
			value: 0,
		}
		f.points[fmt.Sprintf("%d,%d", x, y)] = pos
	}
	pos.value++

	if f.minY > y {
		f.minY = y
	}
	if f.minX > x {
		f.minX = x
	}
	if f.maxY < y {
		f.maxY = y
	}
	if f.maxX < x {
		f.maxX = x
	}
}

func (f *floor) allPositions() (result []*position) {
	for _, pos := range f.points {
		result = append(result, pos)
	}

	return
}

func countOverlaps(lines []string, filter func(x1, y1, x2, y2 int) bool) (count int) {
	floor := NewFloor()
	floor.setPoints(lines, filter)

	// fmt.Println(floor)

	for _, p := range floor.allPositions() {
		if p.value > 1 {
			count++
		}
	}

	return
}

func notDiagonal(x1, y1, x2, y2 int) bool {
	return x1 == x2 || y1 == y2
}

func anyLine(x1, y1, x2, y2 int) bool {
	return x1 == x2 || y1 == y2 || math.Abs(float64(x1-x2)) == math.Abs(float64(y1-y2))
}

func (f *floor) setPoints(lines []string, filter func(x1, y1, x2, y2 int) bool) {
	for _, line := range lines {
		matches := lineMatcher.FindAllStringSubmatch(line, -1)
		x1, err := strconv.Atoi(matches[0][1])
		check(err)
		y1, err := strconv.Atoi(matches[0][2])
		check(err)
		x2, err := strconv.Atoi(matches[0][3])
		check(err)
		y2, err := strconv.Atoi(matches[0][4])
		check(err)

		if filter(x1, y1, x2, y2) {
			switch {
			case x2 != x1 && y2 == y1:
				stepX := (x2 - x1) / int(math.Abs(float64(x2-x1)))
				x := x1
				for x != x2 {
					f.set(x, y1)
					x += stepX
				}
				f.set(x2, y1)
			case y2 != y1 && x2 == x1:
				stepY := (y2 - y1) / int(math.Abs(float64(y2-y1)))
				y := y1
				for y != y2 {
					f.set(x1, y)
					y += stepY
				}
				f.set(x1, y2)
			default:
				stepX := (x2 - x1) / int(math.Abs(float64(x2-x1)))
				x := x1
				stepY := (y2 - y1) / int(math.Abs(float64(y2-y1)))
				y := y1
				for y != y2 {
					f.set(x, y)
					x += stepX
					y += stepY
				}
				f.set(x2, y2)

			}
		}
	}
}

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Println(countOverlaps(readInput("input.txt"), notDiagonal))
}

func part2() {
	fmt.Println(countOverlaps(readInput("input.txt"), anyLine))
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
