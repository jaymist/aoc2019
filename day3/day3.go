package main

import (
	"fmt"
	"math"
	"strconv"

	"go.uber.org/zap"
)

func main() {
	logger = zap.NewExample()
	maxX, maxY := maxOverallGrid()
	logger.Info("max coordinates", zap.Int("max-x", maxX), zap.Int("max-y", maxY))
	makeGrid(maxX, maxY)

	centerY := len(grid) - 1
	centerX := 0

	unravelWires(centerX, centerY)
	scanGrid(centerX, centerY)

	logger.Info("shortest distance", zap.Float64("distance", distance))
	for y := range grid {
		fmt.Println(y, ":\t", grid[y])
	}
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
				logger.Info(
					"down movement",
					zap.Int("curr-x", curX),
					zap.Int("y", y),
					zap.Int("current-val", grid[y][curX]),
				)
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

	logger.Info(
		"ending coordinates",
		zap.Int("curr-x", curX),
		zap.Int("curr-y", curY),
	)

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
