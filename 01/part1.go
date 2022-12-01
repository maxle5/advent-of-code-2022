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
	largestSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// if belongs to the same elf
		if len(line) > 0 {
			calorie, _ := strconv.Atoi(line)
			sum = sum + calorie
		} else {
			if sum > largestSum {
				largestSum = sum
			}

			sum = 0
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	println(largestSum)
}
