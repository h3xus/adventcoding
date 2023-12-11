package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	targetCubes := map[string]int{"red": 12, "green": 13, "blue": 14}

	file, err := os.Open("lefile2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumOfIDs := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		gameID, err := strconv.Atoi(parts[0][5:])
		if err != nil {
			fmt.Println("Error converting game ID:", err)
			return
		}

		subsets := strings.Split(parts[1], "; ")

		possible := true
		for _, subset := range subsets {
			cubeCounts := make(map[string]int)
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				parts := strings.Split(cube, " ")
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Println("Error converting cube count:", err)
					return
				}
				color := parts[1]
				cubeCounts[color] += count
			}

			for color, count := range cubeCounts {
				if count > targetCubes[color] {
					possible = false
					break
				}
			}

			if !possible {
				break
			}
		}

		if possible {
			sumOfIDs += gameID
		}
	}

	// Print the result
	fmt.Println("Sum of IDs of possible games:", sumOfIDs)
}
