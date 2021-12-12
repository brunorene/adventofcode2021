package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var (
	isLower = regexp.MustCompile("[a-z]+")
)

type cave struct {
	name       string
	neighbours map[string]struct{}
}

func newCaves(lines []string) (caves map[string]cave) {
	caves = make(map[string]cave)

	for _, line := range lines {
		names := strings.Split(line, "-")
		c0, exists := caves[names[0]]
		if !exists {
			c0 = cave{
				name:       names[0],
				neighbours: make(map[string]struct{}),
			}
			caves[names[0]] = c0
		}
		c1, exists := caves[names[1]]
		if !exists {
			c1 = cave{
				name:       names[1],
				neighbours: make(map[string]struct{}),
			}
			caves[names[1]] = c1
		}
		c0.neighbours[names[1]] = struct{}{}
		c1.neighbours[names[0]] = struct{}{}
	}

	return
}

func walk(current cave, caves map[string]cave, smallVisited map[string]int, canVisit func(map[string]int) bool) (sum int) {
	switch {
	case current.name == "start":
		for neighbourName := range current.neighbours {
			sum += walk(caves[neighbourName], caves, smallVisited, canVisit)
		}
		return
	case current.name == "end":
		return 1
	case isLower.MatchString(current.name):
		newVisited := make(map[string]int, len(smallVisited))
		for k, v := range smallVisited {
			newVisited[k] = v
		}
		_, exists := smallVisited[current.name]
		if exists {
			if !canVisit(smallVisited) {
				return 0
			}
			newVisited[current.name]++
		} else {
			newVisited[current.name] = 1
		}
		for neighbourName := range current.neighbours {
			if neighbourName != "start" {
				sum += walk(caves[neighbourName], caves, newVisited, canVisit)
			}
		}
		return
	default: // isUpper
		for neighbourName := range current.neighbours {
			if neighbourName != "start" {
				sum += walk(caves[neighbourName], caves, smallVisited, canVisit)
			}
		}
		return
	}
}

func countPaths(lines []string) int {
	caves := newCaves(lines)

	return walk(caves["start"], caves, make(map[string]int), func(m map[string]int) bool {
		return false
	})
}

func countPathsMostTwice(lines []string) int {
	caves := newCaves(lines)

	return walk(caves["start"], caves, make(map[string]int), singleMostTwice)
}

func singleMostTwice(m map[string]int) bool {
	allVisits := 0
	for _, v := range m {
		allVisits += v
	}

	return allVisits == len(m)
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countPaths(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countPathsMostTwice(lines))
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
