package api

import (
	"errors"
	"main/core/parser"
)

type SCmdMethods interface {
	// New constructs a new SCmd
	New(name string, description string)
	Register()
	syntaxify() parser.Rule
	get()
}

type SCmd struct {
	// Name which is used to call this command
	Name string
	// Description used to generate the documentation
	Description string
	// List of pointers to other SCmd structs
	SCmd []*SCmd
	// List of pointers to Cmd structs
	Cmd []*Cmd
}

// syntaxify a SCmd and returns a parser.Rule
//	used to build the SRT of the CLI
func (sc *SCmd) syntaxify() parser.Rule {
	var r parser.Rule
	m := make(parser.SRT)

	// RuleSpec has only Srt specified to true
	r.Spec = parser.RuleSpec{Srt: true}
	// syntaxify sc.Cmd
	for _, c := range sc.Cmd {
		m[c.Name] = c.syntaxify()
	}
	// syntaxify sc.SCmd
	for _, ssc := range sc.SCmd {
		m[ssc.Name] = ssc.syntaxify()
	}
	r.Srt = m
	return r
}

func (sc *SCmd) get(ast *parser.AST) (Func, *parser.AST, error) {
	for _, cmd := range sc.Cmd {
		if cmd.Name == ast.Name {
			return cmd.F, ast, nil
		}
	}
	for _, sc := range sc.SCmd {
		if sc.Name == ast.Name {
			return sc.get(ast.Ast)
		}
	}
	return nil, ast, errors.New("error while trying to execute the command")
}
