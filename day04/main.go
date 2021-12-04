package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type cell struct {
	x, y, value int
	marked      bool
}

type board struct {
	perValue map[int]*cell
	perCoord map[string]*cell
}

func NewBoard() *board {
	return &board{
		perValue: make(map[int]*cell),
		perCoord: make(map[string]*cell),
	}
}

func (b *board) create(x, y, val int) {
	c := cell{x: x, y: y, value: val}

	b.perValue[val] = &c
	b.perCoord[fmt.Sprintf("%d,%d", x, y)] = &c
}

func (b *board) mark(val int) bool {
	c, exists := b.perValue[val]
	if !exists {
		return false
	}

	c.marked = true

	return true
}

func (b *board) isMarked(x, y int) bool {
	c, exists := b.perCoord[fmt.Sprintf("%d,%d", x, y)]
	if !exists {
		return false
	}

	return c.marked
}

func (b *board) getCoords(val int) (x, y int, exists bool) {
	c, exists := b.perValue[val]
	if !exists {
		return 0, 0, false
	}

	return c.x, c.y, true
}

func (b *board) unmarkedValues() (values []int) {
	for v, c := range b.perValue {
		if !c.marked {
			values = append(values, v)
		}
	}

	return
}

func (b *board) getValue(x, y int) (value int, exists bool) {
	c, exists := b.perCoord[fmt.Sprintf("%d,%d", x, y)]
	if !exists {
		return 0, false
	}

	return c.value, true
}

func (b *board) String() (result string) {
	for y := 0; y < 5; y++ {
		line := ""
		for x := 0; x < 5; x++ {
			v, _ := b.getValue(x, y)
			if b.isMarked(x, y) {
				line += "  *"
			} else {
				line += fmt.Sprintf("%3d", v)
			}
		}
		result += fmt.Sprintf("%s\n", line)
	}

	return
}

func (b *board) play(value int) (score int, win bool) {
	exists := b.mark(value)
	if !exists {
		return 0, false
	}

	currentX, currentY, _ := b.getCoords(value)

	hasRow := 0
	for x := 0; x < 5; x++ {
		if b.isMarked(x, currentY) {
			hasRow++
		}
	}

	hasCol := 0
	for y := 0; y < 5; y++ {
		if b.isMarked(currentX, y) {
			hasCol++
		}
	}

	if hasRow == 5 || hasCol == 5 {
		sum := 0
		for _, v := range b.unmarkedValues() {
			sum += v
		}

		return value * sum, true
	}

	return 0, false
}

func main() {
	part1()
	part2()
}

func part1() {
	fmt.Printf("%d\n", bingo(getDrawsAndBoards(readInput("input.txt"))))
}

func part2() {
	fmt.Printf("%d\n", lastBingo(getDrawsAndBoards(readInput("input.txt"))))
}

func lastBingo(draws []int, boards []*board) int {
	for i := 0; i < len(draws); i++ {
		// fmt.Printf("---> %d\n\n", draws[i])

		winnersIdx := []int{}

		for idx, b := range boards {
			score, win := b.play(draws[i])

			// fmt.Println(b)

			if win {
				if len(boards) == 1 || i == len(draws)-1 {
					return score
				}

				winnersIdx = append([]int{idx}, winnersIdx...)
			}
		}

		for _, idx := range winnersIdx {
			boards[idx] = boards[len(boards)-1]
			boards = boards[:len(boards)-1]
		}
	}

	return -1
}

func bingo(draws []int, boards []*board) int {
	for i := 0; i < len(draws); i++ {
		for _, b := range boards {
			score, win := b.play(draws[i])
			if win {
				return score
			}
		}
	}

	return -1
}

func getDrawsAndBoards(lines []string) (draws []int, boards []*board) {
	draws = getDraws(lines[0])

	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, getBoard(lines[i:i+5]))
	}

	return
}

func getBoard(lines []string) (board *board) {
	board = NewBoard()

	sep := regexp.MustCompile(`\s+`)

	for y, line := range lines {
		cells := sep.Split(strings.Trim(line, " "), -1)
		for x, val := range cells {
			num, err := strconv.Atoi(val)
			check(err)

			board.create(x, y, num)
		}
	}

	return
}

func getDraws(line string) (result []int) {
	values := strings.Split(line, ",")

	for _, n := range values {
		num, err := strconv.Atoi(n)
		check(err)

		result = append(result, num)
	}

	return
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
