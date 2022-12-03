package main

import (
	"fmt"
	"os"
	"strings"
)

var values = map[string]int{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13,
	"n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26,
	"A": 27, "B": 28, "C": 29, "D": 30, "E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36, "K": 37, "L": 38, "M": 39,
	"N": 40, "O": 41, "P": 42, "Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48, "W": 49, "X": 50, "Y": 51, "Z": 52,
}

func main() {
	dir, _ := os.Getwd()
	data, _ := os.ReadFile(dir + "/src/day-03/input.txt")
	dataArray := strings.Split(string(data), "\n")

	total := 0

	for i := 0; i < len(dataArray); i += 3 {
		if dataArray[i] == "" {
			continue
		}

		first := strings.Split(dataArray[i], "")
		second := strings.Split(dataArray[i+1], "")
		third := strings.Split(dataArray[i+2], "")

		intersection := findFirstIntersection(arrayIntersection(first, second), third)
		total += values[intersection]
	}

	fmt.Println("TOTAL: ", total)
}

func arrayIntersection(left []string, right []string) []string {
	intersections := make([]string, 0)
	for _, character := range left {
		if arrayContainsElement(right, character) && !arrayContainsElement(intersections, character) {
			intersections = append(intersections, character)
		}
	}

	return intersections
}

func findFirstIntersection(left []string, right []string) string {
	for _, character := range left {
		if arrayContainsElement(right, character) {
			return character
		}
	}

	panic("Could not find intersection")
}

func arrayContainsElement(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
