package lexer

import (
	"fmt"
	"unicode/utf8"
)

type Token struct {
	Type   TokenType
	Line   int
	Col    int
	Lexeme string
}

type TokenType int

const (
	INTEGER TokenType = iota

	// Operators
	PLUS    // +
	MINUS   // -
	BINTANG // *
	GARING  // /

	KURUNG_BUKA
	KURUNG_TUTUP

	IDENT

	EOF
)

func (t Token) String() string {
	switch t.Type {
	case INTEGER:
		return t.Lexeme
	case KURUNG_BUKA:
		return "("
	case KURUNG_TUTUP:
		return ")"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case BINTANG:
		return "*"
	case GARING:
		return "/"
	case IDENT:
		return t.Lexeme
	case EOF:
		return "EOF"
	}

	return ""
}

type Lexer struct {
	src    string
	tokens []Token

	// Lexer state
	start   int
	current int
	line    int
	col     int
}

func New(src string) *Lexer {
	return &Lexer{
		src: src,

		line: 1,
		col:  1,
	}
}

func (l *Lexer) Tokens() []Token {
	for !l.shouldEnd() {
		l.start = l.current
		l.lex()
	}

	l.tokens = append(l.tokens, Token{Type: EOF, Lexeme: "", Line: l.line})

	return l.tokens
}

func (l *Lexer) lex() {
	c := l.next()

	switch c {
	case ' ', '\t', '\r':
		// lewati whitespace
	case '(':
		l.tokens = append(l.tokens, Token{Type: KURUNG_BUKA, Lexeme: "", Line: l.line, Col: l.col})
	case ')':
		l.tokens = append(l.tokens, Token{Type: KURUNG_TUTUP, Lexeme: "", Line: l.line, Col: l.col})
	case '+':
		l.tokens = append(l.tokens, Token{Type: PLUS, Lexeme: "", Line: l.line, Col: l.col})
	case '*':
		l.tokens = append(l.tokens, Token{Type: BINTANG, Lexeme: "", Line: l.line, Col: l.col})
	case '-':
		l.tokens = append(l.tokens, Token{Type: MINUS, Lexeme: "", Line: l.line, Col: l.col})
	case '/':
		l.tokens = append(l.tokens, Token{Type: GARING, Lexeme: "", Line: l.line, Col: l.col})
	case '\n':
		l.line++
	default:
		if isDigit(c) {
			l.num()
		} else if isAlph(c) {
			l.ident()
		} else {
			fmt.Printf("\n")
		}
	}
}

func (l *Lexer) peek() rune {
	if l.shouldEnd() {
		return 0
	}
	r, _ := utf8.DecodeRune([]byte(l.src[l.current:]))
	return r
}

func (l *Lexer) num() {
	for isDigit(l.peek()) {
		l.forward()
	}
	l.tokens = append(l.tokens, Token{
		Type:   INTEGER,
		Lexeme: string(l.src[l.start:l.current]),
		Line:   l.line,
	})
}

func (l *Lexer) ident() {
	for isAlphaNum(l.peek()) {
		l.forward()
	}
	l.tokens = append(l.tokens, Token{
		Type:   IDENT,
		Lexeme: string(l.src[l.start:l.current]),
		Line:   l.line,
	})
}

func (l *Lexer) shouldEnd() bool {
	return l.current >= len(l.src)
}

func (l *Lexer) next() rune {
	r, w := utf8.DecodeRune([]byte(l.src[l.current:]))
	l.current += w
	return r
}

func (l *Lexer) forward() {
	l.current++
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isAlph(c rune) bool {
	return 'a' <= c && c <= 'z'
}

func isAlphaNum(c rune) bool {
	return isAlph(c) || isDigit(c)
}
