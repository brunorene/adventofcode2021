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

// func timeTrack(start time.Time, name string) {
// 	elapsed := time.Since(start)
// 	fmt.Printf("%s took %s\n", name, elapsed)
// }

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

func dijkstra(start floor) (distance int) {
	next := priorityQueue{}
	next.push(start)
	cacheDistance := make(map[string]int)
	var exists bool

	for next.len() > 0 {
		current := next.pop()

		validNeighbours := current.neighbours()

		fmt.Println(current.distance, next.len(), len(validNeighbours))
		fmt.Println(current)
		for _, neighbour := range validNeighbours {
			neighbour.distance, exists = cacheDistance[neighbour.hash()]
			if !exists {
				neighbour.distance = math.MaxInt
			}

			newDistance := current.distance + neighbour.calculateCost(current)

			if newDistance < neighbour.distance {
				neighbour.distance = newDistance
				cacheDistance[neighbour.hash()] = newDistance

				if neighbour.allHome() {
					fmt.Println("******")
					fmt.Println(neighbour.distance)
					fmt.Println(neighbour)
					return neighbour.distance
				}

				next.push(neighbour)
			}
		}
	}

	panic("no route found")
}

type floor struct {
	grid     [3][11]byte
	cost     int
	distance int
}

func (f floor) hash() (value string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			value += fmt.Sprintf("%d,", f.grid[y][x])
		}
	}

	return
}

func (f floor) copySpace(eraseX, eraseY int) floor {
	newFloor := floor{}

	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if f.grid[y][x] != 0 {
				newFloor.grid[y][x] = f.grid[y][x]
			}
		}
	}

	newFloor.grid[eraseY][eraseX] = 0

	return newFloor
}

func (f floor) String() (out string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if y == 0 {
				switch f.grid[y][x] {
				case 0:
					out += "."
				default:
					out += string(f.grid[y][x])
				}
			} else {
				switch f.grid[y][x] {
				case 0:
					if x%2 == 1 || x == 0 || x == 10 {
						out += "#"
					} else {
						out += "."
					}
				default:
					out += string(f.grid[y][x])
				}
			}
		}
		out += "\n"
	}

	return
}

func (f floor) allHome() bool {
	return f.grid[1][2] == 'A' &&
		f.grid[2][2] == 'A' &&
		f.grid[1][4] == 'B' &&
		f.grid[2][4] == 'B' &&
		f.grid[1][6] == 'C' &&
		f.grid[2][6] == 'C' &&
		f.grid[1][8] == 'D' &&
		f.grid[2][8] == 'D'
}

func (f floor) diff(origin floor) (x1, y1, x2, y2 int, isDiff bool) {
	isDiff = true
	count := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if origin.grid[y][x] != 0 && f.grid[y][x] == 0 { // start
				x1 = x
				y1 = y
				count++
			}
			if origin.grid[y][x] == 0 && f.grid[y][x] != 0 { // end
				x2 = x
				y2 = y
				count++
			}
			if count == 2 {
				return
			}
		}
	}

	isDiff = false

	return
}

func (f floor) pathClear(origin floor) bool {
	x1, y1, x2, y2, isDiff := f.diff(origin)

	if !isDiff {
		return false
	}

	// go up
	if y1 == 2 && origin.grid[1][x1] != 0 {
		return false
	}

	step := 1
	if x2 < x1 {
		step = -1
	}

	fmt.Printf("1: %d,%d 2: %d,%d step: %d\n", x1, y1, x2, y2, step)
	for x := x1 + step; x != x2; x += step {
		if origin.grid[0][x] != 0 {
			return false
		}
	}

	if y2 == 2 {
		if origin.grid[1][x2] != 0 {
			return false
		}
	}

	return origin.grid[y2][x2] == 0
}

func (f *floor) calculateCost(origin floor) int {
	x1, y1, x2, y2, isDiff := f.diff(origin)

	if !isDiff {
		return 0
	}

	f.cost = int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))) *
		int(math.Pow10(int(origin.grid[y1][x1]-'A')))

	return f.cost
}

// xx,yy
// 00,00 01,00 02,00 03,00 04,00 05,00 06,00 07,00 08,00 09,00 10,00
// ##### ##### 02,01 ##### 04,01 ##### 06,01 ##### 08,01 ##### #####
// ##### ##### 02,02 ##### 04,02 ##### 06,02 ##### 08,02 ##### #####
func (f floor) neighbours() (out []floor) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			// pod to be moved
			if f.grid[y][x] != 0 {
				// in the rooms
				if y > 0 {
					// on the right place
					if y == 2 && f.grid[y][x] == byte('A'+(x/2-1)) {
						continue
					}
					if y == 1 && f.grid[y][x] == byte('A'+(x/2-1)) && f.grid[y+1][x] == byte('A'+(x/2-1)) {
						continue
					}
					// blocked
					if y == 2 && f.grid[y-1][x] != 0 {
						continue
					}

					for _, h := range []int{0, 1, 3, 5, 7, 9, 10} {
						if f.grid[0][h] == 0 {
							nb := f.copySpace(x, y)
							nb.grid[0][h] = f.grid[y][x]
							if nb.pathClear(f) {
								out = append(out, nb)
							}
						}
					}
				}

				// in the hallway or directly to room from other room
				if f.grid[2][(f.grid[y][x]-'A'+1)*2] == 0 {
					nb := f.copySpace(x, y)
					nb.grid[2][(f.grid[y][x]-'A'+1)*2] = f.grid[y][x]
					if nb.pathClear(f) {
						out = append(out, nb)
					}
				}

				if f.grid[2][(f.grid[y][x]-'A'+1)*2] == 0 {
					nb := f.copySpace(x, y)
					nb.grid[1][(f.grid[y][x]-'A'+1)*2] = f.grid[y][x]
					if nb.pathClear(f) {
						out = append(out, nb)
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
				f.grid[l-1][idx-1] = byte(c)
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
