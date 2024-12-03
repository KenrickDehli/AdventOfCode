package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting program...")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.ReplaceAll(string(file), "\n", "")

	numberRegex := regexp.MustCompile(`mul\((?<left>[0-9]*),(?<right>[0-9]*)\)`)
	invalidRange := regexp.MustCompile(`don't\(\).*?do\(\)`)

	enabledRange := invalidRange.ReplaceAllString(input, "")
	enabledMulti := numberRegex.FindAllStringSubmatch(enabledRange, -1)
	allValidMulti := numberRegex.FindAllStringSubmatch(input, -1)

	//partOne
	sumUncorruptedMultiplications := calculateMultiplication(allValidMulti)

	//partTwo
	sumEnabledMultiplications := calculateMultiplication(enabledMulti)

	fmt.Println("Finished program. The sum of uncorrupted multiplications:", sumUncorruptedMultiplications)
	fmt.Println("Finished program. The sum of uncorrupted enabled multiplications:", sumEnabledMultiplications)
}

func calculateMultiplication(slice [][]string) int {
	var amountSafeReports int = 0
	for i := range slice {
		for j := 1; j < len(slice[i])-1; j++ {
			left, err := strconv.Atoi(slice[i][j])
			if err != nil {
				log.Fatal(err)
			}
			right, err := strconv.Atoi(slice[i][j+1])
			if err != nil {
				log.Fatal(err)
			}
			amountSafeReports += left * right
		}
	}
	return amountSafeReports
}
