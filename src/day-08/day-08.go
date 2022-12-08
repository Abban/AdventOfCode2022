package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector2 struct {
	x int
	y int
}

type Directions struct {
	up    Vector2
	down  Vector2
	left  Vector2
	right Vector2
}

func (v *Vector2) Add(toAdd Vector2) {
	v.x += toAdd.x
	v.y += toAdd.y
}

func main() {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/day-08/input.txt")
	scanner := bufio.NewScanner(file)

	trees := make([][]int, 0)

	for scanner.Scan() {
		trees = append(trees, getMatrixRow(scanner.Text()))
	}

	size := Vector2{x: len(trees[0]), y: len(trees)}
	directions := Directions{
		up:    Vector2{x: 0, y: -1},
		down:  Vector2{x: 0, y: 1},
		left:  Vector2{x: -1, y: 0},
		right: Vector2{x: 1, y: 0},
	}

	part1 := 0
	part2 := 0
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			treePosition := Vector2{x: x, y: y}
			if canSeeToEdge(treePosition, trees, size, directions) {
				part1++
			}

			scenicScore := getScenicScore(treePosition, trees, size, directions)
			if scenicScore > part2 {
				part2 = scenicScore
			}
		}
	}
	fmt.Println("Part 01: ", part1)
	fmt.Println("Part 02: ", part2)
}

func canSeeToEdge(position Vector2, trees [][]int, size Vector2, directions Directions) bool {
	return isEdge(position, size) ||
		visibleInDirection(position, directions.up, size, trees) ||
		visibleInDirection(position, directions.down, size, trees) ||
		visibleInDirection(position, directions.left, size, trees) ||
		visibleInDirection(position, directions.right, size, trees)
}

func getScenicScore(position Vector2, trees [][]int, size Vector2, directions Directions) int {
	if isEdge(position, size) {
		return 0
	}

	return countVisibleTreesInDirection(position, directions.up, size, trees) *
		countVisibleTreesInDirection(position, directions.down, size, trees) *
		countVisibleTreesInDirection(position, directions.left, size, trees) *
		countVisibleTreesInDirection(position, directions.right, size, trees)
}

func isEdge(treePosition Vector2, size Vector2) bool {
	if treePosition.x == 0 || treePosition.y == 0 {
		return true
	}
	if treePosition.x == size.x-1 || treePosition.y == size.y-1 {
		return true
	}
	return false
}

func visibleInDirection(treePosition Vector2, direction Vector2, size Vector2, trees [][]int) bool {
	height := trees[treePosition.y][treePosition.x]
	currentPosition := treePosition

	for {
		currentPosition.Add(direction)

		if goneOffGrid(currentPosition, size) {
			return true
		}

		if height <= trees[currentPosition.y][currentPosition.x] {
			return false
		}
	}
}

func countVisibleTreesInDirection(treePosition Vector2, direction Vector2, size Vector2, trees [][]int) int {
	height := trees[treePosition.y][treePosition.x]
	currentPosition := treePosition
	visible := 0

	for {
		currentPosition.Add(direction)

		if goneOffGrid(currentPosition, size) {
			break
		}

		if height <= trees[currentPosition.y][currentPosition.x] {
			visible++
			break
		}

		visible++
	}

	return visible
}

func goneOffGrid(position Vector2, size Vector2) bool {
	if position.x < 0 || position.y < 0 {
		return true
	}
	if position.x == size.x || position.y == size.y {
		return true
	}
	return false
}

func getMatrixRow(text string) []int {
	row := make([]int, 0)

	for _, character := range strings.Split(text, "") {
		converted, _ := strconv.Atoi(character)
		row = append(row, converted)
	}

	return row
}
