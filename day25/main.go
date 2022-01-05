package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type coords struct {
	x, y int
}

type seafloor [][]rune

func newSeaFloor(lines []string) (s seafloor) {
	for _, line := range lines {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		s = append(s, row)
	}

	return
}

func (s seafloor) equals(other seafloor) bool {
	for y, row := range s {
		for x, c := range row {
			if other[y][x] != c {
				return false
			}
		}
	}

	return true
}

func (s seafloor) copy() (cp seafloor) {
	for y, row := range s {
		cp = append(cp, []rune{})
		for _, c := range row {
			cp[y] = append(cp[y], c)
		}
	}

	return
}

func (s seafloor) String() (out string) {

	for _, row := range s {
		for _, c := range row {
			out += string(c)
		}
		out += "\n"
	}

	return
}

func (s seafloor) moveCucumbers() seafloor {
	eastS := s.copy()

	// east
	for y, row := range s {
		for x, c := range row {
			if c == '>' {
				nextX := (x + 1) % len(s[0])
				if s[y][nextX] == '.' {
					eastS[y][x] = '.'
					eastS[y][nextX] = '>'
				}
			}
		}
	}

	southS := eastS.copy()

	// south
	for y, row := range eastS {
		for x, c := range row {
			if c == 'v' {
				nextY := (y + 1) % len(s)
				if eastS[nextY][x] == '.' {
					southS[y][x] = '.'
					southS[nextY][x] = 'v'
				}
			}
		}
	}

	fmt.Println(southS)

	return southS
}

func findStop(lines []string) int {
	s := newSeaFloor(lines)
	fmt.Println(s)

	step := 0

	for {
		other := s.moveCucumbers()

		step++

		if s.equals(other) {
			break
		}

		s = other
	}

	return step
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", findStop(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", len(lines))
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
