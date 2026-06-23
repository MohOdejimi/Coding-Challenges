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

func (l *Lexer) readString() string{
	str := ""
	for l.ch != '"' {
		str += string(l.ch)
		l.readChar()
	}
	return str
}

func (l *Lexer) readTrue() bool {
	str := "true"
		for _, ch := range str {
			if rune(l.ch) != ch {
				return false
			}
			l.readChar()
	}
	return true
}

func (l *Lexer) readFalse() bool {
	str := "false"
	for _, ch := range str {
		if rune(l.ch) != ch {
			return false
		}
		l.readChar()
	}
	return true
}

func (l *Lexer) readNull() bool {
	str := "null" 
	for _, ch := range str {
		if rune(l.ch) !=  ch {
			return false
		}
		l.readChar()
	}
	return true
}

func (l *Lexer) isNum() bool {
	return l.ch >= '0' && l.ch <= '9'
}

func (l *Lexer) readNum() string {
	num := ""
	for l.isNum() {
		num += string(l.ch)
		l.readChar()
	}
	return num
}


func (l *Lexer) NextToken() (token.Token) {
	var tok token.Token	
		
	l.skipWhitespace()

	switch l.ch {	

	case '{':
		tok = token.Token{
			Type: token.LEFT_BRACE, 
			Literal: string(l.ch),
		}				
	case '}':
		tok = token.Token{
			Type: token.RIGHT_BRACE,
			Literal: string(l.ch),
		}
	case '[': 
		tok =  token.Token{
			Type: token.LEFT_BRACKET,
			Literal: string(l.ch),
		}
	case ']': 
		tok = token.Token{
			Type: token.RIGHT_BRACKET,
			Literal: string(l.ch),
		}			
	case 0:
		tok = token.Token{
			Type: token.EOF,
			Literal: "",
		}
		return tok
	case '"':	
		l.readChar()
		str := l.readString()
		tok =  token.Token{
			Type: token.STRING,
			Literal : str,
		} 
	case ':':
		tok = token.Token{
			Type: token.COLON,
			Literal: string(l.ch),
		}
	case ',': 
		tok = token.Token{
			Type: token.COMMA,
			Literal: string(l.ch),
		}	
	case 't': 
			isTrue := l.readTrue()
			if isTrue {
				tok = token.Token{
					Type: token.BOOLEAN,
					Literal: string("true"),
				}
			} else {
				tok =  token.Token{
					Type: token.ILLEGAL,
					Literal: "ILLEGAL",
				}
			}
			return tok
	case 'f': 
			isNotTrue := l.readFalse()
			if isNotTrue {
				tok =  token.Token{
					Type: token.BOOLEAN,
					Literal: "false",
				}
			} else {
				tok = token.Token{
					Type: token.ILLEGAL,
					Literal: "ILLEGAL",
				}
			}
			return tok
	case 'n': 
			isNull := l.readNull()
			if isNull {
				tok = token.Token{
					Type: token.NULL,
					Literal: "null",
				}
			} else {
				tok = token.Token{
					Type: token.ILLEGAL,
					Literal: "ILLEGAL",
				}
			}
			return tok 
	default:	
		num := l.readNum() 
		if num != "" {
			tok = token.Token{
				Type: token.NUMBER,
				Literal: num,
			}
			return tok
		}
		tok = token.Token{
			Type: token.ILLEGAL, 
			Literal: string(l.ch),
		}		
	}				

	l.readChar()
	return tok
}	