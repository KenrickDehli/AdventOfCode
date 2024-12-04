package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting program...")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(file), "\n")
	validForward := "XMAS"
	validBackward := "SAMX"
	countOfRows := len(input) - 2
	countOfXMAS := 0

	for i := range input {
		for j := range input[i] {
			substrings := buildString(input, i, j, countOfRows, len(input[i]))
			for k := range substrings {
				if strings.Contains(substrings[k], validForward) || strings.Contains(substrings[k], validBackward) {
					fmt.Println(i, substrings[k])
					countOfXMAS++
				}
			}
		}
	}

	fmt.Println("Finished program. The sum of XMAS:", countOfXMAS)
}

func calcBounds(value int, bound int, checknegative bool) int {
	if checknegative {
		if value < bound {
			return bound
		}
		return value
	}
	if value > bound {
		return bound
	}
	return value

}

func buildString(str []string, rowIndex int, columnIndex int, countOfRows int, countOfColumns int) []string {
	var right, left, up, down, upright, upleft, lowleft, lowright string
	right = string(str[rowIndex][columnIndex:calcBounds(columnIndex+4, countOfColumns-1, false)])
	left = string(str[rowIndex][calcBounds(columnIndex-4, 0, true):columnIndex])

	up = string(str[calcBounds(rowIndex, 0, true)][columnIndex]) + string(str[calcBounds(rowIndex-1, 0, true)][columnIndex]) + string(str[calcBounds(rowIndex-2, 0, true)][columnIndex]) + string(str[calcBounds(rowIndex-3, 0, true)][columnIndex])
	down = string(str[calcBounds(rowIndex, countOfRows, false)][columnIndex]) + string(str[calcBounds(rowIndex+1, countOfRows, false)][columnIndex]) + string(str[calcBounds(rowIndex+2, countOfRows, false)][columnIndex]) + string(str[calcBounds(rowIndex+3, countOfRows, false)][columnIndex])

	upright = string(str[calcBounds(rowIndex, 0, true)][calcBounds(columnIndex, countOfColumns-1, false)]) + string(str[calcBounds(rowIndex-1, 0, true)][calcBounds(columnIndex+1, countOfColumns-1, false)]) + string(str[calcBounds(rowIndex-2, 0, true)][calcBounds(columnIndex+2, countOfColumns-1, false)]) + string(str[calcBounds(rowIndex-3, 0, true)][calcBounds(columnIndex+3, countOfColumns-1, false)])
	upleft = string(str[calcBounds(rowIndex, 0, true)][calcBounds(columnIndex, 0, true)]) + string(str[calcBounds(rowIndex-1, 0, true)][calcBounds(columnIndex-1, 0, true)]) + string(str[calcBounds(rowIndex-2, 0, true)][calcBounds(columnIndex-2, 0, true)]) + string(str[calcBounds(rowIndex-3, 0, true)][calcBounds(columnIndex-3, 0, true)])

	lowleft = string(str[calcBounds(rowIndex, countOfRows, false)][calcBounds(columnIndex, 0, true)]) + string(str[calcBounds(rowIndex+1, countOfRows, false)][calcBounds(columnIndex-1, 0, true)]) + string(str[calcBounds(rowIndex+2, countOfRows, false)][calcBounds(columnIndex-2, 0, true)]) + string(str[calcBounds(rowIndex+3, countOfRows, false)][calcBounds(columnIndex-3, 0, true)])
	lowright = string(str[calcBounds(rowIndex, countOfRows, false)][calcBounds(columnIndex, countOfColumns-1, false)]) + string(str[calcBounds(rowIndex+1, countOfRows, false)][calcBounds(columnIndex+1, countOfColumns-1, false)]) + string(str[calcBounds(rowIndex+2, countOfRows, false)][calcBounds(columnIndex+2, countOfColumns-1, false)]) + string(str[calcBounds(rowIndex+3, countOfRows, false)][calcBounds(columnIndex+3, countOfColumns-1, false)])

	return []string{right, left, up, down, upright, upleft, lowleft, lowright}
}
