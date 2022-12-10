package main

import (
	"testing"

	"github.com/galElmalah/aoc-2022/day-7/fileSystem"
	"github.com/galElmalah/aoc-2022/day-7/token"
	"github.com/galElmalah/aoc-2022/util"
)

func TestPart1(t *testing.T) {
	t.Run("Run Part 1 solution on example.txt input", func(t *testing.T) {
		data := util.ReadFile("./example.txt")

		got := Part1(data)
		want := 95437

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Run Part 2 solution on example.txt input", func(t *testing.T) {
		data := util.ReadFile("./example.txt")

		got := Part2(data)
		want := 24933642

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

}

func TestTokenizer(t *testing.T) {

	t.Run("Command", func(t *testing.T) {
		input := "$ cd /\n$ ls a\ndir a\n8504156 c.dat"

		got := token.Tokenize(input)
		want := []token.Token{{
			Type:    token.Cd,
			Literal: "$ cd /",
		}, {
			Type:    token.Ls,
			Literal: "$ ls a",
		},
			{
				Type:    token.Dir,
				Literal: "dir a",
			},
			{
				Type:    token.File,
				Literal: "8504156 c.dat",
			}}

		for i, tt := range got {

			if tt.Type != want[i].Type || tt.Literal != want[i].Literal {
				t.Errorf("got  %v want %v ", tt, want[i])
			}
		}
	})

}

func TestFileTree(t *testing.T) {

	t.Run("Command", func(t *testing.T) {
		input := "$ cd /\n$ ls\ndir a\ndir b\n50 c.dat\n$ cd a\n$ ls\n$ ls\n50 c.dat\n50 c.dat\n25 c.dat\n$ cd ..\n$ cd b\n$ ls\n50 c.dat$ cd .."

		got := fileSystem.NewFileSystem(token.Tokenize(input))

		str := ""
		got.Walk(func(t *fileSystem.FileSystemNode) {
			str += t.Name
		})

		if str != "/ab" {
			t.Errorf("got  %v want %v ", str, "/ab")
		}

	})

	t.Run("Holds correct size", func(t *testing.T) {
		input := "$ cd /\n$ ls\ndir a\ndir b\n50 c.dat\n$ cd a\n$ ls\n$ ls\n50 c.dat\n50 c.dat\n25 c.dat\n$ cd ..\n$ cd b\n$ ls\n50 c.dat$ cd .."

		got := fileSystem.NewFileSystem(token.Tokenize(input))

		if got.Size() != 225 {
			t.Errorf("got  %v want %v ", got.Size(), 225)
		}

	})

}
