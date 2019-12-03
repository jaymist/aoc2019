package main

import (
	"math"
	"strconv"

	"go.uber.org/zap"
)

func main() {
	logger = zap.NewExample()
	maxX, maxY := maxOverallGrid()
	logger.Info("max coordinates", zap.Int("max-x", maxX), zap.Int("max-y", maxY))
	makeGrid(maxX, maxY)

	centerY := len(grid) / 2
	centerX := len(grid[0]) / 2

	unravelWires(centerX, centerY)
	scanGrid(centerX, centerY)

	signalDelay(centerX, centerY)

	logger.Info("shortest distance", zap.Float64("distance", distance))
}

func signalDelay(centerX, centerY int) {
	interX, interY, delay := followWire(centerX, centerY, 0, wire1)

}

func followWire(centerX, centerY, delay int, wire []string) (int, int, int) {
	curX := centerX
	curY := centerY
	foundIntersection := false

	for _, movement := range wire {
		direction, size := dirAndSize(movement)

		switch direction {
		case "R":
			for x := curX + 1; x <= (curX + size); x++ {
				delay++
				if grid[curY][x] == 3 {
					foundIntersection = true
					break
				}
			}
			curX += size

		case "L":
			for x := curX - 1; x >= (curX - size); x-- {
				delay++
				if grid[curY][x] == 3 {
					foundIntersection = true
					break
				}
			}
			curX -= size
		case "D":
			for y := curY + 1; y <= (curY + size); y++ {
				delay++
				if grid[y][curX] == 3 {
					foundIntersection = true
					break
				}
			}
			curY += size
		case "U":
			for y := curY - 1; y >= (curY - size); y-- {
				delay++
				if grid[y][curX] == 3 {
					foundIntersection = true
					break
				}
			}
			curY -= size
		}

		if foundIntersection {
			break
		}
	}

	return curX, curY, delay
}

func scanGrid(centerX, centerY int) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 3 {
				xDist := math.Abs(float64(centerX - x))
				yDist := math.Abs(float64(centerY - y))

				dist := math.Abs(float64(xDist + yDist))

				logger.Info(
					"found intersection",
					zap.Int("center-x", centerX),
					zap.Int("center-y", centerY),
					zap.Int("intersect-x", x),
					zap.Int("intersect-y", y),
					zap.Float64("distance-x", xDist),
					zap.Float64("distance-y", yDist),
					zap.Float64("distance", dist),
				)

				distance = math.Min(distance, dist)
			}
		}
	}
}

func unravelWires(centerX, centerY int) {
	unravelWire(centerX, centerY, wire1, 1)
	unravelWire(centerX, centerY, wire2, 2)
}

func unravelWire(centerX, centerY int, wire []string, val int) {
	logger.Info(
		"starting coordinates",
		zap.Int("center-x", centerX),
		zap.Int("center-y", centerY),
	)

	curX := centerX
	curY := centerY

	for _, movement := range wire {
		direction, size := dirAndSize(movement)

		logger.Info(
			"current position",
			zap.Int("curr-x", curX),
			zap.Int("curr-y", curY),
			zap.String("direction", direction),
			zap.Int("size", size),
		)

		switch direction {
		case "R":
			for x := curX + 1; x <= (curX + size); x++ {
				grid[curY][x] += val
			}
			curX += size

		case "L":
			for x := curX - 1; x >= (curX - size); x-- {
				grid[curY][x] += val
			}
			curX -= size
		case "D":
			for y := curY + 1; y <= (curY + size); y++ {
				grid[y][curX] += val
			}
			curY += size
		case "U":
			for y := curY - 1; y >= (curY - size); y-- {
				grid[y][curX] += val
			}
			curY -= size
		}
	}
}

func makeGrid(maxX, maxY int) {
	grid = make([][]int, maxY*3)
	for i := range grid {
		grid[i] = make([]int, maxX*3)
	}

	logger.Info(
		"overall grid size",
		zap.Int("x", len(grid[0])),
		zap.Int("y", len(grid)),
	)
}

func maxOverallGrid() (int, int) {
	maxX, maxY := maxGridForWire(wire1)
	tempX, tempY := maxGridForWire(wire2)

	maxX = int(math.Max(float64(maxX), float64(tempX)))
	maxY = int(math.Max(float64(maxY), float64(tempY)))

	return maxX, maxY
}

func maxGridForWire(wire []string) (int, int) {
	maxX := 0
	maxY := 0
	xPos := 0
	yPos := 0

	for _, movement := range wire {

		direction, size := dirAndSize(movement)

		switch direction {
		case "R":
			xPos += size
			maxX = int(math.Max(math.Abs(float64(xPos)), float64(maxX)))
		case "L":
			xPos -= size
			maxX = int(math.Max(math.Abs(float64(xPos)), float64(maxX)))
		case "D":
			yPos -= size
			maxY = int(math.Max(math.Abs(float64(yPos)), float64(maxY)))
		case "U":
			yPos += size
			maxY = int(math.Max(math.Abs(float64(yPos)), float64(maxY)))
		}
	}
	return maxX, maxY
}

func dirAndSize(movement string) (string, int) {
	direction := string(movement[0])
	size, _ := strconv.Atoi(movement[1:len(movement)])

	return direction, size
}

var distance = math.Inf(1)
var grid [][]int
var logger *zap.Logger
var wire1 = []string{
	"R8",
	"U5",
	"L5",
	"D3",
}

var wire2 = []string{
	"U7",
	"R6",
	"D4",
	"L4",
}
