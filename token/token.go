package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ANTYA   = "ANTYA"   // EOF
	ABAIDHA = "ABAIDHA" // Illegal

	/* -- Punctuation/Delimiters -- */
	BINDU                    = "."            // Dot
	ALPA_BIRAAM              = ","            // Comma
	ARDHA_BIRAAM             = ";"            // Semicolon
	KHULLAA_SAANO_KOSHTHAK   = "("            // Opening parenthesis
	BANDA_SAANO_KOSHTHAK     = ")"            // Closing parenthesis
	KHULLAA_MADHYAM_KOSHTHAK = "{"            // Opening curly bracket
	BANDA_MADHYAM_KOSHTHAK   = "}"            // Closing curly bracket
	PURNA_BIRAAM             = "PURNA_BIRAAM" // Statement terminator

	/* -- Operators -- */
	BARAABAR   = "=" // Assignment operator
	JODA       = "+" // Plus
	GHATAU     = "-" // Minus
	GUNAN      = "*" // Multiplication
	BHAAG      = "/" // Division
	SHESHAFALA = "%" // Modulo
	BIPARIT    = "!" // Negation

	/* -- Keywords -- */
	MAANAU  = "MAANAU"  // Let
	KAARYA  = "KAARYA"  // Function
	FIRTAA  = "FIRTAA"  // Return
	YADI    = "YADI"    // If
	NABHAYE = "NABHAYE" // Else
	CHAKRA  = "CHAKRA"  // For
	JABAKI  = "JABAKI"  // While
	GARA    = "GARA"    // Do
	ROKA    = "ROKA"    // Break
	CHAAPA  = "PRINT"   // Print
	SATYA   = "SATYA"   // True
	ASATYA  = "ASATYA"  // False

	/* -- Literals -- */
	IDENT        = "IDENT"
	ANK          = "ANK"          // Integer
	DASAMALAB    = "DASAMALAB"    // Float
	AKKSHYARA    = "AKKSHYARA"    // Character
	SATYA_ASATYA = "SATYA_ASATYA" // Boolean
	BAAKYA       = "BAAKYA"       // String
)

var keywords = map[string]TokenType{
	/* -- Keywords -- */
	"मानौ":   MAANAU,
	"कार्य":  KAARYA,
	"फिर्ता": FIRTAA,
	"यदि":    YADI,
	"नभए":    NABHAYE,
	"चक्र":   CHAKRA,
	"जबकि":   JABAKI,
	"गर":     GARA,
	"रोक":    ROKA,
	"छाप":    CHAAPA,
	"सत्य":   SATYA,
	"असत्य":  ASATYA,

	/* -- Literals -- */
	"अंक":   ANK,
	"दशमलव": DASAMALAB,
	"अक्षर": AKKSHYARA,
	"वाक्य": BAAKYA,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
