package day13

import (
	"adventofcode-2025/common"
	"math"
)

type button struct {
	x, y int
}

type prizeLocation struct {
	x, y int
}

type multiplier struct {
	a, b int
}

type game struct {
	buttonA button
	buttonB button
	prize   prizeLocation
	tokens  int
	valid   bool
}

func calculateMCM(a, b, c int) (int, int) {
	d := 0
	e := 0

	multipliers := make([]multiplier, 0)

	for d < 100 {
		e = 0
		for e < 100 {
			if (a*d)+(b*e) == c {
				multipliers = append(multipliers, multiplier{d, e})
			} else if (a*d)+(b*e) > c {
				break
			}

			e++
		}
		d++
	}

	// finding cheapest multipliers
	min := math.MaxInt
	for _, m := range multipliers {
		if m.a+m.b < min {
			min = m.a + m.b
			d = m.a
			e = m.b
		}
	}

	if d < 0 || e < 0 {
		panic("invalid multipliers")
	}

	return d, e
}

func calculateCost(a, b int) int {
	return a*3 + b*1
}

func parseGames(lines []string) []game {
	games := make([]game, 0)
	for i := 0; i < len(lines); i += 4 {
		g := game{}

		str := common.ReplaceStr(lines[i], []string{"Button A: X+", " Y+"}, "")
		parts := common.SplitStr(str, ":,")
		g.buttonA = button{x: common.ParseStrToInt(parts[0]), y: common.ParseStrToInt(parts[1])}

		str = common.ReplaceStr(lines[i+1], []string{"Button B: X+", " Y+"}, "")
		parts = common.SplitStr(str, ":,")
		g.buttonB = button{x: common.ParseStrToInt(parts[0]), y: common.ParseStrToInt(parts[1])}

		str = common.ReplaceStr(lines[i+2], []string{"Prize: X=", " Y="}, "")
		parts = common.SplitStr(str, ":,")
		g.prize = prizeLocation{x: common.ParseStrToInt(parts[0]), y: common.ParseStrToInt(parts[1])}

		g.valid = true

		a, b := calculateMCM(g.buttonA.x, g.buttonB.x, g.prize.x)
		c, d := calculateMCM(g.buttonA.y, g.buttonB.y, g.prize.y)
		g.valid = !(a > 100 || b > 100) && (a*g.buttonA.x)+(b*g.buttonB.x) == g.prize.x
		g.valid = g.valid && !(c > 100 || d > 100) && (a*g.buttonA.y)+(b*g.buttonB.y) == g.prize.y

		g.tokens = calculateCost(a, b)

		games = append(games, g)
	}

	return games
}

func Puzzle1() int {
	lines := common.ReadFileByLines("./inputs/day13.txt")

	games := parseGames(lines)

	sum := 0
	for _, g := range games {
		if g.valid {
			sum += g.tokens
		}
	}
	// = 39996
	return sum
}
