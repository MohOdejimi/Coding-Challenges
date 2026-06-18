package main 


import (
	"fmt"
	"os"

	"jsonParser/lexer"
	"jsonParser/token"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: jsonParser <file.json>")
		return
	}

	input := args[0]
	lexer := lexer.New(input)
	for {
		tok := lexer.NextToken()
		fmt.Printf("Token: %v, Type: %v, Literal: %v\n", tok, tok.Type, tok.Literal)
		if tok.Type == token.EOF {
			break
		}
	}
}
