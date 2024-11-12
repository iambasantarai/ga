package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF        = "EOF"
	ILLEGAL    = "ILLEGAL"
	TERMINATOR = "TERMINATOR"

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
	MAIN   = "MAIN"
	RETURN = "RETURN"
	PRINT  = "PRINT"

	// Identifiers
	INT            = "INT"
	IDENT          = "IDENT"
	STRING_LITERAL = "STRING_LITERAL"
)

var keywords = map[string]TokenType{
	"पूर्णांक": INT,

	"मुख्य":    MAIN,
	"प्रदर्शन": PRINT,
	"फिर्ता":   RETURN,
	"।":        TERMINATOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
