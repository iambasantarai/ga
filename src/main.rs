use std::{fs, path};
mod token;
mod lexer;
mod utils;

use token::TokenKind;
use lexer::Lexer;
use path::Path;

fn main() {
    let file_path = Path::new("examples/init.ga");
    let source = fs::read_to_string(file_path).expect("Can't read this file.");

    let mut lexer = Lexer::new(source);

    loop {
        let token = lexer.next_token();

        match token.kind {
            TokenKind::Eof  => {
                    println!("{:?}", token);
                    break;
            },
            _ =>  {
                println!("{:?}", token)
            },
        }
    }
}
