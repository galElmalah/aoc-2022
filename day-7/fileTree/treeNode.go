package fileTree

import (
	"strings"

	"github.com/galElmalah/aoc-2022/util"
)

type FileNode struct {
	name string
	size int
}

type DirNode struct {
	name string
}

type CdNode struct {
	to string
}

func CreateFileNode(str string) FileNode {
	parts := strings.Split(str, " ")
	return FileNode{
		name: parts[1],
		size: util.ParseInt(parts[0]),
	}
}

func CreateDirNode(str string) DirNode {
	parts := strings.Split(str, " ")
	return DirNode{
		name: parts[1],
	}
}

func CreateCdNode(str string) CdNode {
	parts := strings.Split(str, " ")
	return CdNode{
		to: parts[2],
	}
}
