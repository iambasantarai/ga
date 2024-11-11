package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	// Keywords
	MAIN     = "MAIN"
	FUNCTION = "FUNCTIOn"
	RETURN   = "RETURN"

	// Identifiers
	IDENT = "IDENT"
	INT   = "INT"
)

var keywords = map[string]TokenType{
	"पूर्णांक": INT,

	"मुख्य":  MAIN,
	"कार्य":  FUNCTION,
	"फिर्ता": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
