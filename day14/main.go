package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"runtime"
	"strings"
)

func processInput(lines []string) (template string, insertions map[string]string) {
	template = lines[0]
	insertions = make(map[string]string)
	matcher := regexp.MustCompile(`([A-Z]{2}) -> ([A-Z])`)

	for i := 2; i < len(lines); i++ {
		matches := matcher.FindAllStringSubmatch(lines[i], -1)
		insertions[matches[0][1]] = matches[0][2]
	}

	return
}

func steps(pair string, insertions map[string]string, level int, finalLevel int, cache map[string]map[rune]int64) (counters map[rune]int64) {
	c, exists := cache[fmt.Sprintf("%s,%d", pair, level)]
	if exists {
		return c
	}

	counters = make(map[rune]int64)

	if finalLevel == level {
		for _, p := range pair {
			counters[p]++
		}

		return
	}

	for _, next := range []string{pair[0:1] + insertions[pair], insertions[pair] + pair[1:]} {
		for k, v := range steps(next, insertions, level+1, finalLevel, cache) {
			counters[k] += v
		}
	}

	cache[fmt.Sprintf("%s,%d", pair, level)] = counters

	return
}

func mostMinusless(lines []string, stepCount int) int64 {
	template, insertions := processInput(lines)
	counters := make(map[rune]int64)
	cache := make(map[string]map[rune]int64)

	for i := 0; i < len(template)-1; i++ {
		for k, v := range steps(template[i:i+2], insertions, 0, stepCount, cache) {
			counters[k] += v
		}
	}

	counters[rune(template[len(template)-1])]++
	counters[rune(template[0])]++

	most := int64(0)
	less := int64(math.MaxInt64)
	for _, v := range counters {
		if v < less {
			less = v
		}
		if v > most {
			most = v
		}
	}

	return most/2 - less/2
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", mostMinusless(lines, 10))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", mostMinusless(lines, 40))
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
