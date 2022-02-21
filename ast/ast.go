package ast

import "github.com/DanielOchoa/grendel/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// Program starts with a root statement node.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Satisfies Statement and Node interface.
// <LetStatement> <Name> = <Value>
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

// Purely to differentiate and get compile time errors if we wrongly assign to expression node.
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier. E.g. 'let <identifier> = <expression>.
// Satisfies Node and Expression interfaces.
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // variable name as string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
