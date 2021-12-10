package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

func neighbours(x, y int, lines []string) (list []byte) {
	if y > 0 {
		list = append(list, lines[y-1][x])
	}
	if y < len(lines)-1 {
		list = append(list, lines[y+1][x])
	}
	if x > 0 {
		list = append(list, lines[y][x-1])
	}
	if x < len(lines[0])-1 {
		list = append(list, lines[y][x+1])
	}

	return
}

func getLowValues(lines []string) (xs, ys []int, vals []byte) {
	for y, line := range lines {
		for x, cell := range line {
			val := byte(cell)
			isMin := true
			for _, n := range neighbours(x, y, lines) {
				if val >= n {
					isMin = false
					break
				}
			}
			if isMin {
				xs = append(xs, x)
				ys = append(ys, y)
				vals = append(vals, val)
			}
		}
	}

	return
}

func sumLowValues(lines []string) (sum int) {
	_, _, vals := getLowValues(lines)
	for _, v := range vals {
		sum += 1 + int(v-'0')
	}

	return
}

func spread(x, y int, lines []string) (size int) {
	size++
	current := lines[y][x]
	ar := []byte(lines[y])
	ar[x] = '9'
	lines[y] = string(ar)
	if x < len(lines[0])-1 && lines[y][x+1] > current && lines[y][x+1] < '9' {
		size += spread(x+1, y, lines)
	}
	if x > 0 && lines[y][x-1] > current && lines[y][x-1] < '9' {
		size += spread(x-1, y, lines)
	}
	if y < len(lines)-1 && lines[y+1][x] > current && lines[y+1][x] < '9' {
		size += spread(x, y+1, lines)
	}
	if y > 0 && lines[y-1][x] > current && lines[y-1][x] < '9' {
		size += spread(x, y-1, lines)
	}
	return size
}

func basins(lines []string) (mult int) {
	xs, ys, _ := getLowValues(lines)
	sizes := []int{}
	for idx := range xs {
		sizes = append(sizes, spread(xs[idx], ys[idx], lines))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return sizes[0] * sizes[1] * sizes[2]
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", sumLowValues(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", basins(lines))
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
