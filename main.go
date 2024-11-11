package main

import (
	"fmt"

	"github.com/iambasantarai/ga/lexer"
	"github.com/iambasantarai/ga/token"
)

func main() {
	input := "पूर्णांक मुख्य() { फिर्ता 0; }"

	lexer := lexer.New(input)

	for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
		fmt.Printf("TokenType: %s, Literal=%q\n", tok.Type, tok.Literal)
	}
}
