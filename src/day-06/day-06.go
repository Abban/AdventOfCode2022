package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	filePath := dir + "/src/day-06/input.txt"

	fmt.Println("PART 1: ", findPositionOfFirstPacketWithUniqueCharacters(filePath, 4))
	fmt.Println("PART 2: ", findPositionOfFirstPacketWithUniqueCharacters(filePath, 14))
}

func findPositionOfFirstPacketWithUniqueCharacters(filePath string, packetSize int) int {
	file, _ := os.Open(filePath)
	reader := bufio.NewReader(file)
	buf := make([]byte, 1)

	i := 0
	packet := ""

	for {
		_, err := reader.Read(buf)
		if err != nil {
			break
		}
		i++

		character := string(buf)

		if i <= packetSize {
			packet += character
		} else {
			packet = packet[1:packetSize] + character
		}

		if i < 4 {
			continue
		}

		if !stringContainsDuplicateCharacter(packet) {
			break
		}
	}

	return i
}

func stringContainsDuplicateCharacter(input string) bool {
	for i := 0; i < len(input); i++ {
		if strings.Count(input, string(input[i])) > 1 {
			return true
		}
	}
	return false
}
