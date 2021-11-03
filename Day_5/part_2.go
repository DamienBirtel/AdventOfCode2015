package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("./input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	nice := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {

		}
	}
	fmt.Println(nice)
}
