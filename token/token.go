package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
	DELIMITER = "DELIMITER"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	OPAREN = "("
	CPAREN = ")"
	OCURLY = "{"
	CCURLY = "}"

	// Operators
	ASSIGN   = "="
	ASTERISK = "*"
	BANG     = "!"
	MINUS    = "-"
	PLUS     = "+"
	SLASH    = "/"

	// Keywords
	BREAK  = "BREAK"
	ELSE   = "ELSE"
	FUNC   = "FUNC"
	IF     = "IF"
	LET    = "LET"
	PRINT  = "PRINT"
	RETURN = "RETURN"

	// Types
	BOOL = "BOOL"
	CHAR = "CHAR"
	INT  = "INT"

	// Loops
	DO    = "DO"
	FOR   = "FOR"
	WHILE = "WHILE"

	// Identifiers
	IDENT = "IDENT"

	STRING_LITERAL = "STRING_LITERAL"
)

var keywords = map[string]TokenType{
	"मानौं": LET,

	"संख्या":    INT,
	"अक्षर":     CHAR,
	"सत्यअसत्य": BOOL,

	"छाप": PRINT,

	"कार्य":  FUNC,
	"फिर्ता": RETURN,

	"गर":     DO,
	"जबसम्म": WHILE,
	"चक्र":   FOR,

	"यदि":  IF,
	"अथवा": ELSE,

	"रोक": BREAK,

	"।": DELIMITER,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
