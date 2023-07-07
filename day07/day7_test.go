package day07

import (
	"advent/utils"
	"testing"
)

func TestDay7PartOne(t *testing.T) {
	input := utils.ReadFile("../test_input/day7test.txt")

	actual := day7PartOne(input)
	expected := 95437

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDay7PartTwo(t *testing.T) {
	input := utils.ReadFile("../test_input/day7test.txt")

	actual := day7PartTwo(input)
	expected := 24933642

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDirectory(t *testing.T) {
	dir := NewDir("root")

	dir.AddFile("file1", 10)
	dir.AddFile("file2", 15)

	// assert size is correct
	if dir.Size() != 25 {
		t.Errorf("Expected size to be 10, got %d", dir.Size())
	}
}

func TestAddDir(t *testing.T) {
	dir := NewDir("root")
	dir2 := NewDir("dir2")
	dir2.AddFile("file1", 10)
	dir.AddDir(dir2)
	dir.AddFile("file2", 15)

	// assert size is correct
	if dir.Size() != 25 {
		t.Errorf("Expected size to be 0, got %d", dir.Size())
	}
}

func TestPwd(t *testing.T) {
	dir := NewDir("root")
	dir.MkDir("dir2")

	workDir := dir.GetDir("dir2")

	if workDir.Pwd() != "root/dir2" {
		t.Errorf("Expected pwd to be root/dir2, got %s", workDir.Pwd())
	}
}
