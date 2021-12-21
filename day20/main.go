package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
)

type imageGrid struct {
	algorithm              []int
	image                  map[string]int
	minX, maxX, minY, maxY int
}

func processInput(lines []string) (image imageGrid) {
	for _, c := range lines[0] {
		switch c {
		case '.':
			image.algorithm = append(image.algorithm, 0)
		case '#':
			image.algorithm = append(image.algorithm, 1)
		}
	}

	image.image = make(map[string]int)
	image.minY = math.MaxInt
	image.minX = math.MaxInt
	for y, line := range lines[2:] {
		for x, c := range line {
			if c == '#' {
				image.image[fmt.Sprintf("%d,%d", x, y)] = 1
				if x < image.minX {
					image.minX = x
				}
				if x > image.maxX {
					image.maxX = x
				}
				if y < image.minY {
					image.minY = y
				}
				if y > image.maxY {
					image.maxY = y
				}
			}
		}
	}

	return
}

func (im *imageGrid) applyAlgorithm() {
	
}

func part1() {
	lines := readInput("input.txt")

	img := processInput(lines)
	fmt.Printf("%+v", img)

	fmt.Printf("%d\n", len(lines))
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
