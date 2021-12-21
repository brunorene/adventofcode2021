package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type beacon struct {
	x, y, z   int
	distances []int
	hash      []int
}

func (b beacon) GetHash() (hash string) {
	h := b.hash
	if len(b.hash) == 0 {
		h = b.distances
	}
	for _, i := range h {
		hash += fmt.Sprintf("%04d", i)
	}

	return
}

type scanner struct {
	beacons     []beacon
	sharedCount []int
}

func distance(b1 beacon, b2 beacon) int {
	return int(math.Sqrt(math.Pow(float64(b1.x-b2.x), 2) + math.Pow(float64(b1.y-b2.y), 2) + math.Pow(float64(b1.z-b2.z), 2)))
}

func calculateDistances(in []scanner) []scanner {
	for idxSc := range in {
		for idxB := range in[idxSc].beacons {
			for idxOtherB := range in[idxSc].beacons {
				dist := distance(in[idxSc].beacons[idxB], in[idxSc].beacons[idxOtherB])
				if dist != 0 {
					in[idxSc].beacons[idxB].distances = append(in[idxSc].beacons[idxB].distances, dist)
				}
			}
			sort.Ints(in[idxSc].beacons[idxB].distances)
		}
	}

	return in
}

func traverseBeacons(scanners []scanner) []scanner {
	for i1, sc1 := range scanners {
		for i2, sc2 := range scanners {
			if i1 != i2 {
				for _, b1 := range sc1.beacons {
					for _, b2 := range sc2.beacons {
						match := matchBeacons(b1, b2)
						if len(match) >= 11 {
							sc1.sharedCount[i2]++
						}
					}
				}
			}
		}
	}

	return scanners
}

func matchBeacons(b1, b2 beacon) (match []int) {
	for _, d1 := range b1.distances {
		findIdx := sort.SearchInts(b2.distances, d1)
		if findIdx < len(b2.distances)-1 && b2.distances[findIdx] == d1 {
			match = append(match, d1)
		}
	}

	return
}

func processLines(lines []string) (scanners []scanner) {
	for _, line := range lines {
		if strings.Index(line, "---") == 0 {
			scanners = append(scanners, scanner{})
		}
		if strings.Index(line, ",") > 0 {
			coords := strings.Split(line, ",")
			x, err := strconv.Atoi(coords[0])
			check(err)
			y, err := strconv.Atoi(coords[1])
			check(err)
			z, err := strconv.Atoi(coords[2])
			check(err)

			scanners[len(scanners)-1].beacons = append(scanners[len(scanners)-1].beacons, beacon{x, y, z, []int{}, []int{}})
		}
	}

	for idx := range scanners {
		scanners[idx].sharedCount = make([]int, len(scanners))
	}

	return
}

func part1() {
	lines := readInput("input.txt")

	scanners := traverseBeacons(calculateDistances(processLines(lines)))

	for _, sc := range scanners {
		fmt.Printf("%v ", sc.sharedCount)
		fmt.Println()
	}

	fmt.Printf("%d\n", len(scanners))
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
