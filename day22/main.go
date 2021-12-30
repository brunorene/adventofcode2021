package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

type regionCube struct {
	minX, maxX, minY, maxY, minZ, maxZ int
	value                              *big.Int
}

func intersection(r1, r2 regionCube, modifier int64) (inters regionCube, exists bool) {
	minX := r1.minX
	maxX := r1.maxX
	minY := r1.minY
	maxY := r1.maxY
	minZ := r1.minZ
	maxZ := r1.maxZ

	var existsMin int
	var existsMax int

	if r1.minX <= r2.minX && r2.minX <= r1.maxX {
		minX = r2.minX
		existsMin++
	}
	if r1.minX <= r2.maxX && r2.maxX <= r1.maxX {
		maxX = r2.maxX
		existsMax++
	}
	if r1.minY <= r2.minY && r2.minY <= r1.maxY {
		minY = r2.minY
		existsMin++
	}
	if r1.minY <= r2.maxY && r2.maxY <= r1.maxY {
		maxY = r2.maxY
		existsMax++
	}
	if r1.minZ <= r2.minZ && r2.minZ <= r1.maxZ {
		minZ = r2.minZ
		existsMin++
	}
	if r1.minZ <= r2.maxZ && r2.maxZ <= r1.maxZ {
		maxZ = r2.maxZ
		existsMax++
	}

	return newCube(minX, maxX, minY, maxY, minZ, maxZ, modifier), existsMax == 3 || existsMin == 3
}

func newCube(minX, maxX, minY, maxY, minZ, maxZ int, modifier int64) regionCube {
	return regionCube{minX, maxX, minY, maxY, minZ, maxZ, big.NewInt(modifier *
		int64(math.Abs(float64(maxX-minX))) *
		int64(math.Abs(float64(maxY-minY))) *
		int64(math.Abs(float64(maxZ-minZ)))),
	}
}

func countAllOns(lines []string) *big.Int {
	var onOff string
	var minX, maxX, minY, maxY, minZ, maxZ int
	var onCubes []regionCube
	var sum *big.Int = big.NewInt(0)
	for _, line := range lines {
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &onOff, &minX, &maxX, &minY, &maxY, &minZ, &maxZ)

		var modifier int64
		if onOff == "on" {
			modifier = 1
		}
		cube := newCube(minX, maxX, minY, maxY, minZ, maxZ, modifier)
		sum.Add(sum, cube.value)
		for _, other := range onCubes {
			inters, exists := intersection(other, cube, -1)
			if exists {
				sum.Add(sum, inters.value)
			}
		}
		if onOff == "on" {
			onCubes = append(onCubes, cube)
		}
	}

	return sum
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countOns(lines))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", countAllOns(lines))
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
