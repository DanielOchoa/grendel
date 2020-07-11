package ast

import "github.com/danielochoa/grendel/token"

// Node - holds tokens, a node is part of our AST (abstract syntax tree) our parser will be a
// recursive descent parser (Pratt Parser)
type Node interface {
	TokenLiteral() string
}

// Statement - represents a statement in programming speak.
type Statement interface {
	Node
	statementNode()
}

// Expression - represents an expression vs a statement in language speak.
type Expression interface {
	Node
	expressionNode()
}

// Program - root node of every AST our parser produces.
type Program struct {
	Statements []Statement
}

// TokenLiteral - satisfies node interface.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement - a statement node.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral - to satisfy node interface
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier - for example, in `let x = 5`, Identifier is the `x`. For simplicity
// Identifier satisfies the expression interface even though it's a statement (?).
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral - satisfy node interface
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }