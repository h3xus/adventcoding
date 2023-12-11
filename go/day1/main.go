package sumDigits

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func sumDigits(s string) string {
	var numz string
	var tempArr []int

	for _, char := range s {
		if unicode.IsDigit(char) {
			digit, err := strconv.Atoi(string(char))

			if err != nil {
				fmt.Println(err)
			} else {
				tempArr = append(tempArr, digit)
				numz = strconv.Itoa(tempArr[0]) + strconv.Itoa(tempArr[len(tempArr)-1])
			}
		}
	}
	return numz
}

func main() {
	file, err := os.Open("./lefile1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineSum := sumDigits(line)

		fmt.Printf("Line: %s, Double Number: %s\n", line, lineSum)
		linedigit, err := strconv.Atoi(string(lineSum))

		if err != nil {
			fmt.Println(err)
		} else {
			total += linedigit
		}

	}
	fmt.Printf("Total: %s\n", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
