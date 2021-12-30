package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type alu struct {
	heap       []int
	x, y, w, z int
}

func part1() {
	lines := readInput("input.txt")

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
