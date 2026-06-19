package parser

import (
	"fmt"


	"jsonParser/lexer"
	"jsonParser/token"
)

type Parser struct {
	lexer *lexer.Lexer
	currentToken token.Token
	peekToken token.Token 
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseObject() error {
	openingToken := p.currentToken 

	if openingToken.Type != token.LEFT_BRACE {
		return fmt.Errorf("Opening Token is not {")
	}

	p.nextToken()

	currentToken := p.currentToken

	if currentToken.Type == token.RIGHT_BRACE {
		return nil
	}

	if currentToken.Type != token.STRING {
		return fmt.Errorf("Current Token is not a string")
	}

	p.nextToken()

	currentToken = p.currentToken 

	if currentToken.Type != token.COLON {
		return fmt.Errorf("Current token is %v not : which is the expected", currentToken.Literal)
	}

	p.nextToken()

	currentToken = p.currentToken 

	if currentToken.Type != token.STRING {
		return fmt.Errorf("Current Token is not a string")
	}

	p.nextToken()

	currentToken =  p.currentToken

	if currentToken.Type != token.RIGHT_BRACE {
		return fmt.Errorf("Closing Token is not }")
	}

	return nil 
}

func (p *Parser) Parse() error{
	var err error

	err = p.parseObject()
	
	if err != nil {
		return err
	}

	p.nextToken()

	if p.currentToken.Type != token.EOF {
		return fmt.Errorf("Current Input is not a valid json string")
	}

	return nil
}