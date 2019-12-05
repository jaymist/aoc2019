package main

import (
	"go.uber.org/zap"
	"sort"
	"strconv"
	"strings"
)

var logger, _ = zap.NewDevelopment()

func main() {
	start := 172851
	end := 675869
	allPasswords := make(map[int]bool)

	for i := start; i <= end; i++ {
		digits := convertToDigits(i)

		ispass, password := isPassword(digits)
		if ispass && password >= start && password <= end {
			allPasswords[password] = true
		}
	}

	logger.Info("passwords", zap.Int("length", len(allPasswords)))
}

func isPassword(digits []int) (bool, int) {
	sort.Ints(digits)
	possiblePassword := false
	password := false
	digitStr := make([]string, len(digits))
	digitCount := make(map[int]int)

	for index, digit := range digits {
		text := strconv.Itoa(digit)
		digitStr[index] = text
		digitCount[digit] += 1

		if index > 0 && digit == digits[index-1] {
			possiblePassword = true
		}
	}

	if possiblePassword {
		for _, count := range digitCount {
			if count == 2 {
				password = true
			}
		}
	}
	result, _ := strconv.Atoi(strings.Join(digitStr, ""))
	return password, result
}

func convertToDigits(val int) []int {
	var digits = make([]int, 6)

	for i := (len(digits) - 1); i >= 0; i-- {
		digit := 0
		digit, val = extractDigit(val)
		digits[i] = digit
	}
	return digits
}

func extractDigit(val int) (int, int) {
	remainder := val / 10
	digit := val % 10

	return digit, remainder
}
