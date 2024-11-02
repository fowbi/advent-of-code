package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("No input provided")
	}
}

func part_01(input string) int {
	games := parseInput(input)

	score := 0
	for _, game := range games {
		me := Me[game.Me]
		if me == game.Opponent {
			score += 3
		}

		if WinningOpposite[me] == game.Opponent {
			score += 6
		}

		score += int(me)
	}

	return score
}

func part_02(input string) int {
	games := parseInput(input)

	score := 0
	for _, game := range games {
		switch game.Me {
		case "X":
			score += int(WinningOpposite[game.Opponent])
		case "Y":
			score += 3 + int(game.Opponent)
		case "Z":
			score += 6 + int(LosingOpposite[game.Opponent])
		}
	}

	return score
}

func main() {
	fmt.Println("Solution part 1:", part_01(input))
	fmt.Println("Solution part 2:", part_02(input))
}

type HandShape int

const (
	HandShapeRock HandShape = iota + 1
	HandShapePaper
	HandShapeScissors
)

var Opponent = map[string]HandShape{
	"A": HandShapeRock,
	"B": HandShapePaper,
	"C": HandShapeScissors,
}
var Me = map[string]HandShape{
	"X": HandShapeRock,
	"Y": HandShapePaper,
	"Z": HandShapeScissors,
}

var LosingOpposite = map[HandShape]HandShape{
	HandShapeRock:     HandShapePaper,
	HandShapePaper:    HandShapeScissors,
	HandShapeScissors: HandShapeRock,
}

var WinningOpposite = map[HandShape]HandShape{
	HandShapeRock:     HandShapeScissors,
	HandShapePaper:    HandShapeRock,
	HandShapeScissors: HandShapePaper,
}

type Game struct {
	Opponent HandShape
	Me       string
}

func parseInput(input string) (games []Game) {
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		players := strings.SplitN(row, " ", 2)

		games = append(games, Game{
			Opponent: Opponent[players[0]],
			Me:       players[1],
		})
	}

	return games
}
