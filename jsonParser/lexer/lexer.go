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


var SeenKey bool

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

func (l *Lexer) readString() string{
	str := ""
	for l.ch != '"' {
		str += string(l.ch)
		l.readChar()
	}
	return str
}

func (l *Lexer) NextToken() (token.Token) {
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
	case '"':	
		l.readChar()
		str := l.readString()
		tok =  token.Token{
			Type: token.STRING,
			Literal : str,
		} 
		SeenKey = true
	case ':':
		tok = token.Token{
			Type: token.COLON,
			Literal: string(l.ch),
		}		
	default:	
		tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}		
	}				

	l.readChar()
	return tok
}	