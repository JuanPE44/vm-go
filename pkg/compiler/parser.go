package compiler

import (
	"strconv"
	"vm-go/pkg/ast"
)

type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseExpression() ast.Node {
	return p.parseBinaryOp(0)
}

var precedences = map[TokenType]int{
	TokenPlus:  10,
	TokenMinus: 10,
	TokenStar:  20,
	TokenSlash: 20,
}

func (p *Parser) parseBinaryOp(precedence int) ast.Node {
	left := p.parsePrimary()

	for p.curToken.Type != TokenEOF && precedence < precedences[p.curToken.Type] {
		op := p.curToken.Literal
		level := precedences[p.curToken.Type]

		p.nextToken()
		right := p.parseBinaryOp(level)

		left = &ast.BinaryOpNode{
			Left:     left,
			Operator: op,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parsePrimary() ast.Node {
	switch p.curToken.Type {
	case TokenNumber:
		val, _ := strconv.Atoi(p.curToken.Literal)
		node := &ast.NumberNode{Value: val}
		p.nextToken()
		return node

	case TokenLParen:
		p.nextToken()
		node := p.parseBinaryOp(0)
		if p.curToken.Type != TokenRParen {
			panic("Se esperaba ')'")
		}
		p.nextToken()
		return node
	}
	return nil
}
