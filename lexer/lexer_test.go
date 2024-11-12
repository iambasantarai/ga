package lexer

import (
	"testing"

	"github.com/iambasantarai/ga/token"
)

func TestNextToken(t *testing.T) {
	input := `पूर्णांक मुख्य() {
	       प्रदर्शन("सोच्छौ के मेरो बारे?")।
	       फिर्ता 0।
	   }`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INT, "पूर्णांक"},
		{token.MAIN, "मुख्य"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.PRINT, "प्रदर्शन"},
		{token.LPAREN, "("},
		{token.STRING_LITERAL, "सोच्छौ के मेरो बारे?"},
		{token.RPAREN, ")"},
		{token.TERMINATOR, "।"},
		{token.RETURN, "फिर्ता"},
		{token.INT, "0"},
		{token.TERMINATOR, "।"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tt.expectedType,
				tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - tokentype wrong. expected=%q, got=%q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
