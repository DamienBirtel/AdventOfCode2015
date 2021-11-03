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
	for _, char := range text {
		if char == '(' {
			answer++
		}
		if char == ')' {
			answer--
		}
	}
	fmt.Println(answer)
}
