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

// a -> 0,.,2,3,.,5,6,7,8,9
// b -> 0,.,.,.,4,5,6,.,8,9
// c -> 0,1,2,3,4,.,.,7,8,9
// d -> .,.,2,3,4,5,6,.,8,9
// e -> 0,.,2,.,.,.,6,.,8,.
// f -> 0,1,.,3,4,5,6,7,8,9
// g -> 0,.,2,3,.,5,6,.,8,9

// 1 2s -> #cf#              = 2s
// 4 4s -> *cf* #bd#         = 4s
// 7 3s -> *cf* #a#          = 3s
// 8 7s -> *                 = 7s
// 3 5s -> *cf* *a* #dg#     = 5s + *a* + *cf*
// 9 6s -> *cf* *a* *dg* #b# = 6s + *cf* + *a* + *dg*

// 0 6s -> *cf* *a*      *b* = 6s + *cf* + *a* + *b*
// 6 6s ->      *a* *dg* *b* = 6s + *a* + *dg* + *b*
// 2 5s -> *a* *dg*          = 5s + *a* + *dg* + !*b*
// 5 5s -> *a* *dg* *b*      = 5s + *a* + *dg* + *b*

type matcher struct {
	patt2int map[string]int
	int2patt []string
}

func finalResult(lines []string) (sum int) {
	for _, line := range lines {
		patterns, output := processLine(line)
		matcher := newMatcher(patterns)
		sum += matcher.findNumber(output)
	}

	return
}

func containsAll(str string, match ...string) bool {
	for _, m := range match {
		if !strings.Contains(str, m) {
			return false
		}
	}

	return true
}

func newMatcher(patterns []string) (m *matcher) {
	m = &matcher{
		patt2int: map[string]int{},
		int2patt: make([]string, 10),
	}

	var segC, segF, segA, segD, segG, segB string

	// find 1
	var rest []string
	for _, p := range patterns {
		if len(p) == 2 {
			m.patt2int[p] = 1
			m.int2patt[1] = p
			segC = string([]byte{p[0]})
			segF = string([]byte{p[1]})
		} else {
			rest = append(rest, p)
		}
	}

	// find 4,7,8
	for idx := len(rest) - 1; idx >= 0; idx-- {
		switch len(rest[idx]) {
		case 4:
			m.patt2int[rest[idx]] = 4
			m.int2patt[4] = rest[idx]
			rest[idx] = rest[len(rest)-1]
			rest = rest[:len(rest)-1]
		case 3:
			m.patt2int[rest[idx]] = 7
			m.int2patt[7] = rest[idx]
			segA = strings.Trim(rest[idx], segC+segF)
			rest[idx] = rest[len(rest)-1]
			rest = rest[:len(rest)-1]
		case 7:
			m.patt2int[rest[idx]] = 8
			m.int2patt[8] = rest[idx]
			rest[idx] = rest[len(rest)-1]
			rest = rest[:len(rest)-1]
		}
	}

	// find 3
	for idx, p := range rest {
		if len(p) == 5 && containsAll(p, segC, segF, segA) {
			m.patt2int[p] = 3
			m.int2patt[3] = p
			seg := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(p, segC, ""), segF, ""), segA, "")
			segD = string([]byte{seg[0]})
			segG = string([]byte{seg[1]})
			rest[idx] = rest[len(rest)-1]
			rest = rest[:len(rest)-1]
			break
		}
	}

	// find 9
	for idx, p := range rest {
		if len(p) == 6 && containsAll(p, segC, segF, segA, segD, segG) {
			m.patt2int[p] = 9
			m.int2patt[9] = p
			segB = strings.Trim(p, segC+segF+segA+segD+segG)
			rest[idx] = rest[len(rest)-1]
			rest = rest[:len(rest)-1]
			break
		}
	}

	// find 0,2,5,6
	for _, p := range rest {
		switch len(p) {
		case 5:
			// find 2
			if containsAll(p, segA, segD, segG) && !strings.Contains(p, segB) {
				m.patt2int[p] = 2
				m.int2patt[2] = p
			}
			// find 5
			if containsAll(p, segA, segB, segD, segG) {
				m.patt2int[p] = 5
				m.int2patt[5] = p
			}
		case 6:
			// find 0
			if containsAll(p, segC, segF, segA, segB) {
				m.patt2int[p] = 0
				m.int2patt[0] = p
			}
			// find 6
			if containsAll(p, segA, segB, segD, segG) {
				m.patt2int[p] = 6
				m.int2patt[6] = p
			}
		}
	}

	return
}

func (m *matcher) findNumber(output []string) (result int) {
	for idx, out := range output {
		result += m.patt2int[out] * int(math.Pow10(len(output)-1-idx))
	}

	return
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func processLine(line string) (patterns []string, output []string) {
	pair := strings.Split(line, "|")
	patterns = strings.Split(strings.TrimSpace(pair[0]), " ")
	for idx, p := range patterns {
		patterns[idx] = sortString(p)
	}
	output = strings.Split(strings.TrimSpace(pair[1]), " ")
	for idx, o := range output {
		output[idx] = sortString(o)
	}

	return
}

func countEasyPatterns(output []string) (count int) {
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
