package ast

import (
	"github.com/fatihGunal/monkey-compiler/internal/token"
)

// Abstract syntax tree AST
// Example: let x = 5 + 3
//
// - Statement: the full line "let x = 5 + 3" — it performs an action.
// - Identifier: "x" — the name used to refer to a value.
// - Expression: "5 + 3" — computes and returns a value (8).

type Node interface {
	TokenLiteral() string
}

// does not produce values e.g. let x = 5 + 3
type Statement interface {
	Node
	statementNode()
}

// Expressions produce values e.g. add(5, 5)
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statement []Statement
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}
