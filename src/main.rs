mod token;
mod lexer;
mod utils;

use token::TokenKind;
use lexer::Lexer;

fn main() {
    let source = r#"कार्य मुख्य() {
            छाप("सोच्छौ के मेरो बारे?")।
            छाप(१२३)।
        }
        "#.to_string();
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
