package day10

import (
	"advent/utils"
	"regexp"
	"strconv"
)

type Instr int

const (
	Addx Instr = iota
	Noop
)

type Cmd struct {
	Instr Instr
	Args  string
}

type Computer struct {
	cycle    int
	register int
	History  []int
}

func (c *Computer) RunProgram(program []Cmd) int {
	total := 0
	waitCycle := 0
	storedAdd := 0

	i := 0

	c.History = append(c.History, c.register)
	println("cycle", c.cycle, "reg", c.register)

	for i < len(program) {
		cmd := program[i]
		c.cycle += 1

		if (c.cycle+20)%40 == 0 {
			println("saving cycle", c.cycle, "reg", c.register)
			total += (c.cycle * c.register)
		}

		if waitCycle > 1 {
			waitCycle--
		} else if waitCycle == 1 {
			waitCycle--
			c.register += storedAdd
		} else {
			switch cmd.Instr {
			case Addx:
				x, err := strconv.Atoi(cmd.Args)
				utils.Check(err)
				storedAdd = x
				waitCycle = 1
				println("queue add at pos", i, storedAdd)
			case Noop:
				// do nothings
			}
			i++
		}
		c.History = append(c.History, c.register)
		println("cycle", c.cycle, "reg", c.register)
	}
	return total
}

func NewComputer() *Computer {
	return &Computer{
		cycle:    0,
		register: 1,
		History:  make([]int, 0),
	}
}

func parseInstructions(input []string) []Cmd {
	r := regexp.MustCompile(`([a-z]*).*`)

	cmds := []Cmd{}
	for _, line := range input {
		matches := r.FindAllStringSubmatch(line, -1)
		switch matches[0][1] {
		case "addx":
			argex := regexp.MustCompile(`([a-z]*).*\s([-0-9]*)`)
			arg := argex.FindAllStringSubmatch(line, -1)[0][2]
			cmds = append(cmds, Cmd{Instr: Addx, Args: arg})
		case "noop":
			cmds = append(cmds, Cmd{Instr: Noop})
		}
	}
	return cmds
}

func renderCRT(hist []int) [][]string {
	yHeight := len(hist) / 40
	xHeight := 40
	screen := make([][]string, yHeight)

	for i := 0; i < len(screen); i++ {
		screen[i] = make([]string, xHeight)
	}

	// spritePos represents end of cycle value
	for i := 1; i < len(hist); i++ {
		var spritePos int
		spritePos = hist[i-1] // value at end of previous cycle
		yPos := (i - 1) / 40
		xPos := (i - 1) % 40
		if utils.Abs(spritePos-xPos) < 2 {
			screen[yPos][xPos] = "#"
		} else {
			screen[yPos][xPos] = "."
		}
	}

	for _, line := range screen {
		for _, col := range line {
			print(col)
		}
		println()
	}

	return screen
}
