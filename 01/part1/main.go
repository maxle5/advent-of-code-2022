package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	readfile, err := os.Open("01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	largestSum := 0
	scanner := bufio.NewScanner(readfile)

	// loop through lines in file
	for scanner.Scan() {
		// read the current line and attempt to convert to in
		calorie, err := strconv.Atoi(scanner.Text())

		// check if the current line is a number
		if err == nil {
			// keep a running total for 'this' elf
			sum = sum + calorie
		} else {
			// check if 'this' elf has a larger sum than the top contender
			largestSum = max(largestSum, sum)

			// reset the running total for the 'next' elf
			sum = 0
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	println(largestSum)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
