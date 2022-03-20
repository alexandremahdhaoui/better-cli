package api

import (
	"main/core/parser"
)

type CmdMethods interface {
	New()
	Register()
	syntaxify() parser.Rule
	get()
}

type Args []string
type KWArgs map[string][]string
type Func func(a Args, kw KWArgs)

type Cmd struct {
	// Name which is used to call this command
	Name string
	// Description is the Description used to generate the documentation
	Description string
	// List of Expected Argument specifications
	PArgs ApiPSpec
	// List of Flag specifications
	Flags ApiFSpec
	// The function the cmd will execute, with provided PArgs and Flags
	F Func
}

// syntaxify a SCmd and returns a parser.Rule
//	used to build the SRT of the CLI
func (c *Cmd) syntaxify() parser.Rule {
	var r parser.Rule
	// spec
	ps := c.PArgs.syntaxify()
	r.Spec.Pos = ps.Pos
	r.Spec.PosR = ps.PosR
	r.Spec.Srt = false
	r.Spec.Opt = len(c.Flags) > 0
	// opt
	r.Opt = c.Flags.syntaxify()
	return r
}
