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
	topThreeElves := [3]int{}
	scanner := bufio.NewScanner(readfile)

	// loop through the lines in file
	for scanner.Scan() {
		// read the current line and attempt to convert to in
		calorie, err := strconv.Atoi(scanner.Text())

		// check if the current line is a number
		if err == nil {
			// keep a running total for 'this' elf
			sum = sum + calorie
		} else {
			// check if 'this' elf is in the top 3
			indexToReplace := FindIndexToReplace(sum, topThreeElves[:])

			// replace an elf if 'this' elf is in the top 3
			if indexToReplace > -1 {
				topThreeElves[indexToReplace] = sum
			}

			sum = 0
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	println(Sum(topThreeElves[:]))
}

func FindIndexToReplace(calorie int, topThreeElves []int) int {
	for elfIndex, elfCalories := range topThreeElves {
		if calorie > elfCalories {
			return elfIndex
		}
	}

	return -1
}

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum = sum + val
	}

	return sum
}
