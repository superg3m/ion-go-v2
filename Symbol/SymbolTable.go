package Symbol

import (
	"fmt"
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

type Symbol struct {
	Tok      Token.Token
	DeclType TS.Type
}

func CreateSymbol(token Token.Token, declType TS.Type) Symbol {
	return Symbol{
		Tok:      token,
		DeclType: declType,
	}
}

type SymbolTable struct {
	parent      *SymbolTable
	symbols     map[string]Symbol
	offsets     map[Symbol]int
	StackOffset int
	// deferStack []*AST.StatementDefer
}

func CreateSymbolTable(parent *SymbolTable) SymbolTable {
	return SymbolTable{
		parent:  parent,
		symbols: make(map[string]Symbol),
		offsets: make(map[Symbol]int),
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

func (s *SymbolTable) Has(key Token.Token) bool {
	current := s
	for current != nil {
		_, ok := current.symbols[key.Lexeme]
		if ok {
			return true
		}
		current = current.parent
	}

	return false
}

func (s *SymbolTable) GetSymbol(key Token.Token) Symbol {
	current := s
	for current != nil {
		value, ok := current.symbols[key.Lexeme]
		if ok {
			return value
		}
		current = current.parent
	}

	panic(fmt.Sprintf("Line: %d | Undeclared Identifier: %s", key.Line, key.Lexeme))
	return Symbol{}
}

func (s *SymbolTable) GetOffset(key Token.Token) int {
	current := s
	for current != nil {
		value, ok := current.symbols[key.Lexeme]
		if ok {
			return s.offsets[value]
		}
		current = current.parent
	}

	panic(fmt.Sprintf("Line: %d | Undeclared Identifier: %s", key.Line, key.Lexeme))
	return 0
}

func (s *SymbolTable) Set(key Token.Token, symbol Symbol, offset int) {
	current := s
	for current != nil {
		_, ok := current.symbols[key.Lexeme]
		if ok {
			current.symbols[key.Lexeme] = symbol
			current.offsets[symbol] = offset
			return
		}
		current = current.parent
	}

	s.symbols[key.Lexeme] = symbol
	s.offsets[symbol] = offset
}
