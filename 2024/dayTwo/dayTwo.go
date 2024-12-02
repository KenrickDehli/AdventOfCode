package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting program...")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputAsStringsNoSpaces := strings.Split(string(file), "\n")

	var amountSafeReports int = 0

	for i := range inputAsStringsNoSpaces {
		reportNoSpaces := strings.Split(inputAsStringsNoSpaces[i], " ")
		if len(reportNoSpaces) <= 1 {
			break
		}
		var increasing []bool
		var decreasing []bool
		var differenceIsSafe bool = false
		var report []int

		for j, v := range reportNoSpaces {
			if currentNumber, err := strconv.Atoi(v); err == nil {
				report = append(report, currentNumber)

				if j > 0 {
					increasing = append(increasing, report[j] > report[j-1])
					decreasing = append(decreasing, report[j] < report[j-1])
					difference := report[j] - report[j-1]
					if difference < 0 {
						difference *= -1
					}
					differenceIsSafe = difference >= 1 && difference <= 3
					if !differenceIsSafe {
						break
					}
				}
			}

		}
		sameIncreasingPattern := !slices.Contains(increasing, !increasing[0])
		sameDecreasingPattern := !slices.Contains(decreasing, !decreasing[0])
		if sameIncreasingPattern && sameDecreasingPattern && differenceIsSafe {
			amountSafeReports += 1
		}
	}
	fmt.Println("Finished program. The amount of safe reports:", amountSafeReports)
}
