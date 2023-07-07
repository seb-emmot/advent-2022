package day07

import (
	"advent/utils"
	"regexp"
	"strconv"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []File
}

func (d *Directory) AddFile(name string, size int) {
	d.Files = append(d.Files, File{Name: name, Size: size})
}

func (d *Directory) MkDir(name string) {
	dir := NewDir(name)
	dir.Parent = d
	d.Children = append(d.Children, dir)
}

func (d *Directory) AddDir(dir *Directory) {
	dir.Parent = d
	d.Children = append(d.Children, dir)
}

func (d *Directory) GetDir(name string) *Directory {
	for _, dir := range d.Children {
		if dir.Name == name {
			return dir
		}
	}
	return nil
}

func (d *Directory) Pwd() string {
	if d.Parent == nil {
		return d.Name
	}
	return d.Parent.Pwd() + "/" + d.Name
}

func (d *Directory) Size() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}
	for _, dir := range d.Children {
		size += dir.Size()
	}
	return size
}

func (d *Directory) PrintTree(prefix string) {
	println(prefix, d.Name)
	for _, dir := range d.Children {
		dir.PrintTree(prefix + "  ")
	}
	for _, file := range d.Files {
		println(prefix+"   ", file.Name, file.Size)
	}
}

func NewDir(name string) *Directory {
	return &Directory{
		Name:     name,
		Children: make([]*Directory, 0),
		Files:    make([]File, 0),
	}
}

func buildDirs(input []string) *Directory {
	r := regexp.MustCompile(`(?m)^\$\s(.*)`)

	var curDir *Directory
	var root *Directory
	for _, line := range input {
		res := r.FindStringSubmatch(line)
		if res != nil {
			cmd := res[1][0:2]
			switch cmd {
			case "cd":
				arg := res[1][3:]
				if curDir == nil {
					curDir = NewDir(arg)
					root = curDir
				} else if arg == ".." {
					curDir = curDir.Parent
				} else {
					newDir := curDir.GetDir(arg)
					if newDir == nil {
						curDir.MkDir(arg)
						curDir = curDir.GetDir(arg)
					} else {
						curDir = newDir
					}
				}
			case "ls":
				// not needed.
			}
		} else {
			q := regexp.MustCompile(`^([0-9]*)\s(.*)`)
			match := q.FindStringSubmatch(line)
			if match != nil {
				size, err := strconv.Atoi(match[1])
				utils.Check(err)
				name := match[2]
				curDir.AddFile(name, size)
			}
		}
	}
	return root
}

func getSizeLessThan(dir *Directory, size int) int {
	sum := 0
	if dir.Size() < size {
		sum += dir.Size()
	}

	if dir.Children == nil {
		return sum
	} else {
		for _, child := range dir.Children {
			sum += getSizeLessThan(child, size)
		}
		return sum
	}
}

func day7PartOne(input []string) int {
	root := buildDirs(input)
	println("root size", root.Size())
	return getSizeLessThan(root, 100000)
}

func getBestChild(dir *Directory, size int) *Directory {
	best := dir
	if dir.Children != nil {
		for _, child := range dir.Children {
			bestChild := getBestChild(child, size)
			if bestChild.Size() >= size && bestChild.Size()-size < best.Size()-size {
				best = bestChild
			}
		}
	}
	return best
}

func day7PartTwo(input []string) int {
	root := buildDirs(input)

	totalSpace := 70000000
	freeSpace := totalSpace - root.Size()
	neededSpace := 30000000 - freeSpace

	println("need to clear", neededSpace, "bytes")
	bestChild := getBestChild(root, neededSpace)
	println("best child", bestChild.Name, bestChild.Size())
	return getBestChild(root, neededSpace).Size()
}
