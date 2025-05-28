use std::{env, fs, path, process};
mod token;
mod lexer;
mod utils;

use token::TokenKind;
use lexer::Lexer;
use path::Path;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Missing source file path argument.");
        process::exit(1);
    }

    let file_path = &args[1];

    let file_ext = Path::new(file_path).extension().and_then(|ext| ext.to_str());
    if file_ext != Some("ga") {
        eprintln!("Invalid file extension.\nPlease provide a file with the '.ga' extension.");
        process::exit(1);
    }

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
