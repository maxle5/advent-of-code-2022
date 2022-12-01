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

	scanner := bufio.NewScanner(readfile)

	sum := 0
	topThreeElves := [3]int{}
	for scanner.Scan() {
		calorieStr := scanner.Text()

		// if belongs to the same elf
		if len(calorieStr) > 0 {
			calorie, _ := strconv.Atoi(calorieStr)
			sum = sum + calorie
		} else {
			indexToReplace := FindIndexToReplace(sum, topThreeElves[:])
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
