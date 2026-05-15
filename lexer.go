package main

import (
	"fmt"
)

type Lexer struct {
	input string
	pos   int
	ch    byte
	line  int
	col   int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
		pos:   0,
		line:  1,
		col:   0,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.pos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.pos]
		l.pos++
	}
	if l.ch == '\n' {
		l.line++
		l.col = 0
	} else {
		l.col++
	}
}

func (l *Lexer) peekChar() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.pos - 1
	for isDigit(l.ch) || isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start : l.pos-1]
}

func (l *Lexer) readNumber() string {
	start := l.pos - 1
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start : l.pos-1]
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'z')
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}

func lookupKeyword(ident string) TokenType {
	switch ident {
	case "not":
		return T_NOT
	case "or":
		return T_OR
	case "and":
		return T_AND
	case "div":
		return T_DIVOP
	case "mod":
		return T_MOD
	default:
		return T_IDENT
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	startLine := l.line
	startCol := l.col
	var t Token

	switch l.ch {
	case 0:
		t = Token{Type: T_EOF, Value: "", Line: startLine, Col: startCol}
	case '{':
		t = Token{Type: T_LBRACE, Value: "{", Line: startLine, Col: startCol}
		l.readChar()
	case '}':
		t = Token{Type: T_RBRACE, Value: "}", Line: startLine, Col: startCol}
		l.readChar()
	case '(':
		t = Token{Type: T_LPAREN, Value: "(", Line: startLine, Col: startCol}
		l.readChar()
	case ')':
		t = Token{Type: T_RPAREN, Value: ")", Line: startLine, Col: startCol}
		l.readChar()
	case ';':
		t = Token{Type: T_SEMICOLON, Value: ";", Line: startLine, Col: startCol}
		l.readChar()
	case '+':
		t = Token{Type: T_PLUS, Value: "+", Line: startLine, Col: startCol}
		l.readChar()
	case '-':
		t = Token{Type: T_MINUS, Value: "-", Line: startLine, Col: startCol}
		l.readChar()
	case '*':
		t = Token{Type: T_MUL, Value: "*", Line: startLine, Col: startCol}
		l.readChar()
	case '/':
		t = Token{Type: T_DIV, Value: "/", Line: startLine, Col: startCol}
		l.readChar()
	case '=':
		t = Token{Type: T_EQ, Value: "=", Line: startLine, Col: startCol}
		l.readChar()
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			t = Token{Type: T_LE, Value: "<=", Line: startLine, Col: startCol}
			l.readChar()
		} else if l.peekChar() == '>' {
			l.readChar()
			t = Token{Type: T_NE, Value: "<>", Line: startLine, Col: startCol}
			l.readChar()
		} else {
			t = Token{Type: T_LT, Value: "<", Line: startLine, Col: startCol}
			l.readChar()
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			t = Token{Type: T_GE, Value: ">=", Line: startLine, Col: startCol}
			l.readChar()
		} else {
			t = Token{Type: T_GT, Value: ">", Line: startLine, Col: startCol}
			l.readChar()
		}
	case ':':
		if l.peekChar() == '=' {
			l.readChar()
			t = Token{Type: T_ASSIGN, Value: ":=", Line: startLine, Col: startCol}
			l.readChar()
		} else {
			t = Token{Type: T_ILLEGAL, Value: ":", Line: startLine, Col: startCol}
			l.readChar()
		}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tokType := lookupKeyword(ident)
			t = Token{Type: tokType, Value: ident, Line: startLine, Col: startCol}
		} else if isDigit(l.ch) {
			number := l.readNumber()
			t = Token{Type: T_NUMBER, Value: number, Line: startLine, Col: startCol}
		} else {
			t = Token{Type: T_ILLEGAL, Value: string(l.ch), Line: startLine, Col: startCol}
			l.readChar()
		}
	}
	return t
}

func (l *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token
	for {
		t := l.NextToken()
		tokens = append(tokens, t)
		if t.Type == T_EOF {
			break
		}
		if t.Type == T_ILLEGAL {
			return tokens, fmt.Errorf("illegal character '%s' at %d:%d", t.Value, t.Line, t.Col)
		}
	}
	return tokens, nil
}
