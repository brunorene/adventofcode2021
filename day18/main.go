package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func add(pairA string, pairB string) string {
	return completeReduce(fmt.Sprintf("[%s,%s]", pairA, pairB))
}

func pairLevel(pair string, index int) (level int) {
	for i := index - 1; i >= 0; i-- {
		switch pair[i] {
		case '[':
			level++
		case ']':
			level--
		}
	}

	return
}

func completeReduce(pair string) string {
	for {
		reduced := reduce(pair)
		if reduced == pair {
			return reduced
		}
		pair = reduced
	}
}

func reduce(pair string) string {
	numMatch := regexp.MustCompile(`\d+`)
	for idx, c := range pair {
		if c == '[' {
			level := pairLevel(pair, idx)
			if level == 4 {
				substr := pair[idx:]
				closingBracketIndex := strings.Index(substr, "]")
				values := strings.Split(pair[idx+1:idx+closingBracketIndex], ",")
				left, err := strconv.Atoi(values[0])
				check(err)
				right, err := strconv.Atoi(values[1])
				check(err)
				remainingLeft := pair[:idx]
				remainingRight := pair[idx+closingBracketIndex+1:]

				lastOnLeft := numMatch.FindAllStringIndex(remainingLeft, -1)
				firstOnRight := numMatch.FindStringIndex(remainingRight)

				if len(lastOnLeft) > 0 {
					start := lastOnLeft[len(lastOnLeft)-1][0]
					end := lastOnLeft[len(lastOnLeft)-1][1]
					oldLeft, err := strconv.Atoi(remainingLeft[start:end])
					check(err)
					left += oldLeft
					remainingLeft = fmt.Sprintf("%s%d%s", remainingLeft[:start], left, remainingLeft[end:])
				}
				if len(firstOnRight) > 0 {
					start := firstOnRight[0]
					end := firstOnRight[1]
					oldRight, err := strconv.Atoi(remainingRight[start:end])
					check(err)
					right += oldRight
					remainingRight = fmt.Sprintf("%s%d%s", remainingRight[:start], right, remainingRight[end:])
				}

				return remainingLeft + "0" + remainingRight
			}
		}
	}

	numMatch = regexp.MustCompile(`\d\d\d*`)
	doubleDigitNum := numMatch.FindStringIndex(pair)
	if len(doubleDigitNum) == 2 {
		old, err := strconv.Atoi(pair[doubleDigitNum[0]:doubleDigitNum[1]])
		check(err)
		left := old / 2
		right := old/2 + old%2
		return fmt.Sprintf("%s[%d,%d]%s", pair[:doubleDigitNum[0]], left, right, pair[doubleDigitNum[1]:])
	}

	return pair
}

func addLines(lines []string) string {
	current := lines[0]

	for i := 1; i < len(lines); i++ {
		current = add(current, lines[i])
	}

	return current
}

func magnitude(pair string) string {
	pairMatch := regexp.MustCompile(`\[(\d+),(\d+)\]`)
	matches := pairMatch.FindAllStringSubmatchIndex(pair, 1)
	if len(matches) > 0 {
		match := matches[0]
		left, err := strconv.Atoi(pair[match[2]:match[3]])
		check(err)
		right, err := strconv.Atoi(pair[match[4]:match[5]])
		check(err)
		result := 3*left + 2*right
		pair = fmt.Sprintf("%s%d%s", pair[:match[0]], result, pair[match[1]:])
		matches = pairMatch.FindAllStringSubmatchIndex(pair, 1)
	}

	if len(matches) > 0 {
		return magnitude(pair)
	}

	return pair
}

func maxMagnitude(pairs []string) (max int) {
	for a := 0; a < len(pairs); a++ {
		for b := 0; b < len(pairs); b++ {
			if a != b {
				current, err := strconv.Atoi(magnitude(add(pairs[a], pairs[b])))
				check(err)
				if current > max {
					max = current
				}
			}
		}
	}

	return
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%s\n", magnitude(addLines(lines)))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", maxMagnitude(lines))
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
