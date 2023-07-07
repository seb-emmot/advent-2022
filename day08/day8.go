package day08

import (
	"advent/utils"
	"strconv"
	"strings"
)

type HeightMap struct {
	Grid [][]int
}

func NewHeightMap(input []string) *HeightMap {
	dy := len(input)
	grid := make([][]int, dy)

	for x, line := range input {
		dx := len(line)
		grid[x] = make([]int, dx)
		for y, char := range line {
			val, err := strconv.Atoi(string(char))
			utils.Check(err)
			grid[x][y] = val
		}
	}

	return &HeightMap{Grid: grid}
}

func (h *HeightMap) ComputeVisible() []utils.Point {
	visible := []utils.Point{}

	for y := 1; y <= len(h.Grid)-2; y++ {
		// print("checking line ", y, ". ")
		row := h.Grid[y]
		maxLeft := row[0]
		maxRight := row[len(row)-1]
		for x := 1; x <= len(row)-2; x++ {
			leftCandidate := row[x]
			if leftCandidate > maxLeft {
				visible = append(visible, utils.Point{X: x, Y: y})
				maxLeft = leftCandidate
				// print(leftCandidate, " visible left ")
			}
			rightCandidate := row[len(row)-(1+x)]
			if rightCandidate > maxRight {
				pt := utils.Point{X: len(row) - (1 + x), Y: y}
				visible = append(visible, pt)
				maxRight = rightCandidate
				// print(rightCandidate, " visible right ")
			}
		}
		// println()
	}

	for x := 1; x <= len(h.Grid[0])-2; x++ {
		// print("checking col ", x, ". ")
		col := getCol(h.Grid, x)
		maxTop := col[0]
		maxBot := col[len(col)-1]
		for y := 1; y <= len(col)-2; y++ {
			topCandidate := col[y]
			if topCandidate > maxTop {
				visible = append(visible, utils.Point{X: x, Y: y})
				maxTop = topCandidate
				// print(topCandidate, " visible top ")
			}
			botCandidate := col[len(col)-(1+y)]
			if botCandidate > maxBot {
				pt := utils.Point{X: x, Y: len(col) - (1 + y)}
				visible = append(visible, pt)
				maxBot = botCandidate
				// print(botCandidate, " visible bot ")
			}
		}
		// println()
	}

	x := utils.Set(visible)
	return x
	// return len(visible)
}

func getCol(input [][]int, index int) []int {
	col := []int{}
	for _, row := range input {
		col = append(col, row[index])
	}
	return col
}

func (h *HeightMap) Circumferance() int {
	return ((len(h.Grid) + len(h.Grid[0])) * 2) - 4
}

func (h *HeightMap) Print() {
	for i, row := range h.Grid {
		if i == 0 {
			println(strings.Repeat("_", len(row)+2))
		}
		print("|")
		for _, k := range row {
			print(k)
		}
		print("|\n")
	}
}

func printBlocked(input []utils.Point) {
	maxX := 0
	maxY := 0

	for _, p := range input {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	pMap := make([][]int, maxY+1)

	for i := 0; i <= maxY; i++ {
		pMap[i] = make([]int, maxX+1)
	}

	for _, p := range input {
		pMap[p.Y][p.X] = 1
	}

	q := &HeightMap{pMap}
	q.Print()
}

func day8PartOne(input []string) int {
	hm := NewHeightMap(input)
	hm.Print()

	points := hm.ComputeVisible()
	totalVis := hm.Circumferance() + len(points)

	printBlocked(points)

	return totalVis
}

func (h *HeightMap) ScenicScore(x int, y int) int {
	tilesNorth := 1
	tilesSouth := 1
	tilesWest := 1
	tilesEast := 1

	myTree := h.Grid[y][x]

	for i := 1; y-i > 0; i++ {
		if myTree > h.Grid[y-i][x] {
			tilesNorth++
		} else {
			break
		}
	}

	for i := 1; y+i < len(h.Grid)-1; i++ {
		if myTree > h.Grid[y+i][x] {
			tilesSouth++
		} else {
			break
		}
	}

	for i := 1; x-i > 0; i++ {
		if myTree > h.Grid[y][x-i] {
			tilesWest++
		} else {
			break
		}
	}

	for i := 1; x+i < len(h.Grid)-1; i++ {
		if myTree > h.Grid[y][x+i] {
			tilesEast++
		} else {
			break
		}
	}

	return tilesNorth * tilesSouth * tilesEast * tilesWest
}

func (h *HeightMap) ScenicScores() []int {
	scores := []int{}
	for y := 1; y < len(h.Grid)-1; y++ {
		for x := 1; x < len(h.Grid[0])-1; x++ {
			score := h.ScenicScore(x, y)
			scores = append(scores, score)
		}
	}

	return scores
}

func day8PartTwo(input []string) int {
	hm := NewHeightMap(input)

	scores := hm.ScenicScores()
	maxScore := utils.MaxArray(scores)

	return maxScore
}
