package main

import (
	"bufio"
	"fmt"
	"os"
)

var badChars = map[[2]rune]int{
	{'a', 'b'}: 0,
	{'c', 'd'}: 0,
	{'p', 'q'}: 0,
	{'x', 'y'}: 0,
}

var vowelChars = map[rune]int{
	'a': 0,
	'e': 0,
	'i': 0,
	'o': 0,
	'u': 0,
}

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
		chars := [2]rune{'0', '0'}
		double := false
		naughty := false
		vowels := 0
		for i, char := range line {

			_, ok := vowelChars[char]
			if ok {
				vowels++
			}

			if i == 0 {
				chars[1] = char
				continue
			}

			chars[0] = chars[1]
			chars[1] = char
			_, ok = badChars[chars]
			if ok {
				naughty = true
				break
			}

			if chars[0] == chars[1] {
				double = true
			}
		}
		if !naughty && double && vowels > 2 {
			nice++
		}
	}
	fmt.Println(nice)
}
