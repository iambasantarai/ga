package lexer

import (
	"testing"

	"github.com/iambasantarai/ga/token"
)

func TestNextToken(t *testing.T) {
	t.Run("test a portion of source code", func(t *testing.T) {
		input := `कार्य मुख्य() {
            ("सोच्छौ के मेरो बारे?") छाप ।
        }`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.FUNC, "कार्य"},
			{token.IDENT, "मुख्य"},
			{token.OPAREN, "("},
			{token.CPAREN, ")"},
			{token.OCURLY, "{"},
			{token.OPAREN, "("},
			{token.STRING_LITERAL, "सोच्छौ के मेरो बारे?"},
			{token.CPAREN, ")"},
			{token.PRINT, "छाप"},
			{token.DELIMITER, "।"},
			{token.CCURLY, "}"},
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
	})
}
