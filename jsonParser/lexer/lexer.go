package lexer 

import (
	"jsonParser/token"
)

type Lexer struct {
	input string
	position int 
	readPosition int 
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token	
		
	l.skipWhitespace()

	switch l.ch {	

	case '{':
		tok = token.Token{Type: token.LEFT_BRACE, Literal: string(l.ch)}				
	case '}':
		tok = token.Token{Type: token.RIGHT_BRACE, Literal: string(l.ch)}	
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
		return tok
	default:	
		tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}		
	}				

	l.readChar()
	return tok
}	