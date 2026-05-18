package Symbol

import (
	"fmt"
	"ion/go/v2/AST"
	"ion/go/v2/Token"
)

type SymbolTable struct {
	parent      *SymbolTable
	variables   map[string]AST.Expression
	StackOffset int
	// deferStack []*AST.StatementDefer
}

func CreateSymbolTable(parent *SymbolTable) SymbolTable {
	return SymbolTable{
		parent:    parent,
		variables: make(map[string]AST.Expression),
		// deferStack: make([]*AST.StatementDefer, 0),
	}
}

/*
func (s *SymbolTable) ResolveDeferStack() {
	for i := len(s.deferStack) - 1; i >= 0; i-- {
		interpretNode(s.deferStack[i].DeferredNode, s)
	}

	s.deferStack = nil
}



func (s *SymbolTable) AddDeferStatement(deferStatement *AST.StatementDefer) {
	s.deferStack = append(s.deferStack, deferStatement)
}
*/

func (s *SymbolTable) has(key Token.Token) bool {
	current := s
	for current != nil {
		_, ok := current.variables[key.Lexeme]
		if ok {
			return true
		}
		current = current.parent
	}

	return false
}

func (s *SymbolTable) get(key Token.Token) AST.Expression {
	current := s
	for current != nil {
		value, ok := current.variables[key.Lexeme]
		if ok {
			return value
		}
		current = current.parent
	}

	panic(fmt.Sprintf("Line: %d | Undeclared Identifier: %s", key.Line, key.Lexeme))
	return nil
}

func (s *SymbolTable) set(key Token.Token, value AST.Expression) {
	current := s
	for current != nil {
		_, ok := current.variables[key.Lexeme]
		if ok {
			current.variables[key.Lexeme] = value
			return
		}
		current = current.parent
	}

	s.variables[key.Lexeme] = value
}
