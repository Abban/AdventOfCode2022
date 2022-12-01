package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	data, _ := os.ReadFile(dir + "/src/day-01/input.txt")
	dataArray := strings.Split(string(data), "\n")
	number := 0
	combined := make([]int, 0)

	for i := 0; i < len(dataArray); i++ {
		if dataArray[i] == "" || i == len(dataArray)-1 {
			combined = append(combined, number)
			number = 0
		}
		numberToAdd, _ := strconv.ParseInt(dataArray[i], 10, 64)
		number += int(numberToAdd)
	}

	sort.Slice(combined, func(p, q int) bool {
		return combined[p] > combined[q]
	})

	fmt.Println("HIGHEST: ", combined[0]+combined[1]+combined[2])
}
