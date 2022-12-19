package day9

import (
	"advent/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Move int

const (
	Up Move = iota
	Down
	Left
	Right
)

type RopeGame struct {
	Tails       []*Position
	TailVisited []Position
}

func MakeRopeGame(knots int) *RopeGame {
	rg := RopeGame{
		TailVisited: make([]Position, 0),
		Tails:       make([]*Position, knots),
	}

	for i := 0; i < knots; i++ {
		rg.Tails[i] = &Position{}
	}

	return &rg
}

func (r *RopeGame) moveTails() {
	for i := 1; i < len(r.Tails); i++ {
		node := r.Tails[i]
		node.MoveTowards(*r.Tails[i-1])
		if i == len(r.Tails)-1 {
			r.TailVisited = append(r.TailVisited, *node)
		}
	}
}

func (r *RopeGame) Move(move Move) {
	head := r.Tails[0]
	switch move {
	case Up:
		head.Y -= 1
	case Down:
		head.Y += 1
	case Left:
		head.X -= 1
	case Right:
		head.X += 1
	}

	r.moveTails()

	// println("Head:", r.Head.X, r.Head.Y, "::", "Tail:", r.Tail.X, r.Tail.Y)
}

func (r *RopeGame) DoMoves(moves []Move) {
	for _, move := range moves {
		r.Move(move)
	}
}

func (r *RopeGame) GetGrid() [][]string {
	minY := 0
	maxY := 0
	minX := 0
	maxX := 0

	for _, p := range r.TailVisited {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	println("min-max, x, y", minX, maxX, minY, maxY)

	grid := make([][]string, maxY-minY+1)

	for i := 0; i < len(grid); i++ {
		row := make([]string, maxX-minX+1)
		for i := 0; i < maxX-minX+1; i++ {
			row[i] = "."
		}
		grid[i] = row
	}

	for _, p := range r.TailVisited {
		yCord := p.Y - minY
		xCord := p.X - minX
		grid[yCord][xCord] = "#"
	}

	return grid
}

type Position struct {
	X int
	Y int
}

func (p *Position) MoveTowards(head Position) {
	xDiff := head.X - p.X
	yDiff := head.Y - p.Y
	if xDiff == 0 && utils.Abs(yDiff) > 1 {
		p.Y += yDiff - (1 * yDiff / utils.Abs(yDiff))

	} else if yDiff == 0 && utils.Abs(xDiff) > 1 {
		p.X += xDiff - (1 * xDiff / utils.Abs(xDiff))

	} else if utils.Abs(yDiff) > 1 && utils.Abs(xDiff) > 1 {
		p.Y += yDiff - (1 * yDiff / utils.Abs(yDiff))
		p.X += xDiff - (1 * xDiff / utils.Abs(xDiff))

	} else if utils.Abs(yDiff) > 1 && utils.Abs(xDiff) > 0 {
		p.Y += yDiff - (1 * yDiff / utils.Abs(yDiff))
		p.X += xDiff

	} else if utils.Abs(yDiff) > 0 && utils.Abs(xDiff) > 1 {
		p.X += xDiff - (1 * xDiff / utils.Abs(xDiff))
		p.Y += yDiff
	}
}

func parseMoves(input []string) []Move {
	// match any move (UDRL) and number of them (0 to inf)
	r := regexp.MustCompile(`([UDRL])\s([0-9]*)`)

	moves := []Move{}

	for _, e := range input {
		res := r.FindAllStringSubmatch(e, -1)
		move := res[0][1]
		count, err := strconv.Atoi(res[0][2])
		utils.Check(err)

		var m Move
		switch move {
		case "U":
			m = Up
		case "D":
			m = Down
		case "R":
			m = Right
		case "L":
			m = Left
		}

		for i := 0; i < count; i++ {
			moves = append(moves, m)
		}
	}

	return moves
}

func Day9PartOne() int {
	b := MakeRopeGame(10)
	moves := parseMoves(utils.ReadFile("test_input/day9p2.txt"))
	b.DoMoves(moves)

	g := b.GetGrid()

	for _, row := range g {
		fmt.Println(row)
	}

	return 0
}
