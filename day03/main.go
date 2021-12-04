package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	zero rune = '0'
	one  rune = '1'
)

func main() {
	part1()
	part2()
}

func part1() {
	gamma, epsilon := calculateRates(readInput("input.txt"))

	fmt.Printf("%d\n", gamma*epsilon)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func maxAndMin(position int, lines []string) (max rune, min rune) {
	counter := make(map[rune]int)

	for _, line := range lines {
		counter[rune(line[position])]++
	}

	if counter[zero] > counter[one] {
		return zero, one
	}

	return one, zero
}

func calculateRates(lines []string) (gamma int64, epsilon int64) {
	var gammaStr string
	var epsilonStr string

	for idx := range lines[0] {
		max, min := maxAndMin(idx, lines)

		gammaStr += string(max)
		epsilonStr += string(min)
	}

	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	check(err)
	epsilon, err = strconv.ParseInt(epsilonStr, 2, 64)
	check(err)

	return
}

func part2() {
	oxygen, co2 := filterLines(readInput("input.txt"), readInput("input.txt"))

	fmt.Printf("%d", oxygen*co2)
}

func filterLines(oxygenIn []string, co2In []string) (oxygen int64, co2 int64) {
	var oxygenStr string
	var co2Str string

	for position := range oxygenIn[0] {
		oxygenIn, co2In = selectLines(position, oxygenIn, co2In)

		if len(oxygenIn) == 1 {
			oxygenStr = oxygenIn[0]
		}

		if len(co2In) == 1 {
			co2Str = co2In[0]
		}

		if oxygenStr != "" && co2Str != "" {
			break
		}
	}

	oxygen, err := strconv.ParseInt(oxygenStr, 2, 64)
	check(err)
	co2, err = strconv.ParseInt(co2Str, 2, 64)
	check(err)

	return
}

func selectLines(position int, oxygenIn []string, co2In []string) (oxygenOut []string, co2Out []string) {
	max, _ := maxAndMin(position, oxygenIn)
	_, min := maxAndMin(position, co2In)

	for _, line := range oxygenIn {
		if rune(line[position]) == max {
			oxygenOut = append(oxygenOut, line)
		}
	}

	for _, line := range co2In {
		if rune(line[position]) == min {
			co2Out = append(co2Out, line)
		}
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
