package main

import (
	"fmt"
	"os"

	"github.com/iambasantarai/ga/lexer"
	"github.com/iambasantarai/ga/token"
)

func main() {
	cliArgs := os.Args
	sourceFilePath := cliArgs[1]

	sourceCode, err := os.ReadFile(sourceFilePath)
	if err != nil {
		panic(err)
	}

	lexer := lexer.New(string(sourceCode))

	for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
		fmt.Printf("TokenType: %s, Literal=%q\n", tok.Type, tok.Literal)
	}
}
