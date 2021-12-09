package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"
)

type charSetList []charSet

func (l charSetList) find(matcher func(charSet) bool) (result charSet, rest charSetList) {
	rest = make(charSetList, 0)
	for _, p := range l {
		if matcher(p) {
			result = p
		} else {
			rest = append(rest, p)
		}
	}

	return result, rest
}

type void struct{}

type charSet map[byte]void

func newCharSet(str string) charSet {
	s := make(charSet)
	for _, c := range str {
		s[byte(c)] = void{}
	}

	return s
}

func (s charSet) key() string {
	var data []byte
	for k := range s {
		data = append(data, k)
	}

	str := strings.Split(string(data), "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func (s charSet) contains(subs charSet) bool {
	for k := range subs {
		_, exists := s[k]
		if !exists {
			return false
		}
	}

	return true
}

func (s1 charSet) minus(s2 charSet) (res charSet) {
	res = make(charSet)
	for k := range s1 {
		_, exists := s2[k]
		if !exists {
			res[k] = void{}
		}
	}

	return
}

type matcher map[string]int

func finalResult(lines []string) (sum int) {
	for _, line := range lines {
		patterns, output := processLine(line)
		matcher := newMatcher(patterns)
		sum += matcher.findNumber(output)
	}

	return
}

func newMatcher(patterns charSetList) (m matcher) {
	one, rest := patterns.find(func(s charSet) bool { return len(s) == 2 })
	four, rest := rest.find(func(s charSet) bool { return len(s) == 4 })
	seven, rest := rest.find(func(s charSet) bool { return len(s) == 3 })
	eight, rest := rest.find(func(s charSet) bool { return len(s) == 7 })
	three, rest := rest.find(func(s charSet) bool { return len(s) == 5 && s.contains(seven) })
	nine, rest := rest.find(func(s charSet) bool { return len(s) == 6 && s.contains(three) })
	five, rest := rest.find(func(s charSet) bool { return len(s) == 5 && s.contains(nine.minus(three)) })
	two, rest := rest.find(func(s charSet) bool { return len(s) == 5 })
	six, rest := rest.find(func(s charSet) bool { return len(s.minus(five)) == 1 })
	zero := rest[0]

	return matcher{
		zero.key():  0,
		one.key():   1,
		two.key():   2,
		three.key(): 3,
		four.key():  4,
		five.key():  5,
		six.key():   6,
		seven.key(): 7,
		eight.key(): 8,
		nine.key():  9,
	}
}

func (m matcher) findNumber(output charSetList) (result int) {
	for idx, out := range output {
		result += m[out.key()] * int(math.Pow10(len(output)-1-idx))
	}

	return
}

func processLine(line string) (patterns charSetList, output charSetList) {
	pair := strings.Split(line, "|")
	for _, p := range strings.Split(strings.TrimSpace(pair[0]), " ") {
		patterns = append(patterns, newCharSet(p))
	}
	for _, o := range strings.Split(strings.TrimSpace(pair[1]), " ") {
		output = append(output, newCharSet(o))
	}

	return
}

func countEasyPatterns(output []charSet) (count int) {
	for _, out := range output {
		if len(out) == 2 || len(out) == 3 || len(out) == 4 || len(out) == 7 {
			count++
		}
	}
	return
}

func countAllEasyPatterns(lines []string) (count int) {
	for _, line := range lines {
		_, output := processLine(line)

		count += countEasyPatterns(output)
	}

	return
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countAllEasyPatterns(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", finalResult(lines))
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
