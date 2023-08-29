package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
    Name string
    Size int

    Parent *Directory
    SubDirectories map[string]*Directory
}

func (d *Directory) AddSubDir(sub *Directory) {
    if _, ok := d.SubDirectories[sub.Name]; !ok {
        d.SubDirectories[sub.Name] = sub
    }
}

func (d *Directory) AddParentDir(par *Directory) {
    d.Parent = par
}

func (d *Directory) AccSize(size int) {
    d.Size += size
}

func (d *Directory) GetSize() int {
    totalSize := d.Size
    for _, sub := range d.SubDirectories {
        totalSize += sub.GetSize()
    }
    return totalSize
}

func NewDir(name string) *Directory {
    return &Directory{Name: name, SubDirectories: make(map[string]*Directory)}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    var curDir *Directory
    rootDir := NewDir("/")
    directories := []*Directory{rootDir}

	for bufScanner := bufio.NewScanner(file); bufScanner.Scan(); {
		line := bufScanner.Text()
        inputs := strings.Split(line, " ")

        if inputs[0] == "$" && inputs[1] == "cd" {
            directory := inputs[2]
            switch directory {
            case "/":
                curDir = rootDir
            case "..":
                curDir = curDir.Parent
            default:
                newDir := NewDir(directory)
                newDir.AddParentDir(curDir)
                curDir.AddSubDir(newDir)
                curDir = newDir
                directories = append(directories, newDir)
            }
        }

        if inputs[0] == "$" || inputs[0] == "dir" {
            continue
        }

        size, _ := strconv.Atoi(inputs[0])
        curDir.AccSize(size)
	}

    //totalSize := 0
    //for _, dir := range directories {
    //    if dir.GetSize() <= 100000 {
    //        totalSize += dir.GetSize()
    //    }
    //}
    //fmt.Println(totalSize)

    freeSpace := 70000000 - rootDir.GetSize()
    spaceNeeded := 30000000 - freeSpace
    deleteDirSize := 70000000000000000

    for _, dir := range directories {
        dirSize := dir.GetSize()

        if dirSize < spaceNeeded {
            continue
        }

        if dirSize < deleteDirSize {
            deleteDirSize = dirSize
        }
    }

    fmt.Println(deleteDirSize)
}
