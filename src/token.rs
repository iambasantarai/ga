#[allow(dead_code)]
#[derive(Debug)]
pub enum TokenKind {
    Eof, 
    Illegal,

    /* -- Punctuation/Delimiters -- */
    Comma,
    Colon,
    Period,
    Semicolon,
    OpenParen,
    CloseParen,
    OpenCurly,
    CloseCurly,
    
    /* -- Operators -- */
    Plus,               
    Minus,
    Star,
    Slash,
    Mod,
    Bang,
    Equal,

    /* -- Literals -- */
    Number,
    Character,
    String,
    Identifier,

    /* -- Keywords -- */
    Let,
    Function,
    Print,
    Return,
    If,
    Else,
    True,
    False
}

#[allow(dead_code)]
#[derive(Debug)]
pub struct Token {
    pub kind: TokenKind,
    pub literal: String
}

// Maps Devanagari identifiers to their corresponding keyword token kinds
#[allow(dead_code)]
pub struct DevanagariIdentToTokenMap;

impl DevanagariIdentToTokenMap {
    pub fn lookup(identifier: &str) -> Option<TokenKind> {
        match identifier {
            "मानौ" => Some(TokenKind::Let),
            "कार्य" => Some(TokenKind::Function),
            "छाप" => Some(TokenKind::Print),
            "यदि" => Some(TokenKind::If),
            "नभए" => Some(TokenKind::Else),
            "फिर्ता" => Some(TokenKind::Return),
            "सत्य" => Some(TokenKind::True),
            "असत्य" => Some(TokenKind::False),
            _ => None,
        }
    }
}
