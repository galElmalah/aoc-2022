package fileTree

import (
	"fmt"

	"github.com/galElmalah/aoc-2022/day-7/token"
)

type FileTree struct {
	Size   int
	Parent *FileTree
	Token  token.Token
	name   string
	Sub    []*FileTree
}

func CreateFileTree(tokens []token.Token) *FileTree {
	root := &FileTree{name: "/", Token: token.Token{Type: token.Dir, Literal: "/"}, Sub: []*FileTree{}}
	current := root
	for _, t := range tokens {
		switch t.Type {
		case token.File:
			current.Size += CreateFileNode(t.Literal).size

		case token.Dir:
			node := CreateDirNode(t.Literal)
			current.Sub = append(current.Sub, &FileTree{name: node.name, Token: t, Parent: current, Sub: []*FileTree{}})
		case token.Ls:
			continue
		case token.Cd:
			cdNode := CreateCdNode(t.Literal)
			if cdNode.to == ".." {
				current.Parent.Size += current.Size
				current = current.Parent
			} else {
				for _, c := range current.Sub {
					if cdNode.to == c.name {
						current = c
					}
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

	return root
}

func (tree *FileTree) Walk(visitor func(t *FileTree)) {
	visitor(tree)
	for _, t := range tree.Sub {
		t.Walk(visitor)
	}
}
