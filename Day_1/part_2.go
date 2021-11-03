package main

import (
	"fmt"
	"os"
)

func main() {
	rawInput, err := os.ReadFile("./input.txt")
	if err != nil {
		return
	}
	answer := 0
	text := string(rawInput)
	for i, char := range text {
		if char == '(' {
			answer++
		}
		if char == ')' {
			answer--
		}
		if answer == -1 {
			fmt.Println(i + 1)
			return
		}
	}
}
