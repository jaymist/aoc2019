package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ = zap.NewDevelopment()
	var i int
	var j int
	done := false

	for i = 0; i < 100; i++ {
		for j = 0; j < 100; j++ {
			resetInput()

			input[1] = i
			input[2] = j
			calculateResult()
			logger.Info("running test", zap.Int("result", input[0]), zap.Int("noun", i), zap.Int("verb", j))
			if input[0] == 19690720 {
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	logger.Info("Finsihed", zap.Int("noun", i), zap.Int("verb", j))
	logger.Sugar().Infof("Result: %d", 100*i+j)
}

func calculateResult() {
	finished := false
	index := 0

	for !finished {
		opcode := input[index]

		if opcode == 99 {
			finished = true
			break
		}

		pos1 := input[index+1]
		pos2 := input[index+2]
		posResult := input[index+3]

		val1 := input[pos1]
		val2 := input[pos2]

		if posResult >= len(input) {
			logger.Info("Intcode", zap.Int("pos1", pos1), zap.Int("pos2", pos2), zap.Int("posResult", posResult), zap.Int("val1", val1), zap.Int("val2", val2))
		}

		if opcode == 1 {
			input[posResult] = val1 + val2
		} else if opcode == 2 {
			input[posResult] = val1 * val2
		}
		index += 4
	}
}

func resetInput() {
	input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 9, 23, 27, 1, 5, 27, 31, 1, 5, 31, 35, 1, 35, 13, 39, 1, 39, 9, 43, 1, 5, 43, 47, 1, 47, 6, 51, 1, 51, 13, 55, 1, 55, 9, 59, 1, 59, 13, 63, 2, 63, 13, 67, 1, 67, 10, 71, 1, 71, 6, 75, 2, 10, 75, 79, 2, 10, 79, 83, 1, 5, 83, 87, 2, 6, 87, 91, 1, 91, 6, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 9, 107, 1, 10, 107, 111, 2, 111, 13, 115, 1, 10, 115, 119, 1, 10, 119, 123, 2, 13, 123, 127, 2, 6, 127, 131, 1, 13, 131, 135, 1, 135, 2, 139, 1, 139, 6, 0, 99, 2, 0, 14, 0}
}

var logger *zap.Logger
var input = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 9, 23, 27, 1, 5, 27, 31, 1, 5, 31, 35, 1, 35, 13, 39, 1, 39, 9, 43, 1, 5, 43, 47, 1, 47, 6, 51, 1, 51, 13, 55, 1, 55, 9, 59, 1, 59, 13, 63, 2, 63, 13, 67, 1, 67, 10, 71, 1, 71, 6, 75, 2, 10, 75, 79, 2, 10, 79, 83, 1, 5, 83, 87, 2, 6, 87, 91, 1, 91, 6, 95, 1, 95, 13, 99, 2, 99, 13, 103, 1, 103, 9, 107, 1, 10, 107, 111, 2, 111, 13, 115, 1, 10, 115, 119, 1, 10, 119, 123, 2, 13, 123, 127, 2, 6, 127, 131, 1, 13, 131, 135, 1, 135, 2, 139, 1, 139, 6, 0, 99, 2, 0, 14, 0}
