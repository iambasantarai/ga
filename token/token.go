package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special tokens
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Statement terminator
	DELIMITER = "DELIMITER"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	OPAREN    = "("
	CPAREN    = ")"
	OCURLY    = "{"
	CCURLY    = "}"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	MODULO   = "%"
	BANG     = "!"

	// Comparison Operators
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Logical Operators
	AND = "र"
	OR  = "वा"

	// Keywords
	LET    = "LET"
	FUNC   = "FUNC"
	RETURN = "RETURN"
	IF     = "IF"
	ELSE   = "ELSE"
	FOR    = "FOR"
	WHILE  = "WHILE"
	DO     = "DO"
	BREAK  = "BREAK"
	PRINT  = "PRINT"

	// Data Types
	INT   = "INT"
	FLOAT = "FLOAT"
	CHAR  = "CHAR"
	BOOL  = "BOOL"

	// Boolean values
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Identifiers and literals
	IDENT          = "IDENT"
	INT_LITERAL    = "INT_LITERAL"
	FLOAT_LITERAL  = "FLOAT_LITERAL"
	STRING_LITERAL = "STRING_LITERAL"
)

var keywords = map[string]TokenType{
	"मानौं":     LET,
	"कार्य":     FUNC,
	"फिर्ता":    RETURN,
	"यदि":       IF,
	"अथवा":      ELSE,
	"चक्र":      FOR,
	"जबसम्म":    WHILE,
	"गर":        DO,
	"रोक":       BREAK,
	"छाप":       PRINT,
	"संख्या":    INT,
	"दशमलव":     FLOAT,
	"अक्षर":     CHAR,
	"सत्यअसत्य": BOOL,
	"सत्य":      TRUE,
	"असत्य":     FALSE,
	"र":         AND,
	"वा":        OR,
	"।":         DELIMITER,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
