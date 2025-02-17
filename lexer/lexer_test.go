package lexer

import (
	"testing"

	"ga/token"
)

func TestNextToken(t *testing.T) {
	t.Run("test a portion of source code", func(t *testing.T) {
		input := `कार्य मुख्य() {
            छाप("सोच्छौ के मेरो बारे?")।
        }`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.KAARYA, "कार्य"},
			{token.IDENT, "मुख्य"},
			{token.KHULLAA_SAANO_KOSHTHAK, "("},
			{token.BANDA_SAANO_KOSHTHAK, ")"},
			{token.KHULLAA_MADHYAM_KOSHTHAK, "{"},
			{token.CHAAPA, "छाप"},
			{token.KHULLAA_SAANO_KOSHTHAK, "("},
			{token.BAAKYA, "सोच्छौ के मेरो बारे?"},
			{token.BANDA_SAANO_KOSHTHAK, ")"},
			{token.PURNA_BIRAAM, "।"},
			{token.BANDA_MADHYAM_KOSHTHAK, "}"},
			{token.ANTYA, ""},
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
					"tests[%d] - literal wrong. expected=%q, got=%q",
					i,
					tt.expectedLiteral,
					tok.Literal,
				)
			}
		}
	})
}
