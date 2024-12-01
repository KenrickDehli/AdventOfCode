package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	noSpaceBetweenNumbers := strings.ReplaceAll(string(file), "   ", " ")
	noLineBreaksAndNoSpaceBetweenNumbers := strings.ReplaceAll(noSpaceBetweenNumbers, "\n", " ")

	input := strings.Split(noLineBreaksAndNoSpaceBetweenNumbers, " ")

	var fistArr, secondArr []int
	var secondArrAsString []string

	for i := 0; i < len(input); i++ {
		if currentNumber, err := strconv.ParseInt(input[i], 10, 64); err == nil {
			if i%2 == 0 {
				fistArr = append(fistArr, int(currentNumber))
			} else {
				secondArr = append(secondArr, int(currentNumber))
				secondArrAsString = append(secondArrAsString, input[i])
			}
		}
	}

	sort.Sort(sort.IntSlice(fistArr))
	sort.Sort(sort.IntSlice(secondArr))
	var result int = 0
	var similarityScores int = 0
	secondArrAsOneString := strings.Join(secondArrAsString, "\n")

	for index, value := range fistArr {

		//part 1
		res := fistArr[index] - secondArr[index]
		if res < 0 {
			res = res * -1
		}
		result += res

		//part2
		count := strings.Count(secondArrAsOneString, strconv.Itoa(value))
		similarityScore := value * count
		similarityScores += similarityScore

	}

	fmt.Println("Finished with calculating the differences between the locationIds. Result:", result)
	fmt.Println("Finished with calculating the similarity score. Result:", similarityScores)

}
