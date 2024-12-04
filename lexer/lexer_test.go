package lexer

import (
	"fmt"
	"testing"

	"github.com/iambasantarai/ga/token"
)

func TestNextToken(t *testing.T) {
	t.Run("test keyword tokens", func(t *testing.T) {
		input := `पूर्णांक दशमलव अक्षर मुख्य शून्य प्रदर्शन फिर्ता यदि अन्यथा रोक सही गलत `

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.INT, "पूर्णांक"},
			{token.FLOAT, "दशमलव"},
			{token.CHAR, "अक्षर"},
			{token.MAIN, "मुख्य"},
			{token.VOID, "शून्य"},
			{token.PRINT, "प्रदर्शन"},
			{token.RETURN, "फिर्ता"},
			{token.IF, "यदि"},
			{token.ELSE, "अन्यथा"},
			{token.BREAK, "रोक"},
			{token.TRUE, "सही"},
			{token.FALSE, "गलत"},
		}

		l := New(input)

		for i, tt := range tests {
			tok := l.NextToken()
			fmt.Println(tok)

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

	t.Run("test a portion of source code", func(t *testing.T) {
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
			{token.OPAREN, "("},
			{token.CPAREN, ")"},
			{token.OCURLY, "{"},
			{token.PRINT, "प्रदर्शन"},
			{token.OPAREN, "("},
			{token.STRING_LITERAL, "सोच्छौ के मेरो बारे?"},
			{token.CPAREN, ")"},
			{token.TERMINATOR, "।"},
			{token.RETURN, "फिर्ता"},
			{token.INT, "0"},
			{token.TERMINATOR, "।"},
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
