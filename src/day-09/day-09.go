package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector2 struct {
	x int
	y int
}

func (v *Vector2) Move(toAdd Vector2) {
	v.x += toAdd.x
	v.y += toAdd.y
}

func (v *Vector2) Delta(other Vector2) Vector2 {
	return Vector2{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v *Vector2) ToString() string {
	return fmt.Sprintf("%d,%d", v.x, v.y)
}

func main() {
	dir, _ := os.Getwd()
	fmt.Println("Sample 01:", getChainEndVisitedCount(2, dir+"/src/day-09/sample.txt"))
	fmt.Println("Input 01:", getChainEndVisitedCount(2, dir+"/src/day-09/input.txt"))
	fmt.Println("Sample 02:", getChainEndVisitedCount(10, dir+"/src/day-09/sample.txt"))
	fmt.Println("Input 02:", getChainEndVisitedCount(10, dir+"/src/day-09/input.txt"))
}

func getChainEndVisitedCount(length int, path string) int {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	visited := map[string]int{"0,0": 1}
	chain := makeChain(length)

	for scanner.Scan() {
		scanner.Text()
		var direction string
		var steps int

		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &steps)

		for i := 0; i < steps; i++ {
			chain[0].Move(getDirection(direction))
			for j := 1; j < len(chain); j++ {
				delta := chain[j-1].Delta(chain[j])

				if getIntAbs(delta.x) > 1 || getIntAbs(delta.y) > 1 {
					chain[j].Move(clampVector(delta, Vector2{x: 1, y: 1}, Vector2{x: -1, y: -1}))
					if j == len(chain)-1 {
						visited[chain[j].ToString()] = 1
					}
				}
			}
		}
	}

	return len(visited)
}

func makeChain(count int) []Vector2 {
	chain := make([]Vector2, count)
	for i := 0; i < count; i++ {
		chain[i] = Vector2{x: 0, y: 0}
	}
	return chain
}

func getDirection(direction string) Vector2 {
	switch direction {
	case "U":
		return Vector2{x: 0, y: 1}
	case "D":
		return Vector2{x: 0, y: -1}
	case "L":
		return Vector2{x: -1, y: 0}
	case "R":
		return Vector2{x: 1, y: 0}
	default:
		panic("Could not find direction!")
	}
}

func getIntAbs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func clampVector(vector Vector2, high Vector2, low Vector2) Vector2 {
	if vector.x > high.x {
		vector.x = high.x
	}
	if vector.x < low.x {
		vector.x = low.x
	}
	if vector.y > high.y {
		vector.y = high.y
	}
	if vector.y < low.y {
		vector.y = low.y
	}
	return vector
}
