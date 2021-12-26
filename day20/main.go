package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
)

type coords struct {
	x, y int
}

type imageGrid struct {
	algorithm []int
	image     map[coords]int
	cellSum   map[coords]int
}

func calcAlgorithm(lines []string) int {
	var algorithm []int
	for _, c := range lines[0] {
		switch c {
		case '.':
			algorithm = append(algorithm, 0)
		case '#':
			algorithm = append(algorithm, 1)
		}
	}

	image := make(map[coords]int)
	for y, line := range lines[2:] {
		for x, c := range line {
			if c == '#' {
				// where this bit will be used
				image[coords{x, y}] = 1
			}
		}
	}

	// 1st time
	for y := -10; y < 110; y++ {
		for x := -10; x < 110; x++ {
			cellSum := image[coords{x - 1, y - 1}] * int(math.Pow(2, 8))
			cellSum += image[coords{x, y - 1}] * int(math.Pow(2, 7))
			cellSum += image[coords{x + 1, y - 1}] * int(math.Pow(2, 6))
			cellSum += image[coords{x - 1, y}] * int(math.Pow(2, 5))
			cellSum += image[coords{x, y}] * int(math.Pow(2, 4))
			cellSum += image[coords{x + 1, y}] * int(math.Pow(2, 3))
			cellSum += image[coords{x - 1, y + 1}] * int(math.Pow(2, 2))
			cellSum += image[coords{x, y + 1}] * int(math.Pow(2, 1))
			cellSum += image[coords{x + 1, y + 1}] * int(math.Pow(2, 0))

			image[coords{x, y}] = algorithm[cellSum]
		}
	}

	// for y := -11; y < 111; y++ {
	// 	for x := -11; x < 111; x++ {
	// 		if x == -11 || y == -11 || x == 111 || y == 111 {
	// 			image[coords{x, y}] = 1
	// 		}
	// 	}
	// }

	// 2nd time
	for y := -10; y < 110; y++ {
		for x := -10; x < 110; x++ {
			cellSum := image[coords{x - 1, y - 1}]*int(math.Pow(2, 0)) +
				image[coords{x, y - 1}]*int(math.Pow(2, 1)) +
				image[coords{x + 1, y - 1}]*int(math.Pow(2, 2)) +
				image[coords{x - 1, y}]*int(math.Pow(2, 3)) +
				image[coords{x, y}]*int(math.Pow(2, 4)) +
				image[coords{x + 1, y}]*int(math.Pow(2, 5)) +
				image[coords{x - 1, y + 1}]*int(math.Pow(2, 6)) +
				image[coords{x, y + 1}]*int(math.Pow(2, 7)) +
				image[coords{x + 1, y + 1}]*int(math.Pow(2, 8))

			image[coords{x, y}] = algorithm[cellSum]
		}
	}

	// count lights
	var sum int
	for _, v := range image {
		sum += v
	}

	return sum
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", calcAlgorithm(lines))
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
