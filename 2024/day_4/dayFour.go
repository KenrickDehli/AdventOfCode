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
	//	validForward := "XMAS"
	//	validBackward := "SAMX"
	countOfRows := len(input) - 2

	for i := range input {
		for j := range input[i] {
			rightString := string(input[i][j:calcBounds(j+4, len(input[i])-1, false)])
			leftString := string(input[i][calcBounds(j-4, 0, true):j])
			upperString, lowerString := buildString(input, true, i, j, countOfRows)
			fmt.Println("upper", upperString)
			fmt.Println("right", rightString)
			fmt.Println("left", leftString)
			fmt.Println("lower", lowerString)
		}
	}

	fmt.Println("Finished program. The sum of uncorrupted enabled multiplications:")
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

func buildString(str []string, upperLower bool, rowIndex int, columnIndex int, countOfRows int) (string, string) {
	var ret1, ret2 string
	if upperLower {
		ret1 = string(str[calcBounds(rowIndex-1, 0, true)][columnIndex]) + string(str[calcBounds(rowIndex-2, 0, true)][columnIndex]) + string(str[calcBounds(rowIndex-3, 0, true)][columnIndex]) + string(str[calcBounds(rowIndex-4, 0, true)][columnIndex])
		ret2 = string(str[calcBounds(rowIndex+1, countOfRows, false)][columnIndex]) + string(str[calcBounds(rowIndex+2, countOfRows, false)][columnIndex]) + string(str[calcBounds(rowIndex+3, countOfRows, false)][columnIndex]) + string(str[calcBounds(rowIndex+4, countOfRows, false)][columnIndex])
		return ret1, ret2
	}
	//ret1 = string(str[calcBounds(rowIndex)])
	return ret1, ret2

}
