package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var up = "^"
var down = "v"
var left = "<"
var right = ">"

var positions [][]int

func main() {
	fmt.Println("Starting program...")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	t := strings.Split(string(file), "\n")
	var input []string

	for _, k := range t {
		if strings.Contains(k, ".") {
			input = append(input, k)
		}
	}

	fmt.Println(input)
	sum := 0

	positions = make([][]int, len(input))

	for i := range positions {
		var temp []int
		for range input[0] {
			temp = append(temp, 0)
		}
		positions[i] = append(temp, 0)
	}

	x, y, direction := getPosition(input)
	end := false
	for !end {
		switch direction {
		case "up":
			x, y, direction = goUp(input, x, y)
		case "down":
			x, y, direction = goDown(input, x, y)
		case "left":
			x, y, direction = goLeft(input, x, y)
		case "right":
			x, y, direction = goRight(input, x, y)
		case "leave":
			end = true
		}

	}

	fmt.Println(x, y, direction)

	for _, v := range positions {
		fmt.Println(v)

		for _, k := range v {
			if k == 1 {
				sum++
			}
		}
	}

	fmt.Println("Finished program. The sum of middle page number of the correctly-ordered updates:", sum)
}

func getPosition(guardMap []string) (int, int, string) {
	var x int
	var y int
	var direction string

	for i, v := range guardMap {
		if strings.Contains(v, up) {
			x = i
			y = strings.Index(guardMap[i], up)
			direction = "up"
		}
		if strings.Contains(v, down) {
			x = i
			y = strings.Index(guardMap[i], down)
			direction = "down"
		}
		if strings.Contains(v, left) {
			x = i
			y = strings.Index(guardMap[i], left)
			direction = "left"
		}
		if strings.Contains(v, right) {
			x = i
			y = strings.Index(guardMap[i], right)
			direction = "right"
		}
	}
	return x, y, direction
}

func goUp(guardMap []string, startX int, startY int) (int, int, string) {
	var x int
	var y int
	var direction string
	fmt.Println("going up")

	for i := startX; i > 0; i-- {
		isHashtag := string(guardMap[i-1][startY]) == "#"
		if isHashtag {
			direction = "right"
			x = i
			y = startY
		}
		if i-1 == 0 && !isHashtag {
			direction = "leave"
		}
		positions[i][startY] = 1
	}
	return x, y, direction
}

func goDown(guardMap []string, startX int, startY int) (int, int, string) {
	var x int
	var y int
	var direction string
	fmt.Println("going down")

	for i := startX; i < len(guardMap)-1; i++ {
		fmt.Println(i)
		isHashtag := string(guardMap[i+1][startY]) == "#"
		if isHashtag {
			direction = "left"
			x = i
			y = startY
		}
		if i+1 == len(guardMap)-1 && !isHashtag {
			direction = "leave"
			x = i
			y = startY
			positions[i+1][startY] = 1
		}
		positions[i][startY] = 1
	}
	return x, y, direction
}

func goLeft(guardMap []string, startX int, startY int) (int, int, string) {
	var x int
	var y int
	var direction string
	fmt.Println("going left")

	for i := startY; i > 0; i-- {
		fmt.Println(i)
		isHashtag := string(guardMap[startX][i-1]) == "#"
		if isHashtag {
			direction = "up"
			x = startY
			y = i
		}
		if i-1 == 0 && !isHashtag {
			direction = "leave"
			x = startX
			y = i
			positions[startX][i-1] = 1
		}
		positions[startX][i-1] = 1
	}
	return x, y, direction
}

func goRight(guardMap []string, startX int, startY int) (int, int, string) {
	var x int
	var y int
	var direction string
	fmt.Println("going right")

	for i := startY; i < len(guardMap[startX])-1; i++ {
		fmt.Println(i)
		isHashtag := string(guardMap[startX][i+1]) == "#"
		if isHashtag {
			direction = "down"
			x = i
			y = startY
		}
		if i+1 == len(guardMap[startX])-1 && !isHashtag {
			direction = "leave"
			x = i
			y = startY
			positions[startX][i+1] = 1
		}
		positions[startX][i] = 1
	}
	return x, y, direction
}
