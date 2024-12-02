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
		var report []int

		for k := range reportNoSpaces {

			if currentNumber, err := strconv.Atoi(reportNoSpaces[k]); err == nil {
				report = append(report, currentNumber)
			}

		}
		safe := safe(slices.Clone(report), 0)
		if safe {
			amountSafeReports += 1
		}
		if !safe {
		}
	}
	fmt.Println("Finished program. The amount of safe reports:", amountSafeReports)
}

func safe(report []int, level int) bool {
	if level > 1 {
		return false
	}
	var increasing []bool
	var decreasing []bool
	var differences []bool
	var differenceIsSafe bool = false
	for j := range report {

		if j > 0 {
			increasing = append(increasing, report[j] > report[j-1])
			decreasing = append(decreasing, report[j] < report[j-1])
			difference := report[j] - report[j-1]
			if difference < 0 {
				difference *= -1
			}
			differenceIsSafe = difference >= 1 && difference <= 3
			differences = append(differences, differenceIsSafe)
		}

	}

	safeIncreasingPattern := !slices.Contains(increasing, !increasing[0])
	safeDecreasingPattern := !slices.Contains(decreasing, !decreasing[0])
	safeDifferencePattern := !slices.Contains(differences, !differences[0])

	if safeIncreasingPattern && safeDecreasingPattern && safeDifferencePattern {
		return true
	}

	var indexChange int

	if !safeIncreasingPattern {
		indexChange = slices.Index(increasing, !increasing[0])

		if indexChange == len(increasing)-1 && countOf(increasing, increasing[indexChange]) == 1 {
			indexChange += 1
		}

		if indexChange == 1 && countOf(increasing, increasing[0]) == 1 {
			indexChange -= 1
		}

	} else if !safeDecreasingPattern {
		indexChange = slices.Index(decreasing, !decreasing[0])

		if indexChange == len(decreasing)-1 && countOf(decreasing, decreasing[indexChange]) == 1 {
			indexChange += 1
		}

		if indexChange == 1 && countOf(decreasing, decreasing[0]) == 1 {
			indexChange -= 1
		}
	} else if !safeDifferencePattern {
		indexChange = slices.Index(differences, !differences[0])

		if indexChange == len(differences)-1 && countOf(differences, differences[indexChange]) == 1 {
			indexChange += 1
		}

		if indexChange == 1 && countOf(differences, differences[0]) == 1 {
			indexChange -= 1
		}
	}

	temp := slices.Delete(slices.Clone(report), indexChange, indexChange+1)
	tempLevel := level + 1
	return safe(temp, tempLevel)
}

func countOf(slice []bool, value bool) int {
	var count int = 0
	for i := range slice {
		if slice[i] == value {
			count++
		}
	}
	return count

}

