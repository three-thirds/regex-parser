package parser

import ( // importing fmt, ast and lexer
	"fmt"

	"regex/ast"
	"regex/lexer"
)

type Parser struct {
	lexer   *lexer.Lexer // future tokens come from here
	current lexer.Token  // token we are looking at
}

func New(input string) *Parser {
	l := lexer.New(input)

	p := &Parser{
		lexer: l,
	}

	p.advance() // consume data and go to next one

	return p
}

func (p *Parser) advance() {
	p.current = p.lexer.NextToken()
}

func (p *Parser) Parse() (ast.Node, error) { // Parser parses and returns error if any
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
	return p.parseRepetition()
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
