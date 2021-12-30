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
	queue []*path
}

func (pq priorityQueue) len() int {
	return len(pq.queue)
}

func (pq *priorityQueue) push(p *path) {
	pq.queue = append(pq.queue, p)
	sort.Slice(pq.queue, func(i, j int) bool { return pq.queue[i].cost() > pq.queue[j].cost() })
}

func (pq *priorityQueue) pop() (item *path) {
	item = pq.queue[len(pq.queue)-1]
	pq.queue[len(pq.queue)-1] = nil
	pq.queue = pq.queue[:len(pq.queue)-1]
	return
}

func dijkstra(start floor) (distance int) {
	next := priorityQueue{}
	next.push(&path{start: start, end: start})
	cacheDistance := make(map[string]int)
	var exists bool

	for next.len() > 0 {
		current := next.pop()

		validNeighbours := (*current).end.neighbours()

		fmt.Println(current.accumulatedDistance, next.len(), len(validNeighbours))
		fmt.Println(current)
		for _, neighbour := range validNeighbours {
			neighbour.accumulatedDistance, exists = cacheDistance[neighbour.hash()]
			if !exists {
				neighbour.accumulatedDistance = math.MaxInt
			}

			newDistance := current.accumulatedDistance + neighbour.cost()

			if newDistance < neighbour.accumulatedDistance {
				neighbour.accumulatedDistance = newDistance
				cacheDistance[neighbour.hash()] = newDistance

				if neighbour.allHome() {
					fmt.Println("******")
					fmt.Println(neighbour.accumulatedDistance)
					fmt.Println(neighbour)
					return neighbour.accumulatedDistance
				}

				next.push(&neighbour)
			}
		}
	}

	panic("no route found")
}

type floor [3][11]byte

type path struct {
	start               floor
	end                 floor
	memoizedCost        *int
	accumulatedDistance int
}

func (p path) hash() (value string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			value += fmt.Sprintf("%d,%d,", p.start[y][x], p.end[y][x])
		}
	}

	return
}

func (f floor) copySpace(eraseX, eraseY int) floor {
	newSpace := floor{}

	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if f[y][x] != 0 {
				newSpace[y][x] = f[y][x]
			}
		}
	}

	newSpace[eraseY][eraseX] = 0

	return newSpace
}

func (f floor) String() (out string) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if y == 0 {
				switch f[y][x] {
				case 0:
					out += "."
				default:
					out += string(f[y][x])
				}
			} else {
				switch f[y][x] {
				case 0:
					if x%2 == 1 || x == 0 || x == 10 {
						out += "#"
					} else {
						out += "."
					}
				default:
					out += string(f[y][x])
				}
			}
		}
		out += "\n"
	}

	return
}

func (p path) String() string {
	return p.end.String()
}

func (p path) allHome() bool {
	return p.end[1][2] == 'A' &&
		p.end[2][2] == 'A' &&
		p.end[1][4] == 'B' &&
		p.end[2][4] == 'B' &&
		p.end[1][6] == 'C' &&
		p.end[2][6] == 'C' &&
		p.end[1][8] == 'D' &&
		p.end[2][8] == 'D'
}

func (p path) diff() (x1, y1, x2, y2 int, isDiff bool) {
	isDiff = true
	count := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if p.start[y][x] != 0 && p.end[y][x] == 0 { // start
				x1 = x
				y1 = y
				count++
			}
			if p.start[y][x] == 0 && p.end[y][x] != 0 { // end
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

func (p path) pathClear() bool {
	x1, y1, x2, y2, isDiff := p.diff()

	if !isDiff {
		return false
	}

	// go up
	if y1 == 2 && p.start[1][x1] != 0 {
		return false
	}

	step := 1
	if x2 < x1 {
		step = -1
	}

	for x := x1 + step; x != x2; x += step {
		if p.start[0][x] != 0 {
			return false
		}
	}

	if y2 == 2 {
		if p.start[1][x2] != 0 {
			return false
		}
	}

	return p.start[y2][x2] == 0
}

func (p *path) cost() int {
	x1, y1, x2, y2, isDiff := p.diff()
	if !isDiff {
		return 0
	}

	if p.memoizedCost == nil {
		val := int(math.Abs(float64(x2-x1))+math.Abs(float64(y2-y1))) *
			int(math.Pow10(int(p.start[y1][x1]-'A')))
		p.memoizedCost = &val
	}

	return *p.memoizedCost
}

// xx,yy
// 00,00 01,00 02,00 03,00 04,00 05,00 06,00 07,00 08,00 09,00 10,00
// ##### ##### 02,01 ##### 04,01 ##### 06,01 ##### 08,01 ##### #####
// ##### ##### 02,02 ##### 04,02 ##### 06,02 ##### 08,02 ##### #####
func (f floor) neighbours() (out []path) {
	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			// pod to be moved
			if f[y][x] != 0 {
				// in the rooms
				if y > 0 {
					// on the right place
					if y == 2 && f[y][x] == byte('A'+(x/2-1)) {
						continue
					}
					if y == 1 && f[y][x] == byte('A'+(x/2-1)) && f[y+1][x] == byte('A'+(x/2-1)) {
						continue
					}
					// blocked
					if y == 2 && f[y-1][x] != 0 {
						continue
					}

					for _, h := range []int{0, 1, 3, 5, 7, 9, 10} {
						if f[0][h] == 0 {
							nb := f.copySpace(x, y)
							nb[0][h] = f[y][x]
							p := path{start: f, end: nb}
							if p.pathClear() {
								out = append(out, p)
							}
						}
					}
				}

				// in the hallway or directly to room from other room
				if f[2][(f[y][x]-'A'+1)*2] == 0 {
					nb := f.copySpace(x, y)
					nb[2][(f[y][x]-'A'+1)*2] = f[y][x]
					p := path{start: f, end: nb}
					if p.pathClear() {
						out = append(out, p)
					}
				}

				if f[2][(f[y][x]-'A'+1)*2] == 0 {
					nb := f.copySpace(x, y)
					nb[1][(f[y][x]-'A'+1)*2] = f[y][x]
					p := path{start: f, end: nb}
					if p.pathClear() {
						out = append(out, p)
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
				f[l-1][idx-1] = byte(c)
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
