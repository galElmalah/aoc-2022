package token

type TokenType string

const (
	Cd   TokenType = "Cd"
	Ls   TokenType = "Ls"
	File TokenType = "File"
	Dir  TokenType = "Dir"
)

type Token struct {
	Type    TokenType
	Literal string
}
