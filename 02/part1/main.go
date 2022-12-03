package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func main() {
	readfile, err := os.Open("d:/source/advent-of-code-2022/02/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	totalScore := 0
	scanner := bufio.NewScanner(readfile)

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, " ")

		theirMove, _ := tryParse(result[0])
		myMove, _ := tryParse(result[1])

		totalScore = totalScore + myMove.BaseScore() + myMove.Score(theirMove)
	}

	log.Println(totalScore)
}

func tryParse(move string) (IMove, error) {
	if move == "X" || move == "A" {
		return Rock{}, nil
	} else if move == "Y" || move == "B" {
		return Paper{}, nil
	} else if move == "Z" || move == "C" {
		return Scissors{}, nil
	}

	return nil, errors.New("expected rock, paper, or scissors")
}

type IMove interface {
	BaseScore() int
	Score(opponentMove IMove) int
}

type Rock struct{}
type Paper struct{}
type Scissors struct{}

func (rock Rock) BaseScore() int {
	return 1
}
func (rock Rock) Score(opponentMove IMove) int {
	if _, ok := opponentMove.(Scissors); ok {
		return 6
	} else if _, ok := opponentMove.(Rock); ok {
		return 3
	}

	return 0
}

func (paper Paper) BaseScore() int {
	return 2
}
func (paper Paper) Score(opponentMove IMove) int {
	if _, ok := opponentMove.(Rock); ok {
		return 6
	} else if _, ok := opponentMove.(Paper); ok {
		return 3
	}

	return 0
}

func (scissors Scissors) BaseScore() int {
	return 3
}
func (scissors Scissors) Score(opponentMove IMove) int {
	if _, ok := opponentMove.(Paper); ok {
		return 6
	} else if _, ok := opponentMove.(Scissors); ok {
		return 3
	}

	return 0
}
