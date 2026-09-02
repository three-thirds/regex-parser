package parser

// Package parser converts a stream of tokens into an abstract syntax
// tree (AST) representing the regular expression.

import ( // importing fmt, ast and lexer
	"fmt"

	"regex/ast"
	"regex/lexer"
)

// Parser holds the lexer and the current token during parsing.
type Parser struct {
	lexer   *lexer.Lexer
	current lexer.Token
}

// New creates a Parser for the provided input string.
func New(input string) *Parser {
	l := lexer.New(input)

	p := &Parser{
		lexer: l,
	}

	p.advance() // consume data and go to next one

	return p
}

// advance consumes the next token from the lexer and stores it
// as the parser's current token.
func (p *Parser) advance() {
	p.current = p.lexer.NextToken()
}

// Parse parses the input and returns the root AST node or an error.
func (p *Parser) Parse() (ast.Node, error) {
	node, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if p.current.Type != lexer.TokenEOF {
		return nil, fmt.Errorf(
			"unexpected token after expression: %v",
			p.current,
		)
	}

	return node, nil
}

func (p *Parser) parseExpression() (ast.Node, error) { // as of now expression -> primary so this is useless but in the future experession -> alternation -> ... -> primary, then ts will be useful
	return p.parseAlternation()
}

func (p *Parser) parseAlternation() (ast.Node, error) {
	left, err := p.parseConcatenation()
	if err != nil {
		return nil, err
	}

	for p.current.Type == lexer.TokenPipe {
		p.advance()

		right, err := p.parseConcatenation()
		if err != nil {
			return nil, err
		}

		left = &ast.Alternation{
			Left:  left,
			Right: right,
		}
	}

	return left, nil
}

func (p *Parser) parseConcatenation() (ast.Node, error) {
	left, err := p.parseRepetition()
	if err != nil {
		return nil, err
	}

	for canStartPrimary(p.current.Type) {
		right, err := p.parseRepetition()
		if err != nil {
			return nil, err
		}

		left = &ast.Concat{
			Left:  left,
			Right: right,
		}
	}
	return left, nil
}

func (p *Parser) parseRepetition() (ast.Node, error) {
	node, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}

	for {
		switch p.current.Type {
		case lexer.TokenStar:
			p.advance()

			node = &ast.Star{
				Expr: node,
			}

		case lexer.TokenPlus:
			p.advance()

			node = &ast.Plus{
				Expr: node,
			}
		case lexer.TokenQuestion:
			p.advance()

			node = &ast.Question{
				Expr: node,
			}

		default:
			return node, nil
		}
	}
}

func (p *Parser) parsePrimary() (ast.Node, error) {
	switch p.current.Type {
	case lexer.TokenLiteral:
		value := p.current.Value
		p.advance()

		return &ast.Literal{
			Value: value,
		}, nil

	case lexer.TokenDot:
		p.advance()

		return &ast.Dot{}, nil

	case lexer.TokenLParen:
		p.advance()

		node, err := p.parseExpression()

		if err != nil {
			return nil, err
		}

		if p.current.Type != lexer.TokenRParen {
			return nil, fmt.Errorf(
				"expected ')', got %v",
				p.current,
			)
		}

		p.advance()

		return node, nil

	default:
		return nil, fmt.Errorf(
			"unexpected token: %v",
			p.current,
		)
	}
}

func canStartPrimary(tokenType lexer.TokenType) bool {
	switch tokenType {
	case lexer.TokenLiteral,
		lexer.TokenDot,
		lexer.TokenLParen:
		return true
	default:
		return false
	}
}
