package main

import "fmt"

const (
	P1StartPosition = 1
	P2StartPosition = 9
	P1              = -1
	P2              = 1
)

type gameState struct {
	score int
	position int
}

func play(player int, wins map[int]int64, )

func determPlay(p1Pos, p2Pos int) (result int) {
	scores := map[int]int{P1: 0, P2: 0}
	positions := map[int]int{P1: p1Pos, P2: p2Pos}

	player := P1

	draw := -1

	rolls := 0

	for {
		play := 0
		for i := 0; i < 3; i++ {
			draw++
			draw = draw % 100
			play += draw + 1
			rolls++
		}

		positions[player] = (positions[player] + play) % 10
		scores[player] += positions[player] + 1

		if scores[player] >= 1000 {
			break
		}

		player = -player
	}

	return scores[-player] * rolls
}

func part1() {
	fmt.Printf("%d\n", determPlay(P1StartPosition, P2StartPosition))
}

func part2() {
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
