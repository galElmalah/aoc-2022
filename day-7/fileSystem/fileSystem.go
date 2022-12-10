package fileSystem

import (
	"fmt"

	"github.com/galElmalah/aoc-2022/day-7/token"
)

type FileSystem struct {
	root *FileSystemNode
}

type FileSystemNode struct {
	Size   int
	Parent *FileSystemNode
	Token  token.Token
	Name   string
	Dirs   map[string]*FileSystemNode
}

func newFileSystemNode(name string, token token.Token, parent *FileSystemNode) *FileSystemNode {
	return &FileSystemNode{Name: name, Parent: parent, Token: token, Dirs: map[string]*FileSystemNode{}}
}

func NewFileSystem(tokens []token.Token) *FileSystem {
	root := newFileSystemNode("/", token.Token{Type: token.Dir, Literal: "/"}, nil)
	fileSystem := &FileSystem{root}
	current := root
	for _, t := range tokens {
		switch t.Type {
		case token.File:
			current.Size += CreateFileNode(t.Literal).Size
		case token.Dir:
			node := CreateDirNode(t.Literal)
			current.Dirs[node.Name] = newFileSystemNode(node.Name, t, current)
		case token.Ls:
			continue
		case token.Cd:
			cdNode := CreateCdNode(t.Literal)
			if cdNode.To == ".." {
				current.Parent.Size += current.Size
				current = current.Parent
			} else {
				_, ok := current.Dirs[cdNode.To]
				if ok {
					current = current.Dirs[cdNode.To]
				}
			}
		default:
			fmt.Println("Unexpected token!", t)
		}

	}

	// In case we are not ending at the root node
	for current != root {
		current.Parent.Size += current.Size
		current = current.Parent
	}

	return fileSystem
}

func (tree *FileSystem) Walk(visitor func(t *FileSystemNode)) {
	tree.root.Walk(visitor)
}

func (tree *FileSystem) Size() int {
	return tree.root.Size
}

func (node *FileSystemNode) Walk(visitor func(t *FileSystemNode)) {
	visitor(node)
	for _, t := range node.Dirs {
		t.Walk(visitor)
	}
}
