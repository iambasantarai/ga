use crate::token::{Token, TokenKind, DevanagariIdentToTokenMap};
use crate::utils;

pub struct Lexer {
    input: Vec<char>,
    position: usize,
    read_position: usize,
    character: char
}

#[allow(dead_code)]
impl Lexer {
    pub fn new(input: String) -> Self {
        let input_chars: Vec<char> = input.chars().collect();
        let mut lexer = Self {
            input: input_chars,
            position: 0,
            read_position: 0,
            character: '\0'
        };

        lexer.read_char();
        return lexer;
    }

    pub fn next_token(&mut self) -> Token {
        self.skip_whitespace();

        let token = match self.character {
            '.' => self.create_token(TokenKind::Period),
            ',' => self.create_token(TokenKind::Comma),
            ':' => self.create_token(TokenKind::Colon),
            ';' => self.create_token(TokenKind::Semicolon),
            '(' => self.create_token(TokenKind::OpenParen),
            ')' => self.create_token(TokenKind::CloseParen),
            '{' => self.create_token(TokenKind::OpenCurly),
            '}' => self.create_token(TokenKind::CloseCurly),
            '+' => self.create_token(TokenKind::Plus),
            '-' => self.create_token(TokenKind::Minus),
            '*' => self.create_token(TokenKind::Star),
            '/' => self.create_token(TokenKind::Slash),
            '%' => self.create_token(TokenKind::Mod),
            '!' => self.create_token(TokenKind::Bang),
            '=' => self.create_token(TokenKind::Equal),
            '\0' => Token {
                kind: TokenKind::Eof,
                literal: "".to_string()
            },
            _ =>  {
                if utils::is_devanagari_letter(self.character) {
                    let identifier = self.read_identifier();

                    // Look up the identifier in the Devanagari keyword token map
                    let kind = DevanagariIdentToTokenMap::lookup(&identifier)
                        .unwrap_or(TokenKind::Identifier);

                    return Token {
                        kind,
                        literal: identifier
                    };
                } else if utils::is_digit(self.character) {
                    return Token {
                        kind: TokenKind::Number,
                        literal: self.read_number()
                    };

                } else {
                    self.create_token(TokenKind::Illegal)
                }
            }
        };

        self.read_char();
        return token;
    }

    fn create_token(&self, kind: TokenKind) -> Token {
        return Token{
            kind,
            literal: self.character.to_string()
        }
    }

    fn read_char(&mut self) {
        if self.read_position >= self.input.len() {
            self.character = '\0';
        } else {
            self.character = self.input[self.read_position];
        }
        self.position = self.read_position;
        self.read_position += 1;
    }

    fn peek_char(&self) -> char {
        if self.read_position >= self.input.len() {
            return '\0';
        } else {
            return self.input[self.read_position];
        }
    }

    fn skip_whitespace(&mut self) {
        while self.character.is_whitespace() {
            self.read_char();
        }
    }

    fn read_identifier(&mut self) -> String {
        let position = self.position;
        while utils::is_devanagari_letter(self.character) || utils::is_devanagari_vowel_sign(self.character) {
            self.read_char();
        }

        return self.input[position..self.position].iter().collect();
    }

    fn read_number(&mut self) -> String {
        let position = self.position;
        while utils::is_digit(self.character) {
            self.read_char();
        }

        return self.input[position..self.position].iter().collect();
    }
}
