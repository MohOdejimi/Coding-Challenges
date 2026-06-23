package token 


type TokenType string 

const (
	LEFT_BRACE TokenType = "LEFT_BRACE"
	RIGHT_BRACE TokenType = "RIGHT_BRACE"
	LEFT_BRACKET TokenType = "LEFT_BRACKET"
	RIGHT_BRACKET TokenType = "RIGHT_BRACKET"
	EOF TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"
	STRING TokenType = "STRING"
	COLON TokenType = "COLON"
	COMMA TokenType = "COMMA"
	BOOLEAN TokenType = "BOOLEAN"
	NULL TokenType = "NULL"
	NUMBER TokenType = "NUMBER"
)

type Token struct {
	Position int 
	Type TokenType
	Literal string
}