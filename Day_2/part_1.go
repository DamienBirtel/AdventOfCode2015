package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	totalWrap := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "x")
		sides := []int{}
		for _, word := range text {
			nb, _ := strconv.Atoi(word)
			sides = append(sides, nb)
		}

		smallest := -1
		wrap := 0
		for i := 0; i < 3; i++ {
			for j := i + 1; j < len(sides); j++ {
				dimension := sides[i] * sides[j]
				if smallest < 0 || dimension < smallest {
					smallest = dimension
				}
				wrap += 2 * dimension
			}
		}
		totalWrap += smallest + wrap
	}
	fmt.Println(totalWrap)
}
