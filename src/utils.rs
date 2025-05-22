pub fn is_devanagari_letter(ch: char) -> bool {
    return ('\u{0900}' <= ch && ch <= '\u{097F}') || ('\u{A8E0}' <= ch && ch <= '\u{A8FF}');
}

pub fn is_devanagari_vowel_sign(ch: char) -> bool {
    return '\u{093A}' <= ch && ch <= '\u{094F}';
}

pub fn is_digit(ch: char) -> bool {
    return '\u{0966}' <= ch && ch <= '\u{096F}';
}
