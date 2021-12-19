package main

import (
	"fmt"
	"math"
)

const (
	MinX = 288
	MaxX = 330
	MinY = -96
	MaxY = -50
)

type motionData struct {
	velocity, finalVelocity, displacement, time int
}

func positionsInsideY(initialVel, min, max int) (matches []motionData) {
	pos := 0
	vel := initialVel
	for t := 1; ; t++ {
		pos += vel
		if pos < min {
			break
		}
		if pos >= min && pos <= max {
			matches = append(matches, motionData{initialVel, vel, pos, t})
		}
		vel--
	}

	return
}

func positionsInsideX(initialVel, min, max int) (matches []motionData) {
	pos := 0
	vel := initialVel
	for t := 1; ; t++ {
		pos += vel
		if pos >= min && pos <= max {
			matches = append(matches, motionData{initialVel, vel, pos, t})
			if vel == 0 {
				for time := t; time < 1000; time++ {
					matches = append(matches, motionData{initialVel, vel, pos, time})
				}
			}
		}
		vel--
		if vel < 0 {
			break
		}
	}

	return
}

func allCandidates(min, max int, inside func(vel, min, max int) []motionData) (candidates map[motionData]struct{}) {
	candidates = make(map[motionData]struct{})
	for v := -1000; v < 1000; v++ {
		data := inside(v, min, max)
		if len(data) > 0 {
			for _, d := range data {
				candidates[d] = struct{}{}
			}
		}
	}

	return
}

func displacement(velocity int, time int) int {
	// v*t - (t²-1)/2
	return velocity*time - (int(math.Pow(float64(time), 2))-time)/2
}

func timeAtTop(velocity int) int {
	// v + 1/2 => floor(v)
	return velocity
}

func highestY(minY, maxY int) (max int) {
	dataFinal := allCandidates(minY, maxY, positionsInsideY)
	for probeY := range dataFinal {
		top := timeAtTop(probeY.velocity)
		pos := displacement(probeY.velocity, top)
		if pos > max {
			max = pos
		}
	}

	return max
}

func distinctInitialVelocity(minX, maxX, minY, maxY int) (distinct map[string]int) {
	distinct = make(map[string]int)
	allX := allCandidates(minX, maxX, positionsInsideX)
	allY := allCandidates(minY, maxY, positionsInsideY)

	for mX := range allX {
		for mY := range allY {
			if mX.time == mY.time {
				distinct[fmt.Sprintf("%d,%d", mX.velocity, mY.velocity)]++
			}
		}
	}

	return
}

func lenDistinctInitialVelocity(minX, maxX, minY, maxY int) int {
	vels := distinctInitialVelocity(minX, maxX, minY, maxY)
	return len(vels)
}

func part1() {
	fmt.Printf("%d\n", highestY(MinY, MaxY))
}

func part2() {
	fmt.Printf("%d\n", lenDistinctInitialVelocity(MinX, MaxX, MinY, MaxY))
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
