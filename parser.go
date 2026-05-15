package main

import (
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
}

func (p *Parser) NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: T_EOF, Value: "", Line: -1, Col: -1}
	}
	return p.tokens[p.pos]
}

func (p *Parser) match(expexted TokenType) error {
	if cur := p.current(); cur.Type == expexted {
		p.pos++
		return nil
	} else {
		return fmt.Errorf("line %d, col %d, expected %v, found %v", cur.Line, cur.Col, expexted, cur.Type)
	}
}

func (p *Parser) isAtEnd() bool {
	return p.pos >= len(p.tokens) || p.current().Type == T_EOF
}

func (p *Parser) parseProgram() error {
	return p.parseBlock()
}

func (p *Parser) parseBlock() error {
	if err := p.match(T_LBRACE); err != nil {
		return err
	}
	if err := p.parseStatementList(); err != nil {
		return err
	}
	if err := p.match(T_RBRACE); err != nil {
		return err
	}
	return nil
}

func (p *Parser) parseStatementList() error {
	if err := p.parseStatement(); err != nil {
		return err
	}
	return p.parseTail()
}

func (p *Parser) parseTail() error {
	if p.current().Type == T_SEMICOLON {
		p.match(T_SEMICOLON)
		if err := p.parseStatement(); err != nil {
			return err
		}
		return p.parseTail()
	} else {
		return nil
	}
}

func (p *Parser) parseStatement() error {
	switch p.current().Type {
	case T_IDENT:
		p.match(T_IDENT)
		if err := p.match(T_ASSIGN); err != nil {
			return err
		}
		return p.parseExpression()
	case T_LBRACE:
		return p.parseBlock()
	default:
		return fmt.Errorf("expected ident or {, found %v, line %d, col %d", p.current().Type, p.current().Line, p.current().Col)
	}
}

func (p *Parser) parseExpression() error {
	if err := p.parseSimpleExpression(); err != nil {
		return err
	}
	if isRelOp(p.current().Type) {
		if err := p.parseRelOp(); err != nil {
			return err
		}
		return p.parseSimpleExpression()
	}
	return nil
}

func (p *Parser) parseSimpleExpression() error {
	if p.current().Type == T_PLUS || p.current().Type == T_MINUS {
		if err := p.parseSign(); err != nil {
			return err
		}
	}
	if err := p.parseTerm(); err != nil {
		return err
	}
	for isAddOp(p.current().Type) {
		if err := p.parseAddOp(); err != nil {
			return err
		}
		if err := p.parseTerm(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) parseTerm() error {
	if err := p.parseFactor(); err != nil {
		return err
	}
	for isMulOp(p.current().Type) {
		if err := p.parseMulOp(); err != nil {
			return err
		}
		if err := p.parseFactor(); err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) parseFactor() error {
	switch p.current().Type {
	case T_IDENT:
		p.match(T_IDENT)
	case T_NUMBER:
		p.match(T_NUMBER)
	case T_LPAREN:
		p.match(T_LPAREN)
		if err := p.parseSimpleExpression(); err != nil {
			return err
		}
		p.match(T_RPAREN)
	case T_NOT:
		if err := p.match(T_NOT); err != nil {
			return err
		}
	default:
		return fmt.Errorf("wrong factor %v, line %d, col %d", p.current().Type, p.current().Line, p.current().Col)
	}
	return nil
}

func (p *Parser) parseSign() error {
	if p.current().Type == T_MINUS || p.current().Type == T_PLUS {
		p.match(p.current().Type)
		return nil
	}
	return fmt.Errorf("expected sign")
}

func (p *Parser) parseAddOp() error {
	if isAddOp(p.current().Type) {
		p.match(p.current().Type)
		return nil
	}
	return fmt.Errorf("expected add operation")
}

func (p *Parser) parseMulOp() error {
	if isMulOp(p.current().Type) {
		p.match(p.current().Type)
		return nil
	}
	return fmt.Errorf("expected multiplication operation")
}

func (p *Parser) parseRelOp() error {
	if isRelOp(p.current().Type) {
		p.match(p.current().Type)
		return nil
	}
	return fmt.Errorf("expected relation operation")
}

func isRelOp(t TokenType) bool {
	return t == T_EQ || t == T_NE || t == T_LT || t == T_LE || t == T_GT || t == T_GE
}

func isAddOp(t TokenType) bool {
	return t == T_PLUS || t == T_MINUS || t == T_OR
}

func isMulOp(t TokenType) bool {
	return t == T_MUL || t == T_DIV || t == T_DIVOP || t == T_MOD || t == T_AND
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}
