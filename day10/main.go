package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

const (
	openLegal  = "([{<"
	closeLegal = ")]}>"
)

type runeStack []rune

func newByteStack() runeStack {
	return runeStack{}
}

func (b *runeStack) pop() (item rune, isEmpty bool) {
	if len(*b) == 0 {
		return 0, true
	}
	lastIndex := len(*b) - 1
	element := (*b)[lastIndex]
	*b = (*b)[:lastIndex]
	return element, false
}

func (b *runeStack) push(item rune) {
	*b = append(*b, item)
}

func errorScore(line string) int {
	scores := []int{3, 57, 1197, 25137}

	st := newByteStack()
	for _, c := range line {
		if strings.ContainsRune(openLegal, c) {
			st.push(c)
		}
		if strings.ContainsRune(closeLegal, c) {
			open, empty := st.pop()
			if empty {
				return 0
			}
			index := strings.IndexRune(closeLegal, c)
			if rune(openLegal[index]) != open {
				return scores[index]
			}
		}
	}

	return 0
}

func totalErrorScore(lines []string) (total int) {
	for _, line := range lines {
		total += errorScore(line)
	}

	return
}

func autocompleteScore(line string) (score int) {
	matcher := regexp.MustCompile(`\[\]|\(\)|\{\}|<>`)

	for {
		before := len(line)
		line = matcher.ReplaceAllString(line, "")
		if before == len(line) {
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		index := strings.IndexByte(openLegal, line[i])
		score *= 5
		score += index + 1
	}

	return
}

func totalAutocompleteScore(lines []string) (total int) {
	var scores []int
	for _, line := range lines {
		if errorScore(line) == 0 {
			scores = append(scores, autocompleteScore(line))
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", totalErrorScore(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", totalAutocompleteScore(lines))
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
