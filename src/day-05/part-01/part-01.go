package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/day-05/input.txt")
	scanner := bufio.NewScanner(file)

	cargoColumns := make([]string, 0)
	cargoRows := make([]string, 0)
	initialising := true

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		if string(text[1]) == "1" {
			cargoColumns = fillCargoColumns(len(cleanCargoRow(text)), cargoRows)

			// Can run commands now
			initialising = false
			continue
		}

		if initialising {
			cargoRows = append(cargoRows, cleanCargoRow(text))
			continue
		}

		runCommand(text, cargoColumns)
	}

	text := ""
	for i := 0; i < len(cargoColumns); i++ {
		text += string(cargoColumns[i][len(cargoColumns[i])-1])
	}

	fmt.Println(text)
}

func cleanCargoRow(row string) string {
	text := ""
	for i := 1; i < len(row); i += 4 {
		text += string(row[i])
	}
	return text
}

func fillCargoColumns(columnCount int, cargoRows []string) []string {
	cargoColumns := make([]string, columnCount)

	for i := len(cargoRows) - 1; i >= 0; i-- {
		for j := 0; j < len(cargoRows[i]); j++ {
			character := string(cargoRows[i][j])
			if character == " " {
				continue
			}
			cargoColumns[j] += character
		}
	}

	return cargoColumns
}

func runCommand(command string, cargoColumns []string) {
	var count, from, to int
	fmt.Sscanf(command, "move %d from %d to %d", &count, &from, &to)

	fromIndex := from - 1
	toIndex := to - 1

	for i := 0; i < count; i++ {
		lastCharacterPosition := len(cargoColumns[fromIndex]) - 1
		character := string(cargoColumns[fromIndex][lastCharacterPosition])

		cargoColumns[fromIndex] = cargoColumns[fromIndex][:lastCharacterPosition]
		cargoColumns[toIndex] += character
	}
}
