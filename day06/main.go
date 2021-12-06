package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func readList(line string) (result []int64) {
	result = make([]int64, 9)
	for _, v := range strings.Split(line, ",") {
		num, err := strconv.Atoi(v)
		check(err)

		result[num]++
	}

	return
}

func grow(counters []int64, days int) []int64 {
	for i := 0; i < days; i++ {
		todayCounters := make([]int64, 9)
		for c := 8; c >= 0; c-- {
			switch c {
			case 0:
				growCount := counters[0]
				todayCounters[8] = growCount
				todayCounters[6] = growCount + counters[7]
			case 7:
			default:
				todayCounters[c-1] = counters[c]
			}
		}
		counters = todayCounters
	}
	return counters
}

func fishCount(line string, days int) (count int64) {
	result := readList(line)
	result = grow(result, days)

	for _, c := range result {
		count += c
	}
	return
}

func part1() {
	lines := readInput("input.txt")

	fmt.Println(fishCount(lines[0], 80))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Println(fishCount(lines[0], 256))
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
