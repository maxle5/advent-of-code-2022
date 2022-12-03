package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	readfile, err := os.Open("d:/source/advent-of-code-2022/02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	baseScoreMap := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
	}
	totalScore := 0
	scanner := bufio.NewScanner(readfile)

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, " ")

		theirMove := rune(result[0][0])
		myMove := map[rune]rune{
			'X': 'A',
			'Y': 'B',
			'Z': 'C',
		}[rune(result[1][0])]

		totalScore = totalScore + baseScoreMap[myMove] + score(theirMove, myMove)
	}

	println(totalScore)
}

func score(theirMove rune, myMove rune) int {
	winningMove := map[rune]rune{
		'A': 'B',
		'B': 'C',
		'C': 'A',
	}[theirMove]

	if myMove == winningMove {
		return 6
	} else if myMove == theirMove {
		return 3
	}

	return 0
}
