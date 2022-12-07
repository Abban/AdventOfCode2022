package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + "/src/day-07/input.txt")
	scanner := bufio.NewScanner(file)

	sizes := map[string]int{"/": 0}
	folder := "/"

	for scanner.Scan() {
		text := scanner.Text()

		if shouldSkip(text) {
			continue
		}

		if isCommand(text) {
			if text == "$ cd .." {
				folder = removeEndFolderFromPath(folder)
			} else if text != "$ ls" {
				folder += "/" + getFolderName(text)
			}
			continue
		}

		if !isDir(text) {
			fileSize := getFileSize(text)
			sizes[folder] += fileSize
			parentFolder := folder

			for {
				parentFolder = removeEndFolderFromPath(parentFolder)
				if parentFolder == "" {
					break
				}
				sizes[parentFolder] += fileSize
			}
		}
	}

	part1(sizes)
	part2(sizes)
}

func part1(sizes map[string]int) {
	total := 0
	for folder, size := range sizes {
		if folder == "/" {
			continue
		}
		if size <= 100000 {
			total += size
		}
	}

	fmt.Println("TOTAL: ", total)
}

func part2(sizes map[string]int) {
	spaceNeeded := 30000000 - (70000000 - sizes["/"])

	smallest := 70000000
	for _, size := range sizes {
		if size > spaceNeeded && size < smallest {
			smallest = size
		}
	}

	fmt.Println("SMALLEST: ", smallest)
}

func getFolderName(text string) string {
	parts := strings.Split(text, " ")
	return parts[2]
}

func removeEndFolderFromPath(path string) string {
	index := strings.LastIndex(path, "/")
	return path[:index]
}

func getFileSize(text string) int {
	parts := strings.Split(text, " ")
	size, _ := strconv.Atoi(parts[0])
	return size
}

func shouldSkip(text string) bool {
	return len(text) == 0 || text == "$ cd /"
}

func isCommand(text string) bool {
	return strings.HasPrefix(text, "$")
}

func isDir(text string) bool {
	return strings.HasPrefix(text, "dir")
}
