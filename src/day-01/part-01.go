package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	data, _ := os.ReadFile(dir + "/src/day-01/input.txt")
	dataArray := strings.Split(string(data), "\n")
	number := 0
	highest := 0

	for i := 0; i < len(dataArray); i++ {
		if dataArray[i] == "" || i == len(dataArray)-1 {
			if number > highest {
				highest = number
			}
			number = 0
		}
		numberToAdd, _ := strconv.ParseInt(dataArray[i], 10, 64)
		number += int(numberToAdd)
	}

	fmt.Println("HIGHEST: ", highest)
}
