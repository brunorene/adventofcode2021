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

// #############
// #0123456789T#
// ###1#1#1#1###
//   #2#2#2#2#
//   #########
//    1 2 3 4

type amphiPods struct {
	typesPerPosition map[int]int
	movedPodPosition int
	path             []int
	distance         int
}

func (ap amphiPods) possiblePaths(position, podType int) [][]int {
	if position/10 == podType/10 {
		if position%podType == 2 || (position%podType == 1 && ap.typesPerPosition[position+1] == podType) {
			return [][]int{}
		}
	}

	possibleMoveIns := map[int]map[int][][]int{
		10: {
			0:  {{1, 2, 11}, {1, 2, 11, 12}},
			1:  {{2, 11}, {2, 11, 12}},
			3:  {{2, 11}, {2, 11, 12}},
			5:  {{4, 3, 2, 11}, {4, 3, 2, 11, 12}},
			7:  {{6, 5, 4, 3, 2, 11}, {6, 5, 4, 3, 2, 11, 12}},
			9:  {{8, 76, 5, 4, 3, 2, 11}, {8, 7, 6, 5, 4, 3, 2, 11, 12}},
			10: {{9, 8, 76, 5, 4, 3, 2, 11}, {9, 8, 7, 6, 5, 4, 3, 2, 11, 12}}},
		20: {
			0:  {{1, 2, 3, 4, 21}, {1, 2, 3, 4, 21, 22}},
			1:  {{2, 3, 4, 21}, {2, 3, 4, 21, 22}},
			3:  {{4, 21}, {4, 21, 22}},
			5:  {{4, 21}, {4, 21, 22}},
			7:  {{6, 5, 4, 21}, {6, 5, 4, 21, 22}},
			9:  {{8, 7, 6, 5, 4, 21}, {8, 7, 6, 5, 4, 21, 22}},
			10: {{9, 8, 7, 6, 5, 4, 21}, {9, 8, 7, 6, 5, 4, 21, 22}}},
		30: {
			0:  {{1, 2, 3, 4, 5, 6, 31}, {1, 2, 3, 4, 5, 6, 31, 32}},
			1:  {{2, 3, 4, 5, 6, 31}, {2, 3, 4, 5, 6, 31, 32}},
			3:  {{4, 5, 6, 31}, {4, 5, 6, 31, 32}},
			5:  {{6, 31}, {6, 31, 32}},
			7:  {{6, 31}, {6, 31, 32}},
			9:  {{8, 7, 6, 31}, {8, 7, 6, 31, 32}},
			10: {{9, 8, 7, 6, 31}, {9, 8, 7, 6, 31, 32}}},
		40: {
			0:  {{1, 2, 3, 4, 5, 6, 7, 8, 41}, {1, 2, 3, 4, 5, 6, 7, 8, 41, 42}},
			1:  {{2, 3, 4, 5, 6, 7, 8, 41}, {2, 3, 4, 5, 6, 7, 8, 41, 42}},
			3:  {{4, 5, 6, 7, 8, 41}, {4, 5, 6, 7, 8, 41, 42}},
			5:  {{6, 7, 8, 41}, {6, 7, 8, 41, 42}},
			7:  {{8, 41}, {8, 41, 42}},
			9:  {{8, 41}, {8, 41, 42}},
			10: {{9, 8, 41}, {9, 8, 41, 42}}},
	}
	possibleMoveOuts := map[int][][]int{
		12: {
			{11, 2, 1, 0},
			{11, 2, 1},
			{11, 2, 3},
			{11, 2, 3, 4, 5},
			{11, 2, 3, 4, 5, 6, 7},
			{11, 2, 3, 4, 5, 6, 7, 8, 9},
			{11, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		11: {
			{2, 1, 0},
			{2, 1},
			{2, 3},
			{2, 3, 4, 5},
			{2, 3, 4, 5, 6, 7},
			{2, 3, 4, 5, 6, 7, 8, 9},
			{2, 3, 4, 5, 6, 7, 8, 9, 10}},
		22: {
			{21, 4, 3, 2, 1, 0},
			{21, 4, 3, 2, 1},
			{21, 4, 3},
			{21, 4, 5},
			{21, 4, 5, 6, 7},
			{21, 4, 5, 6, 7, 8, 9},
			{21, 4, 5, 6, 7, 8, 9, 10}},
		21: {
			{4, 3, 2, 1, 0},
			{4, 3, 2, 1},
			{4, 3},
			{4, 5},
			{4, 5, 6, 7},
			{4, 5, 6, 7, 8, 9},
			{4, 5, 6, 7, 8, 9, 10}},
		32: {
			{31, 6, 5, 4, 3, 2, 1, 0},
			{31, 6, 5, 4, 3, 2, 1},
			{31, 6, 5, 4, 3},
			{31, 6, 5},
			{31, 6, 7},
			{31, 6, 7, 8, 9},
			{31, 6, 7, 8, 9, 10}},
		31: {
			{6, 5, 4, 3, 2, 1, 0},
			{6, 5, 4, 3, 2, 1},
			{6, 5, 4, 3},
			{6, 5},
			{6, 7},
			{6, 7, 8, 9},
			{6, 7, 8, 9, 10}},
		42: {
			{41, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			{41, 8, 7, 6, 5, 4, 3, 2, 1},
			{41, 8, 7, 6, 5, 4, 3},
			{41, 8, 7, 6, 5},
			{41, 8, 7},
			{41, 8, 9},
			{41, 8, 9, 10}},
		41: {
			{8, 7, 6, 5, 4, 3, 2, 1, 0},
			{8, 7, 6, 5, 4, 3, 2, 1},
			{8, 7, 6, 5, 4, 3},
			{8, 7, 6, 5},
			{8, 7},
			{8, 9},
			{8, 9, 10}},
	}
	if position <= 10 {
		return possibleMoveIns[podType][position]
	}
	return possibleMoveOuts[position]
}

func (ap amphiPods) key() (k string) {
	for i := 0; i < 43; i++ {
		t, exists := ap.typesPerPosition[i]
		if exists {
			k += fmt.Sprintf("%d:%d,", i, t)
		} else {
			k += ".:.,"
		}
	}

	return
}

func (ap amphiPods) cost() int {
	return int(math.Pow10(ap.typesPerPosition[ap.path[len(ap.path)-1]]/10-1)) * len(ap.path)
}

func (ap amphiPods) isClear() bool {
	for position := range ap.typesPerPosition {
		for _, posPath := range ap.path {
			if position == posPath {
				return false
			}
		}
	}

	if ap.path[len(ap.path)-1]%ap.typesPerPosition[ap.movedPodPosition] == 1 { // positions 11, 21, 31 and 41
		for position, podType := range ap.typesPerPosition {
			if position == ap.path[len(ap.path)-1]+1 {
				return podType == ap.typesPerPosition[ap.movedPodPosition]
			}
		}
		return false
	}

	return true
}

func (ap amphiPods) allHome() bool {
	for position, podType := range ap.typesPerPosition {
		if position/10 != podType/10 || position <= 10 {
			return false
		}
	}
	return true
}

func (ap amphiPods) String() (out string) {
	podStr := func(value int, exists bool) string {
		if !exists {
			return "."
		}
		switch value {
		case 10:
			return "A"
		case 20:
			return "B"
		case 30:
			return "C"
		case 40:
			return "D"
		}
		panic("unknown type: Strin()")
	}
	out = "#############\n"
	out += "#"
	for i := 0; i < 11; i++ {
		val, exists := ap.typesPerPosition[i]
		out += podStr(val, exists)
	}
	out += "#\n###"
	for _, pos := range []int{11, 21, 31, 41} {
		val, exists := ap.typesPerPosition[pos]
		out += podStr(val, exists) + "#"
	}
	out += "##\n  #"
	for _, pos := range []int{12, 22, 32, 42} {
		val, exists := ap.typesPerPosition[pos]
		out += podStr(val, exists) + "#"
	}
	out += "  \n  #########"

	return
}

type priorityQueue struct {
	queue []*amphiPods
}

func (pq priorityQueue) len() int {
	return len(pq.queue)
}

func (pq *priorityQueue) push(item *amphiPods) {
	pq.queue = append(pq.queue, item)
	sort.Slice(pq.queue, func(i, j int) bool { return pq.queue[i].cost() > pq.queue[j].cost() })
}

func (pq *priorityQueue) pop() (item *amphiPods) {
	item = pq.queue[len(pq.queue)-1]
	pq.queue[len(pq.queue)-1] = nil
	pq.queue = pq.queue[:len(pq.queue)-1]
	return
}

func neighbours(ap amphiPods) (nb []*amphiPods) {
	for position, podType := range ap.typesPerPosition {
		moves := ap.possiblePaths(position, podType)
		for _, path := range moves {
			candidate := amphiPods{
				make(map[int]int),
				position,
				path,
				math.MaxInt,
			}
			for otherPos, otherType := range ap.typesPerPosition {
				candidate.typesPerPosition[otherPos] = otherType
			}
			if candidate.isClear() {
				delete(candidate.typesPerPosition, position)
				candidate.typesPerPosition[path[len(path)-1]] = podType
				nb = append(nb, &candidate)
			}
		}
	}

	return
}

func dijkstra(start *amphiPods) (distance int) {
	visited := make(map[string]int)
	next := priorityQueue{}
	visited[start.key()] = 0
	next.push(start)

	for next.len() > 0 {
		current := next.pop()
		dist, exists := visited[current.key()]
		if exists {
			current.distance = dist
		}
		fmt.Println(current.distance)
		fmt.Println(current)

		validNeighbours := neighbours(*current)
		for _, neighbour := range validNeighbours {
			dist, exists := visited[neighbour.key()]
			if exists {
				neighbour.distance = dist
			}
			newDistance := current.distance + neighbour.cost()

			if newDistance < neighbour.distance {
				neighbour.distance = newDistance
				visited[neighbour.key()] = newDistance

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

// Part 1
// #############
// #...........#
// ###A#C#B#C###
//   #D#A#D#B#
//   #########

func part1() {
	fmt.Printf("%d\n", dijkstra(&amphiPods{
		typesPerPosition: map[int]int{12: 40, 11: 10, 22: 10, 21: 30, 32: 40, 31: 20, 42: 20, 41: 30},
	}))
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
