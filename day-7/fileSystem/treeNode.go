package fileSystem

import (
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

type FileNode struct {
	Name string
	Size int
}

type DirNode struct {
	Name string
}

type LsNode struct {
	Dir string
}

type CdNode struct {
	To string
}

func CreateFileNode(str string) FileNode {
	parts := strings.Split(str, " ")
	return FileNode{
		Name: parts[1],
		Size: util.ParseInt(parts[0]),
	}
}

func CreateDirNode(str string) DirNode {
	parts := strings.Split(str, " ")
	return DirNode{
		Name: parts[1],
	}
}

func CreateCdNode(str string) CdNode {
	parts := strings.Split(str, " ")
	return CdNode{
		To: parts[2],
	}
}
