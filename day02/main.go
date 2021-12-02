package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type direction int

const (
	forward direction = iota
	down
	up
)

type movement struct {
	where direction
	steps int
}

type submarine struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	moveInstance(move1)
	moveInstance(move2)
}

func moveInstance(move func([]movement) submarine) {
	movs := readMovements(readInput("input.txt"))

	subm := move(movs)

	fmt.Printf("%d\n\n", subm.depth*subm.horizontal)
}

func move1(movements []movement) (sub submarine) {
	for _, mov := range movements {
		switch mov.where {
		case forward:
			sub.horizontal += mov.steps
		case up:
			sub.depth -= mov.steps
		case down:
			sub.depth += mov.steps
		}
	}
	return
}

func move2(movements []movement) (sub submarine) {
	for _, mov := range movements {
		switch mov.where {
		case forward:
			sub.horizontal += mov.steps
			sub.depth += sub.aim * mov.steps
		case up:
			sub.aim -= mov.steps
		case down:
			sub.aim += mov.steps
		}
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readMovements(list []string) (movements []movement) {
	for _, item := range list {
		parts := strings.Split(item, " ")
		if len(parts) != 2 {
			panic("invalid movement len != 2 " + item)
		}

		var mov movement
		var err error

		switch parts[0] {
		case "forward":
			mov.where = forward
		case "up":
			mov.where = up
		default:
			mov.where = down
		}

		mov.steps, err = strconv.Atoi(parts[1])
		check(err)

		movements = append(movements, mov)
	}

	return
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
