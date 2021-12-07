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

func getSortedPositions(line string) (numbers []int) {
	for _, val := range strings.Split(line, ",") {
		num, err := strconv.Atoi(val)
		check(err)

		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	return
}

func totalFuel(line string) int {
	numbers := getSortedPositions(line)

	median := numbers[len(numbers)/2-1]

	sum := 0

	for _, n := range numbers {
		sum += int(math.Abs(float64(n - median)))
	}

	return sum
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d", totalFuel(lines[0]))
}

func part2() {
	// lines := readInput("input.txt")
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
