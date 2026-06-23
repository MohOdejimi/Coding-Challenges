package main

import (
	"fmt"
	"os"

	"jsonParser/lexer"
	"jsonParser/parser"
)

func main() {
	args := os.Args[1:] 
	file := args[0]

	fileContent, err := os.ReadFile(file)
	if err != nil {
		err = fmt.Errorf("failed to read file %s: %v", file, err)
		fmt.Println(err)
		os.Exit(1)
	}
	input := string(fileContent)

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	result := parser.Parse()

	if result != nil {
		fmt.Println("Invalid JSON")
		fmt.Println(result.Error())
		os.Exit(1)
	} else {
		fmt.Println("Valid JSON")
		os.Exit(0)
	}
}
