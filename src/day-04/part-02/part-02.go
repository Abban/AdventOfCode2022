package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	data, _ := os.ReadFile(dir + "/src/day-04/input.txt")
	dataArray := strings.Split(string(data), "\n")

	total := 0

	for _, row := range dataArray {
		if row == "" {
			continue
		}
		sides := strings.Split(row, ",")
		left := splitToIntValues(sides[0])
		right := splitToIntValues(sides[1])
		if rangeOverlapsRange(left, right) {
			total++
		}
	}

	fmt.Println("TOTAL: ", total)
}

func rangeOverlapsRange(left []int, right []int) bool {
	if left[1] < right[0] {
		return false
	}
	if left[0] > right[1] {
		return false
	}
	return true
}

func splitToIntValues(item string) []int {
	data := strings.Split(item, "-")
	left, _ := strconv.Atoi(data[0])
	right, _ := strconv.Atoi(data[1])
	return []int{left, right}
}
