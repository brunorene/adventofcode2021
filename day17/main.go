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
	velocityX, velocityY, insideRegionX, insideRegionY, time int
}

func displacement(velocity int, time int) int {
	// v*t - (t²-1)/2
	return velocity*time - (int(math.Pow(float64(time), 2))-time)/2
}

func velocity(displacement, time int) (v int, isInteger bool) {
	if (2*displacement+int(math.Pow(float64(time), 2))-time)%(2*time) != 0 {
		return 0, false
	}
	return (2*displacement + int(math.Pow(float64(time), 2)) - time) / (2 * time), true
}

func timeAtTop(velocity int) int {
	// v + 1/2 => floor(v)
	return velocity
}

func between(current, start, end int) bool {
	return current >= start && current <= end
}

func velocityCandidates(min, max int) (data []motionData) {
	return data
}

func candidatesMatchTimeAxis(mainAxis []motionData, match []motionData) (mainMatch map[motionData]int) {
	mainMatch = make(map[motionData]int)
	for _, main := range mainAxis {
		for _, other := range match {
			if main.time == other.time {
				mainMatch[main]++
			}
		}
	}

	return
}

func distinctInitialVelocity(minX, maxX, minY, maxY int) (dataFinal map[motionData]struct{}) {
	dataFinal = make(map[motionData]struct{})
	var dataY []motionData
	for m := minY; m <= maxY; m++ {
		for t := 1; t < 1000; t++ {
			vel, isInt := velocity(m, t)
			if isInt {
				dataY = append(dataY, motionData{
					velocityY:     vel,
					insideRegionY: m,
					time:          t,
				})
			}
		}
	}
	for m := minX; m <= maxX; m++ {
		for t := 1; t < 1000; t++ {
			vel, isInt := velocity(m, t)
			if isInt {
				for _, dataY := range dataY {
					if dataY.time == t {
						dataFinal[motionData{
							velocityX:     vel,
							velocityY:     dataY.velocityY,
							insideRegionX: m,
							insideRegionY: dataY.insideRegionY,
							time:          t,
						}] = struct{}{}
					}
				}
			}
		}
	}

	return
}

func highestY(minX, maxX, minY, maxY int) (max int) {
	dataFinal := distinctInitialVelocity(minX, maxX, minY, maxY)
	for probeY := range dataFinal {
		top := timeAtTop(probeY.velocityY)
		pos := displacement(probeY.velocityY, top)
		if pos > max {
			max = pos
		}
	}

	return max
}

func lenDistinctInitialVelocity(minX, maxX, minY, maxY int) (count int) {
	return len(distinctInitialVelocity(minX, maxX, minY, maxY))
}

func part1() {
	fmt.Printf("%d\n", highestY(MinX, MaxX, MinY, MaxY))
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
