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

	OPAREN = "("
	CPAREN = ")"
	OCURLY = "{"
	CCURLY = "}"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"

	// Keywords
	VOID   = "VOID"
	MAIN   = "MAIN"
	RETURN = "RETURN"
	PRINT  = "PRINT"
	IF     = "IF"
	ELSE   = "ELSE"
	BREAK  = "BREAK"

	// Types
	INT   = "INT"
	FLOAT = "FLOAT"
	CHAR  = "CHAR"

	// Identifiers
	IDENT = "IDENT"

	STRING_LITERAL = "STRING_LITERAL"
)

var keywords = map[string]TokenType{
	"पूर्णांक": INT,
	"दशमलव":    FLOAT,
	"अक्षर":    CHAR,

	"मुख्य":    MAIN,
	"शून्य":    VOID,
	"प्रदर्शन": PRINT,
	"फिर्ता":   RETURN,
	"यदि":      IF,
	"अन्यथा":   ELSE,
	"रोक":      BREAK,

	"।": TERMINATOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
