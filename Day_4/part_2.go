package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		return
	}

	i := 0
	for {
		nb := strconv.Itoa(i)
		hasher := md5.New()
		toHash := append(data, []byte(nb)...)
		hasher.Write(toHash)
		hashedStr := hex.EncodeToString(hasher.Sum(nil))
		if hashedStr[0:6] == "000000" {
			fmt.Println(i)
			break
		}
		i++
	}
}
