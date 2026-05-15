package main

import (
	"fmt"
)

type TokenType int

const (
	T_EOF TokenType = iota
	T_IDENT
	T_NUMBER
	T_ASSIGN
	T_SEMICOLON
	T_LBRACE
	T_RBRACE
	T_LPAREN
	T_RPAREN
	T_PLUS
	T_MINUS
	T_MUL
	T_DIV
	T_EQ
	T_NE
	T_LT
	T_LE
	T_GT
	T_GE
	T_NOT
	T_OR
	T_AND
	T_DIVOP
	T_MOD
	T_ILLEGAL
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

func (t Token) String() string {
	return fmt.Sprintf("(%d,%d) %v %q", t.Line, t.Col, t.Type, t.Value)
}
