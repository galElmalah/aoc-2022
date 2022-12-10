package token

import "strings"

func newToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func Tokenize(raw string) []Token {
	lines := strings.Split(string(raw), "\n")
	tokens := []Token{}
	for _, l := range lines {
		parts := strings.Split(l, " ")
		if parts[0] == "$" {
			if parts[1] == "ls" {
				tokens = append(tokens, newToken(Ls, l))
			} else {
				tokens = append(tokens, newToken(Cd, l))
			}
		} else {
			if parts[0] == "dir" {
				tokens = append(tokens, newToken(Dir, l))
			} else {
				tokens = append(tokens, newToken(File, l))
			}
		}
	}
	return tokens
}
