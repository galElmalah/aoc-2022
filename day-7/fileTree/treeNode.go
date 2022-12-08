package fileTree

import (
	"strconv"
	"strings"
)

type FileNode struct {
	name string
	size int
}

type DirNode struct {
	name string
}

type LsNode struct {
	dir string
}

type CdNode struct {
	to string
}

func praseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func CreateFileNode(str string) FileNode {
	parts := strings.Split(str, " ")
	return FileNode{
		name: parts[1],
		size: praseInt(parts[0]),
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
