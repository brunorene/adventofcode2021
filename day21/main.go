package main

import "fmt"

const (
	P1StartPosition = 2
	P2StartPosition = 10
	P1              = true
	P2              = false
)

type gameState struct {
	score    int
	position int
}

func quantumPlay(player bool, p1Pos int, p1Score int, p2Pos int, p2Score int, roll int) (p1Wins, p2Wins int64) {
	switch player {
	case P1:
		p1Pos = (p1Pos+roll-1)%10 + 1
		p1Score += p1Pos
		if p1Score >= 21 {
			return 1, 0
		}
	case P2:
		p2Pos = (p2Pos+roll-1)%10 + 1
		p2Score += p2Pos
		if p2Score >= 21 {
			return 0, 1
		}
	}

	p1Wins9, p2Wins9 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 9)
	p1Wins8, p2Wins8 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 8)
	p1Wins7, p2Wins7 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 7)
	p1Wins6, p2Wins6 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 6)
	p1Wins5, p2Wins5 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 5)
	p1Wins4, p2Wins4 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 4)
	p1Wins3, p2Wins3 := quantumPlay(!player, p1Pos, p1Score, p2Pos, p2Score, 3)
	return p1Wins9 + p1Wins8*3 + p1Wins7*6 + p1Wins6*7 + p1Wins5*6 + p1Wins4*3 + p1Wins3, p2Wins9 + p2Wins8*3 + p2Wins7*6 + p2Wins6*7 + p2Wins5*6 + p2Wins4*3 + p2Wins3
}

func startGame(p1Start, p2Start int) (p1Wins, p2Wins int64) {
	p1Wins9, p2Wins9 := quantumPlay(P1, p1Start, 0, p2Start, 0, 9)
	p1Wins8, p2Wins8 := quantumPlay(P1, p1Start, 0, p2Start, 0, 8)
	p1Wins7, p2Wins7 := quantumPlay(P1, p1Start, 0, p2Start, 0, 7)
	p1Wins6, p2Wins6 := quantumPlay(P1, p1Start, 0, p2Start, 0, 6)
	p1Wins5, p2Wins5 := quantumPlay(P1, p1Start, 0, p2Start, 0, 5)
	p1Wins4, p2Wins4 := quantumPlay(P1, p1Start, 0, p2Start, 0, 4)
	p1Wins3, p2Wins3 := quantumPlay(P1, p1Start, 0, p2Start, 0, 3)
	return p1Wins9 + p1Wins8*3 + p1Wins7*6 + p1Wins6*7 + p1Wins5*6 + p1Wins4*3 + p1Wins3, p2Wins9 + p2Wins8*3 + p2Wins7*6 + p2Wins6*7 + p2Wins5*6 + p2Wins4*3 + p2Wins3
}

func determPlay(p1Pos, p2Pos int) (result int) {
	scores := map[bool]int{P1: 0, P2: 0}
	positions := map[bool]int{P1: p1Pos, P2: p2Pos}

	player := P1

	draw := 0

	rolls := 0

	for {
		play := 0
		for i := 0; i < 3; i++ {
			draw++
			draw = (draw-1)%100 + 1
			play += draw
			rolls++
		}

		positions[player] = (positions[player]+play-1)%10 + 1
		scores[player] += positions[player]

		if scores[player] >= 1000 {
			break
		}

		player = !player
	}

	return scores[!player] * rolls
}

func part1() {
	fmt.Printf("%d\n", determPlay(P1StartPosition, P2StartPosition))
}

func part2() {
	p1Wins, p2Wins := startGame(P1StartPosition, P2StartPosition)
	fmt.Printf("%d %d\n", p1Wins, p2Wins)
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
