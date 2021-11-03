package main

import (
	"fmt"
	"os"
)

type Coordinates struct {
	x int
	y int
}

func main() {
	rawInput, err := os.ReadFile("./input.txt")
	if err != nil {
		return
	}

	directions := string(rawInput)
	currentCoordinates := Coordinates{0, 0}
	coordinates := map[Coordinates]int{currentCoordinates: 1}
	for _, char := range directions {
		switch char {
		case '^':
			currentCoordinates.x += 1
		case 'v':
			currentCoordinates.x -= 1
		case '<':
			currentCoordinates.y -= 1
		case '>':
			currentCoordinates.y += 1
		}
		coordinates[currentCoordinates] += 1
	}
	fmt.Println(len(coordinates))
}
