package api

import (
	"errors"
	"log"
	"main/core"
	"main/core/parser"
)

type CLIMethods interface {
	Run()
	syntaxify() parser.SRT
	funcify() core.F
	get() Func
}

type CLI struct {
	Name        string
	Description string
	Version     string
	SCmd        []*SCmd
	Cmd         []*Cmd
}

func (c *CLI) Run(s string) {
	core.Interpret(s, c.funcify(), c.syntaxify())
}

// syntaxify from CLI object
func (c *CLI) syntaxify() *parser.SRT {
	m := make(parser.SRT)
	// build parser.Rule for each CLI.SCmd
	for _, sc := range c.SCmd {
		m[sc.Name] = sc.syntaxify()
	}
	// build parser.Rule for each CLI.Cmd
	for _, cmd := range c.Cmd {
		m[cmd.Name] = cmd.syntaxify()
	}
	return &m
}

// funcify
func (c *CLI) funcify() core.F {
	return func(ast *parser.AST) {
		f, ast, err := c.get(ast)
		if err != nil {
			log.Fatalf("%s", err)
		}
		a, kw := convertAst(ast)
		f(a, kw)
	}
}

// get
func (c *CLI) get(ast *parser.AST) (Func, *parser.AST, error) {
	for _, cmd := range c.Cmd {
		if cmd.Name == ast.Name {
			return cmd.F, ast, nil
		}
	}
	for _, sc := range c.SCmd {
		if sc.Name == ast.Name {
			return sc.get(ast.Ast)
		}
	}
	return nil, ast, errors.New("error while trying to execute the command")
}

// convertAst to args/kwargs
func convertAst(ast *parser.AST) (Args, KWArgs) {
	m := make(map[string][]string)
	for _, fl := range ast.Opt.Short {
		m[fl.Name] = fl.Pos
	}
	for _, fl := range ast.Opt.Long {
		m[fl.Name] = fl.Pos
	}
	return ast.Pos, m
}
