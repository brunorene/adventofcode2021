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

type priorityQueue struct {
	queue []floor
}

func (pq priorityQueue) len() int {
	return len(pq.queue)
}

func (pq *priorityQueue) push(p floor) {
	pq.queue = append(pq.queue, p)
	sort.Slice(pq.queue, func(i, j int) bool { return pq.queue[i].cost > pq.queue[j].cost })
}

func (pq *priorityQueue) pop() (item floor) {
	item = pq.queue[len(pq.queue)-1]
	pq.queue = pq.queue[:len(pq.queue)-1]
	return
}

func dijkstra(start floor) int {
	next := priorityQueue{}
	next.push(start)
	cacheDistance := make(map[string]int)
	var exists bool

	for next.len() > 0 {
		current := next.pop()

		validNeighbours := current.neighbours()

		fmt.Println(current.level, current.distance, next.len(), len(validNeighbours))
		fmt.Println(current)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		for _, neighbour := range validNeighbours {
			neighbour.level = current.level + 1

			neighbour.distance, exists = cacheDistance[neighbour.hash()]
			if !exists {
				neighbour.distance = math.MaxInt
			}

			newDistance := current.distance + neighbour.calculateCost(current)

			if newDistance < neighbour.distance {
				neighbour.distance = newDistance
				cacheDistance[neighbour.hash()] = newDistance

				if neighbour.allHome() {
					fmt.Println(neighbour.distance, next.len())
					fmt.Println(neighbour)
					return neighbour.distance
				}

				next.push(neighbour)
			}
		}
	}

	panic("no route found")
}

func allPositions() [][2]int {
	return append(hallwayPositions(), roomPositions()...)
}

func hallwayPositions() [][2]int {
	return [][2]int{
		{0, 0}, {1, 0}, {3, 0}, {5, 0}, {7, 0}, {9, 0}, {10, 0},
	}
}

func roomPositions() [][2]int {
	return [][2]int{
		{2, 1}, {2, 2}, {4, 1}, {4, 2}, {6, 1}, {6, 2}, {8, 1}, {8, 2},
	}
}

type floor struct {
	locations map[[2]int]byte
	cost      int
	distance  int
	level     int
}

func (f floor) clone() (other floor) {
	other = floor{
		locations: make(map[[2]int]byte),
	}

	for _, pos := range allPositions() {
		other.locations[pos] = f.locations[pos]
	}

	return
}

func (f floor) hash() (value string) {
	for _, pos := range allPositions() {
		value += fmt.Sprintf("%d,", f.locations[pos])
	}

	return
}

func (f floor) String() (out string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if y == 0 {
				switch f.locations[[2]int{x, y}] {
				case 0:
					out += "."
				default:
					out += string(f.locations[[2]int{x, y}])
				}
			} else {
				switch f.locations[[2]int{x, y}] {
				case 0:
					if x%2 == 1 || x == 0 || x == 10 {
						out += "#"
					} else {
						out += "."
					}
				default:
					out += string(f.locations[[2]int{x, y}])
				}
			}
		}
		out += "\n"
	}

	return
}

func (f floor) allHome() bool {
	return f.locations[[2]int{2, 1}] == 'A' &&
		f.locations[[2]int{2, 2}] == 'A' &&
		f.locations[[2]int{4, 1}] == 'B' &&
		f.locations[[2]int{4, 2}] == 'B' &&
		f.locations[[2]int{6, 1}] == 'C' &&
		f.locations[[2]int{6, 2}] == 'C' &&
		f.locations[[2]int{8, 1}] == 'D' &&
		f.locations[[2]int{8, 2}] == 'D'
}

func (f floor) pathClear(src, dst [2]int) bool {
	count := 0
	for _, xy := range allPositions() {
		if xy[0] > src[0] && xy[0] <= dst[0] &&
			xy[1] >= src[1] && xy[1] <= dst[1] {
			count++
		}
	}

	if count > 0 {
		return false
	}

	return true
}

func (f floor) move(src, dst [2]int) floor {
	next := f.clone()
	next.locations[dst] = f.locations[src]
	next.locations[src] = 0

	return next
}

// xx,yy
// 00,00 01,00 02,00 03,00 04,00 05,00 06,00 07,00 08,00 09,00 10,00
// ##### ##### 02,01 ##### 04,01 ##### 06,01 ##### 08,01 ##### #####
// ##### ##### 02,02 ##### 04,02 ##### 06,02 ##### 08,02 ##### #####
func (f floor) neighbours() (out []floor) {
	for _, xy := range hallwayPositions() {
		if f.locations[xy] != 0 {
			landingX := int((f.locations[xy] - 'A' + 1) * 2)

			for _, y := range []int{2, 1} {
				if f.pathClear(xy, [2]int{landingX, y}) {
					out = append(out, f.move(xy, [2]int{landingX, y}))

					break
				}
			}
		}
	}

	for _, xy := range roomPositions() {
		if f.locations[xy] != 0 {
			landingX := int((f.locations[xy] - 'A' + 1) * 2)

			if landingX == xy[0] && f.locations[[2]int{landingX, 2}] == f.locations[xy] {
				continue
			}

			var straight2Room bool

			for _, y := range []int{2, 1} {
				if f.pathClear(xy, [2]int{landingX, y}) {
					out = append(out, f.move(xy, [2]int{landingX, y}))

					straight2Room = true

					break
				}
			}

			if !straight2Room {
				for _, hallway := range hallwayPositions() {
					if f.pathClear(xy, hallway) {
						out = append(out, f.move(xy, hallway))
					}
				}
			}
		}
	}

	return
}

func processInput(lines []string) (f floor) {
	for _, l := range []int{2, 3} {
		for idx, c := range lines[l] {
			switch c {
			case '#':
			case ' ':
				continue
			default:
				f.locations[[2]int{l - 1, idx - 1}] = byte(c)
			}
		}
	}

	return
}

// Part 1
// #############
// #...........#
// ###A#C#B#C###
//   #D#A#D#B#
//   #########

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", dijkstra(processInput(lines)))
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
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return
}
