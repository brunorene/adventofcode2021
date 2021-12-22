package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func countOns(lines []string) int {
	var onOff string
	var minX, maxX, minY, maxY, minZ, maxZ int
	onMap := make(map[string]bool)
	for _, line := range lines {
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &onOff, &minX, &maxX, &minY, &maxY, &minZ, &maxZ)

		if minX >= -50 && maxX <= 50 &&
			minY >= -50 && maxY <= 50 &&
			minZ >= -50 && maxZ <= 50 {

			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					for z := minZ; z <= maxZ; z++ {
						key := fmt.Sprintf("%d,%d,%d", x, y, z)
						switch onOff {
						case "on":
							onMap[key] = true
						case "off":
							delete(onMap, key)
						}
					}
				}
			}
		}
	}

	return len(onMap)
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countOns(lines))
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
