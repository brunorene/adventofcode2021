package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

// type amphiPods struct {
// 	space    floor
// 	distance int
// }

// type path struct {
// 	start, end, podType int
// }

// func (p path) cost() int {
// 	return int(math.Pow10(p.podType/10-1)) * p.distance()
// }

// func (p path) distance() int {
// 	return len(p.steps())
// }

// func (ap amphiPods) possiblePaths(start, podType int) (paths []path) {
// 	defer timeTrack(time.Now(), "possiblePaths")
// 	paths = []path{}

// 	// if pod is already at home
// 	if start/10 == podType/10 {
// 		if start%podType == 2 || (start%podType == 1 && ap.typesPerPosition[start+1] == podType) {

// 			return
// 		}
// 	}

// 	// in the hallway
// 	if start <= 10 {
// 		if (path{start, podType + 1, podType}).isClear(ap) {
// 			paths = append(paths, path{start, podType + 1, podType})
// 		}
// 		if (path{start, podType + 2, podType}).isClear(ap) {
// 			paths = append(paths, path{start, podType + 2, podType})
// 		}

// 		return
// 	}
// 	// in rooms
// 	_, entryRoomOccupied := ap.typesPerPosition[podType+1]
// 	_, deepRoomOccupied := ap.typesPerPosition[podType+2]
// 	// straight to the end
// 	if !deepRoomOccupied {
// 		if (path{start, podType + 2, podType}).isClear(ap) {
// 			paths = append(paths, path{start, podType + 2, podType})
// 		}

// 		return
// 	}
// 	if !entryRoomOccupied && ap.typesPerPosition[podType+2] == podType {
// 		if (path{start, podType + 1, podType}).isClear(ap) {
// 			paths = append(paths, path{start, podType + 1, podType})
// 		}

// 		return
// 	}
// 	// moving out to the hallway
// 	for _, end := range []int{0, 1, 3, 5, 7, 9, 10} {
// 		if (path{start, end, podType}).isClear(ap) {
// 			paths = append(paths, path{start, end, podType})
// 		}
// 	}

// 	return
// }

// func (p path) isClear(ap amphiPods) bool {
// 	walk := p.steps()

// 	for endPosition := range ap.typesPerPosition {
// 		for _, posPath := range walk {
// 			if endPosition == posPath {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func (ap amphiPods) allHome() bool {
// 	for position, podType := range ap.typesPerPosition {
// 		if position/10 != podType/10 || position <= 10 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (ap amphiPods) String() (out string) {
// 	podStr := func(value int, exists bool) string {
// 		if !exists {
// 			return "."
// 		}
// 		switch value {
// 		case 10:
// 			return "A"
// 		case 20:
// 			return "B"
// 		case 30:
// 			return "C"
// 		case 40:
// 			return "D"
// 		}
// 		panic("unknown type: Strin()")
// 	}
// 	out = "#############\n"
// 	out += "#"
// 	for i := 0; i < 11; i++ {
// 		val, exists := ap.typesPerPosition[i]
// 		out += podStr(val, exists)
// 	}
// 	out += "#\n###"
// 	for _, pos := range []int{11, 21, 31, 41} {
// 		val, exists := ap.typesPerPosition[pos]
// 		out += podStr(val, exists) + "#"
// 	}
// 	out += "##\n  #"
// 	for _, pos := range []int{12, 22, 32, 42} {
// 		val, exists := ap.typesPerPosition[pos]
// 		out += podStr(val, exists) + "#"
// 	}
// 	out += "  \n  #########"

// 	return
// }

// type priorityQueue struct {
// 	queue []*amphiPods
// }

// func (pq priorityQueue) len() int {
// 	return len(pq.queue)
// }

// func (pq *priorityQueue) push(item *amphiPods) {
// 	defer timeTrack(time.Now(), "priority queue push")
// 	pq.queue = append(pq.queue, item)
// 	sort.Slice(pq.queue, func(i, j int) bool { return pq.queue[i].movingPod.cost() > pq.queue[j].movingPod.cost() })
// }

// func (pq *priorityQueue) pop() (item *amphiPods) {
// 	defer timeTrack(time.Now(), "priority queue pop")
// 	item = pq.queue[len(pq.queue)-1]
// 	pq.queue[len(pq.queue)-1] = nil
// 	pq.queue = pq.queue[:len(pq.queue)-1]
// 	return
// }

// func neighbours(ap amphiPods) (nb []*amphiPods) {
// 	defer timeTrack(time.Now(), "neighbours")
// 	for position, podType := range ap.typesPerPosition {
// 		moves := ap.possiblePaths(position, podType)
// 		for _, path := range moves {
// 			candidate := amphiPods{
// 				make(map[int]int),
// 				path,
// 				math.MaxInt,
// 			}
// 			for otherPos, otherType := range ap.typesPerPosition {
// 				candidate.typesPerPosition[otherPos] = otherType
// 			}

// 			delete(candidate.typesPerPosition, position)
// 			candidate.typesPerPosition[path.end] = podType
// 			nb = append(nb, &candidate)
// 		}
// 	}

// 	return
// }

// func dijkstra(start *amphiPods) (distance int) {
// 	next := priorityQueue{}
// 	next.push(start)

// 	for next.len() > 0 {
// 		current := next.pop()

// 		validNeighbours := neighbours(*current)

// 		fmt.Println(current.distance, next.len(), len(validNeighbours))
// 		fmt.Println(current)
// 		for _, neighbour := range validNeighbours {
// 			newDistance := current.distance + neighbour.movingPod.cost()

// 			if newDistance < neighbour.distance {
// 				neighbour.distance = newDistance

// 				if neighbour.allHome() {
// 					fmt.Println("******")
// 					fmt.Println(neighbour.distance)
// 					fmt.Println(neighbour)
// 					return neighbour.distance
// 				}

// 				next.push(neighbour)
// 			}
// 		}
// 	}

// 	panic("no route found")
// }

type floor [3][11]byte

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

func (start floor) pathClear(end floor) bool {
	return true
}

// xx,yy
// 00,00 01,00 02,00 03,00 04,00 05,00 06,00 07,00 08,00 09,00 10,00
// ##### ##### 02,01 ##### 04,01 ##### 06,01 ##### 08,01 ##### #####
// ##### ##### 02,02 ##### 04,02 ##### 06,02 ##### 08,02 ##### #####
func (f floor) neighbours() (out []floor) {
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
							out = append(out, nb)
						}
					}
				}

				// in the hallway or directly to room from other room
				if f[2][(f[y][x]-'A'+1)*2] == 0 {
					nb := f.copySpace(x, y)
					nb[2][(f[y][x]-'A'+1)*2] = f[y][x]
					out = append(out, nb)
				}

				if f[2][(f[y][x]-'A'+1)*2] == 0 {
					nb := f.copySpace(x, y)
					nb[1][(f[y][x]-'A'+1)*2] = f[y][x]
					out = append(out, nb)
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

	start := processInput(lines)

	fmt.Println(start)

	nbs := processInput(lines).neighbours()

	fmt.Printf("**** %d\n", len(nbs))
	fmt.Println()

	for _, nb := range nbs {
		fmt.Println(nb)
	}
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
