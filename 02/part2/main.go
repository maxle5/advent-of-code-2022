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

		loseMap := map[rune]rune{
			'A': 'C',
			'B': 'A',
			'C': 'B',
		}
		drawMap := map[rune]rune{
			'A': 'A',
			'B': 'B',
			'C': 'C',
		}
		winMap := map[rune]rune{
			'A': 'B',
			'B': 'C',
			'C': 'A',
		}
		winLoseMap := map[rune](map[rune]rune){
			'X': loseMap,
			'Y': drawMap,
			'Z': winMap,
		}
		theirMove := rune(result[0][0])
		myMove := winLoseMap[rune(result[1][0])][theirMove]

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
