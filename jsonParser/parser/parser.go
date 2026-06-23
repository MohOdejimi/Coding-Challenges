package parser

import (
	"fmt"

	"jsonParser/lexer"
	"jsonParser/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
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

func (p *Parser) parseValue() error {
	if p.currentToken.Type == token.LEFT_BRACE {
		err := p.parseObject()
		return err
	}

	if p.currentToken.Type == token.LEFT_BRACKET {
		err := p.parseArray()
		return err
	}

	if p.currentToken.Type != token.BOOLEAN &&
		p.currentToken.Type != token.STRING &&
		p.currentToken.Type != token.NULL &&
		p.currentToken.Type != token.NUMBER {
		return fmt.Errorf("invalid value: unexpected token type %v", p.currentToken.Type)
	}
	return nil
}

func (p *Parser) parsePair() error {
	currentToken := p.currentToken

	if currentToken.Type != token.STRING {
		return fmt.Errorf("expected string key, got %v", currentToken.Type)
	}

	p.nextToken()
	currentToken = p.currentToken

	if currentToken.Type != token.COLON {
		return fmt.Errorf("expected ':' after key, got %q", currentToken.Literal)
	}

	p.nextToken()
	currentToken = p.currentToken

	err := p.parseValue()

	return err
}

func (p *Parser) parseArray() error {
	if p.currentToken.Type != token.LEFT_BRACKET {
		return fmt.Errorf("expected '[' to start array, got %v", p.currentToken.Type)
	}

	p.nextToken()
	if p.currentToken.Type == token.RIGHT_BRACKET {
		return nil
	}

	for {
		err := p.parseValue()
		if err != nil {
			return err
		}

		p.nextToken()

		if p.currentToken.Type == token.RIGHT_BRACKET {
			return nil
		} else if p.currentToken.Type == token.COMMA {
			p.nextToken()
		} else {
			return fmt.Errorf("unexpected token %v (%q)", p.currentToken.Type, p.currentToken.Literal)
		}
	}
}

func (p *Parser) parseObject() error {
	openingToken := p.currentToken

	if openingToken.Type != token.LEFT_BRACE {
		return fmt.Errorf("expected '{' to start object, got %v", openingToken.Type)
	}

	p.nextToken()
	currentToken := p.currentToken

	if currentToken.Type == token.RIGHT_BRACE {
		return nil
	}

	for {
		err := p.parsePair()

		if err != nil {
			return err
		}
		p.nextToken()

		currentToken = p.currentToken
		if currentToken.Type == token.RIGHT_BRACE {
			return nil
		} else if currentToken.Type == token.COMMA {
			p.nextToken()
		} else {
			return fmt.Errorf("unexpected token %v (%q)", currentToken.Type, currentToken.Literal)
		}
	}
}

func (p *Parser) Parse() error {
	var err error

	err = p.parseObject()

	if err != nil {
		return err
	}

	p.nextToken()

	if p.currentToken.Type != token.EOF {
		return fmt.Errorf("expected end of input, found %v", p.currentToken.Type)
	}

	return nil
}
