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

var foldMatcher = regexp.MustCompile(`fold along (x|y)=(\d+)`)

type paper struct {
	grid                   map[string]int
	minX, maxX, minY, maxY int
}

type paperFold struct {
	axis  string
	value int
}

type dot struct {
	x, y, value int
}

func newPaper() *paper {
	return &paper{
		grid: make(map[string]int),
		minX: math.MaxInt,
		minY: math.MaxInt,
	}
}

func (p *paper) allDots() (dots []dot) {
	for k, v := range p.grid {
		coords := strings.Split(k, ",")
		x, err := strconv.Atoi(coords[0])
		check(err)
		y, err := strconv.Atoi(coords[1])
		check(err)
		dots = append(dots, dot{x, y, v})
	}

	return
}

func (p paper) String() (output string) {
	for y := 0; y <= p.maxY; y++ {
		for x := 0; x <= p.maxX; x++ {
			switch p.get(x, y) {
			case 0:
				output += "."
			default:
				output += "#"
			}
		}
		output += "\n"
	}

	return
}

func (p *paper) set(x, y, val int) {
	if val == 0 {
		delete(p.grid, fmt.Sprintf("%d,%d", x, y))

	} else {
		p.grid[fmt.Sprintf("%d,%d", x, y)] = val
	}
	p.maxX = 0
	p.maxY = 0
	p.minX = math.MaxInt
	p.minY = math.MaxInt
	for _, d := range p.allDots() {
		if d.x > p.maxX {
			p.maxX = d.x
		}
		if d.y > p.maxY {
			p.maxY = d.y
		}
		if d.x < p.minX {
			p.minX = d.x
		}
		if d.y < p.minY {
			p.minY = d.y
		}
	}
}

func (p *paper) get(x, y int) int {
	v, exists := p.grid[fmt.Sprintf("%d,%d", x, y)]
	if exists {
		return v
	}
	return 0
}

func (p *paper) fold(f paperFold) {
	// axis == x
	startX := f.value
	startY := 0
	if f.axis == "y" {
		startX = 0
		startY = f.value
	}
	endX := p.maxX
	endY := p.maxY

	for _, d := range p.allDots() {
		if d.x >= startX && d.x <= endX &&
			d.y >= startY && d.y <= endY {
			// axis == x
			newX := f.value - (d.x - f.value)
			newY := d.y
			if f.axis == "y" {
				newY = f.value - (d.y - f.value)
				newX = d.x
			}
			p.set(d.x, d.y, 0)
			p.set(newX, newY, d.value)
		}
	}
}

func (p *paper) countDots() (sum int) {
	for _, d := range p.allDots() {
		sum += p.get(d.x, d.y)
	}

	return
}

func processInput(lines []string) (p *paper, pf []paperFold) {
	var onFolds bool
	p = newPaper()
	for _, line := range lines {
		if len(line) == 0 {
			onFolds = true
			continue
		}
		if onFolds {
			matches := foldMatcher.FindAllStringSubmatch(line, -1)
			val, err := strconv.Atoi(matches[0][2])
			check(err)
			pf = append(pf, paperFold{
				axis:  matches[0][1],
				value: val,
			})
		} else {
			coords := strings.Split(line, ",")
			x, err := strconv.Atoi(coords[0])
			check(err)
			y, err := strconv.Atoi(coords[1])
			check(err)
			p.set(x, y, 1)
		}
	}

	return
}

func countDotsAfter1stFold(lines []string) int {
	p, pf := processInput(lines)

	p.fold(pf[0])

	return p.countDots()
}

func capitalLetters(lines []string) paper {
	p, pf := processInput(lines)

	for _, f := range pf {
		p.fold(f)
	}

	return *p
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countDotsAfter1stFold(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Print(capitalLetters(lines))
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
