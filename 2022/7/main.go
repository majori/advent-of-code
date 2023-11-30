package main

import (
	"advent-of-code/common/aoc"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Dir struct {
	Name    string
	Parent  *Dir
	SubDirs map[string]*Dir
	Files   map[string]*File
}

type File struct {
	Name string
	Size int
}

var p1 int
var p2 []int

func main() {
	task := aoc.Init(2022, 7)
	input := task.AsStringSlice()

	cdReg := regexp.MustCompile(`^\$ cd (.+)`)
	lsReg := regexp.MustCompile(`^\$ ls`)
	dirReg := regexp.MustCompile(`^dir (.+)`)
	fileReg := regexp.MustCompile(`^(\d+) (.+)`)

	var currentDir, rootDir *Dir
	for _, row := range input {
		if cdReg.MatchString(row) {
			name := cdReg.FindStringSubmatch(row)[1]
			if name == "/" {
				rootDir = &Dir{Name: name}
				currentDir = rootDir
			} else if name == ".." {
				currentDir = currentDir.Parent
			} else {
				currentDir = currentDir.SubDirs[name]

			}
		} else if lsReg.MatchString(row) {
		} else if dirReg.MatchString(row) {
			if currentDir.SubDirs == nil {
				currentDir.SubDirs = make(map[string]*Dir)
			}
			name := dirReg.FindStringSubmatch(row)[1]
			currentDir.SubDirs[name] = &Dir{
				Name:   name,
				Parent: currentDir,
			}
		} else if fileReg.MatchString(row) {
			if currentDir.Files == nil {
				currentDir.Files = make(map[string]*File)
			}
			f := fileReg.FindStringSubmatch(row)
			size, _ := strconv.Atoi(f[1])
			name := f[2]
			currentDir.Files[name] = &File{
				Name: name,
				Size: size,
			}
		}
	}

	calculateDirSize(rootDir)
	fmt.Println("Puzzle 1:", p1)

	// unusedSpace := 70_000_000 - totalSize
	// requiredSpace := 30_000_000 - unusedSpace

	sort.Ints(p2)
	fmt.Println("Puzzle 2:", p2[0])
}

func calculateDirSize(dir *Dir) int {
	var size int
	for _, file := range dir.Files {
		size += file.Size
	}

	for _, subDir := range dir.SubDirs {
		size += calculateDirSize(subDir)
	}

	if size <= 100000 {
		p1 += size
	}

	if size >= 10822529 {
		if p2 == nil {
			p2 = make([]int, 0)
		}
		p2 = append(p2, size)
	}

	return size
}
