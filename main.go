package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iambasantarai/ga/lexer"
	"github.com/iambasantarai/ga/token"
)

func main() {
	cliArgs := os.Args
	if len(cliArgs) < 2 {
		panic("Not enough arguments.")
	}

	sourceFilePath := cliArgs[1]
	fileExtension := filepath.Ext(sourceFilePath)
	if fileExtension != ".ga" {
		panic("Unrecognized file type. File must have '.ga' extension.")
	}

	sourceCode, err := os.ReadFile(sourceFilePath)
	if err != nil {
		panic(err)
	}

	lexer := lexer.New(string(sourceCode))

	for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
		fmt.Printf("TokenType: %s, Literal=%q\n", tok.Type, tok.Literal)
	}
}
