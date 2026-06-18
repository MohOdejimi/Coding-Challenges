package token 


type TokenType string 

const (
	LEFT_BRACE TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	EOF TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
)

type Token struct {
	Position int 
	Type TokenType
	Literal string
}