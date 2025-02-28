package lexer

import (
	"ga/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character, _ = utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.position = l.readPosition
		l.readPosition += utf8.RuneLen(l.character)
	}
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()

	var tok token.Token

	switch l.character {
	case '=':
		tok = newToken(token.BARAABAR, l.character)
	case '+':
		tok = newToken(token.JODA, l.character)
	case '-':
		tok = newToken(token.GHATAU, l.character)
	case '*':
		tok = newToken(token.GUNAN, l.character)
	case '/':
		tok = newToken(token.BHAAG, l.character)
	case '%':
		tok = newToken(token.SHESHAFALA, l.character)
	case '!':
		tok = newToken(token.BIPARIT, l.character)
	case ',':
		tok = newToken(token.ALPA_BIRAAM, l.character)
	case ';':
		tok = newToken(token.ARDHA_BIRAAM, l.character)
	case '(':
		tok = newToken(token.KHULLAA_SAANO_KOSHTHAK, l.character)
	case ')':
		tok = newToken(token.BANDA_SAANO_KOSHTHAK, l.character)
	case '{':
		tok = newToken(token.KHULLAA_MADHYAM_KOSHTHAK, l.character)
	case '}':
		tok = newToken(token.BANDA_MADHYAM_KOSHTHAK, l.character)
	case '।':
		tok.Literal = "।"
		tok.Type = token.PURNA_BIRAAM
	case '"':
		tok.Type = token.BAAKYA
		tok.Literal = l.readStringLiteral()
	case 0:
		tok.Literal = ""
		tok.Type = token.ANTYA
	default:
		if isDevanagariLetter(l.character) {
			tok.Literal = l.readKeywordOrIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		} else if isDigit(l.character) {
			tok.Type = token.ANK
			tok.Literal = l.readNumericLiteral()

			return tok
		} else {
			tok = newToken(token.ABAIDHA, l.character)
		}
	}

	l.readChar()

	return tok
}

func newToken(tokenType token.TokenType, character rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.character) {
		l.readChar()
	}
}

func isDevanagariLetter(character rune) bool {
	/*
		Devanagari characters live in the U+0900 to U+097F block
		return character >= '\u0900' && character <= '\u097F'
	*/
	return unicode.Is(unicode.Devanagari, character)
}

func (l *Lexer) readKeywordOrIdentifier() string {
	position := l.position
	for isDevanagariLetter(l.character) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readStringLiteral() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.character == '"' || l.character == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func isDigit(character rune) bool {
	return unicode.IsDigit(character)
}

func (l *Lexer) readNumericLiteral() string {
	position := l.position
	for isDigit(l.character) {
		l.readChar()
	}

	return l.input[position:l.position]
}
