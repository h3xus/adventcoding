package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("lefile2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfPowers := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		subsets := strings.Split(parts[1], "; ")

		minRed, minGreen, minBlue := int(^uint(0)>>1), int(^uint(0)>>1), int(^uint(0)>>1)

		for _, subset := range subsets {
			cubes := strings.Split(subset, ", ")
			for _, cube := range cubes {
				parts := strings.Split(cube, " ")
				count, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Println("Error converting cube count:", err)
					return
				}
				color := parts[1]

				switch color {
				case "red":
					if count < minRed {
						minRed = count
					}
				case "green":
					if count < minGreen {
						minGreen = count
					}
				case "blue":
					if count < minBlue {
						minBlue = count
					}
				}
			}
		}

		power := minRed * minGreen * minBlue

		sumOfPowers += power
	}

	fmt.Println("Sum of the powers of minimum sets:", sumOfPowers)
}
