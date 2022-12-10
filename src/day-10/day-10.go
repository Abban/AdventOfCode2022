package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/day-10/input.txt")
	scanner := bufio.NewScanner(file)

	commands := make([]string, 0)
	x := 1

	for scanner.Scan() {
		text := scanner.Text()
		if text != "noop" {
			commands = append(commands, "buffer")
		}

		commands = append(commands, text)
	}

	signalStrength := 0
	pixels := makePixels(240)

	for i := 0; i < len(commands); i++ {
		cycle := i + 1

		if isImportantCycle(cycle) {
			signalStrength += cycle * x
		}

		position := int(math.Mod(float64(i), 40))
		if x == position || x == position-1 || x == position+1 {
			pixels[i] = "#"
		}

		if commands[i] != "noop" && commands[i] != "buffer" {
			x += parseCommand(commands[i])
		}
	}

	fmt.Println("Part 1:", signalStrength)
	fmt.Println("Part 2:")

	for i := 0; i < len(pixels); i++ {
		fmt.Print(pixels[i])
		if isEndOfRow(i + 1) {
			fmt.Println("")
		}
	}
}

func parseCommand(line string) int {
	cmd, _ := strconv.Atoi(strings.Split(line, " ")[1])
	return cmd
}

func isImportantCycle(cycle int) bool {
	return cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220
}

func isEndOfRow(cycle int) bool {
	return cycle == 40 || cycle == 80 || cycle == 120 || cycle == 160 || cycle == 200 || cycle == 240
}

func makePixels(length int) []string {
	pixels := make([]string, length)
	for i := 0; i < length; i++ {
		pixels[i] = "."
	}
	return pixels
}
