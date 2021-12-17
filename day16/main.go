package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/brunorene/adventofcode2021/day16/packet"
)

func sumNestedVersions(parent packet.Packet) (sum int) {
	for _, p := range parent.GetChildren() {
		switch p.GetType() {
		case packet.Value:
			sum += int(p.GetVersion())
		default:
			sum += sumNestedVersions(p)
		}
	}

	return sum + int(parent.GetVersion())
}

func sumVersions(hex string) (sum int) {
	binary := packet.ToBinary(hex)

	message, _ := packet.NewMessage(binary)

	for _, p := range message {
		switch p.GetType() {
		case packet.Value:
			sum += int(p.GetVersion())
		default:
			sum += sumNestedVersions(p)
		}
	}

	return
}

func part1() {
	lines := readInput("input.txt")

	fmt.Printf("%d\n", sumVersions(lines[0]))
}

func part2() {
	lines := readInput("input.txt")

	fmt.Printf("%v\n", packet.MessageResults(packet.NewMessage(packet.ToBinary(lines[0]))))
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
