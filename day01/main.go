package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	numbers := readNumbers(readInput("input.txt"))

	fmt.Printf("%d\n\n", countIncreases(numbers))
}

func countIncreases(numbers []int) int {
	count := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			count++
		}
	}

	return count
}

func part2() {
	numbers := readNumbers(readInput("input.txt"))

	fmt.Printf("%d\n", countIncreasesWithSlidingWindow(3, numbers))
}

func countIncreasesWithSlidingWindow(size int, numbers []int) int {
	count := 0
	for i := 1; i < len(numbers)-2; i++ {
		sum0 := numbers[i-1] + numbers[i] + numbers[i+1]
		sum1 := numbers[i] + numbers[i+1] + numbers[i+2]
		if sum1 > sum0 {
			count++
		}
	}

	return count
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

func readNumbers(list []string) (nums []int) {
	for _, item := range list {
		val, err := strconv.Atoi(item)
		check(err)
		nums = append(nums, val)
	}

	return
}
