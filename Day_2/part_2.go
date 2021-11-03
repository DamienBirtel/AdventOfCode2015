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

		biggestIndex := 0
		wrap := sides[0]
		for i := 1; i < 3; i++ {
			wrap *= sides[i]
			if sides[i] > sides[biggestIndex] {
				biggestIndex = i
			}
		}
		for i, nb := range sides {
			if i != biggestIndex {
				wrap += 2 * nb
			}
		}
		totalWrap += wrap
	}
	fmt.Println(totalWrap)
}
