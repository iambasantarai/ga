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
            '"' => {
                let literal = self.read_string();
                return Token {
                    kind: TokenKind::String,
                    literal
                }
            }
            '\0' => Token {
                kind: TokenKind::Eof,
                literal: "".to_string()
            },
            _ =>  {
                if utils::is_devanagari_letter(self.character) {
                    if utils::is_digit(self.character) {
                        return Token {
                            kind: TokenKind::Number,
                            literal: self.read_number()
                        };
                    } else {

                        let identifier = self.read_identifier();

                        // Look up the identifier in the Devanagari keyword token map
                        let kind = DevanagariIdentToTokenMap::lookup(&identifier)
                            .unwrap_or(TokenKind::Identifier);

                        return Token {
                            kind,
                            literal: identifier
                        };
                    }
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

    fn read_string(&mut self)-> String {
        self.read_char();
        let mut literal = String::new();

        while self.character != '\"' {
            literal.push(self.character);
            self.read_char();
        }
        self.read_char();

        return  literal
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_a_portion_of_source_code() {
        let source = r#"कार्य काम() {
                सन्देश = "सोच्छौ के मेरो बारे?"।
                छाप(सन्देश)
                छाप(१२३)
            }
            काम()
        "#.to_string();

        let mut lexer = Lexer::new(source);

        let expected_tokens = vec![
            Token { kind: TokenKind::Function, literal: "कार्य".to_string() },
            Token { kind: TokenKind::Identifier, literal: "काम".to_string() },
            Token { kind: TokenKind::OpenParen, literal: "(".to_string() },
            Token { kind: TokenKind::CloseParen, literal: ")".to_string() },
            Token { kind: TokenKind::OpenCurly, literal: "{".to_string() },
            Token { kind: TokenKind::Identifier, literal: "सन्देश".to_string() },
            Token { kind: TokenKind::Equal, literal: "=".to_string() },
            Token { kind: TokenKind::String, literal: "सोच्छौ के मेरो बारे?".to_string() },
            Token { kind: TokenKind::Terminator, literal: "।".to_string() },
            Token { kind: TokenKind::Print, literal: "छाप".to_string() },
            Token { kind: TokenKind::OpenParen, literal: "(".to_string() },
            Token { kind: TokenKind::Identifier, literal: "सन्देश".to_string() },
            Token { kind: TokenKind::CloseParen, literal: ")".to_string() },
            Token { kind: TokenKind::Print, literal: "छाप".to_string() },
            Token { kind: TokenKind::OpenParen, literal: "(".to_string() },
            Token { kind: TokenKind::Number, literal: "१२३".to_string() },
            Token { kind: TokenKind::CloseParen, literal: ")".to_string() },
            Token { kind: TokenKind::CloseCurly, literal: "}".to_string() },
            Token { kind: TokenKind::Identifier, literal: "काम".to_string() },
            Token { kind: TokenKind::OpenParen, literal: "(".to_string() },
            Token { kind: TokenKind::CloseParen, literal: ")".to_string() },
            Token { kind: TokenKind::Eof, literal: "".to_string() },
        ];

        let mut index = 0;
        loop {
            let token = lexer.next_token();
            let expected_token = &expected_tokens[index];

            assert_eq!(token.kind, expected_token.kind, "Kind mismatch at token index {}", index);
            assert_eq!(token.literal, expected_token.literal, "Literal mismatch at token index {}", index);

            // If the token is Eof, stop the loop
            if token.kind == TokenKind::Eof {
                break;
            }

            index += 1;
        }
    }
}
