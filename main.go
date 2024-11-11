package main

import "fmt"

func main() {
	input := "पूर्णांक मुख्य() { फिर्ता 0; }"

	lexer := New(input)

	for tok := lexer.NextToken(); tok.Type != EOF; tok = lexer.NextToken() {
		fmt.Printf("TokenType: %s, Literal=%q\n", tok.Type, tok.Literal)
	}
}
