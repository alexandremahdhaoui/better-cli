package api

import (
	"main/core/parser"
)

//----------------------------------------------------------------------------//
//--------------------------------- Arg Spec ---------------------------------//
//----------------------------------------------------------------------------//

type ArgSpecMethods interface {
	syntaxify() parser.PSpec
}

type ArgSpec struct {
	// Name
	Name string
	// Description used to generate the documentation
	Description string
	// Regex can either be:
	// 	- a Regex which must: start with `^`, end with `$`
	// 	- an alias to a Regex, eg: `lower_alpha_dash`, `float`
	// 		|-> The full list of aliases can be found at ../core/regex/regex.go
	Regex string
}

type ApiPSpec []*ArgSpec

func (ps ApiPSpec) syntaxify() parser.PSpec {
	var pSpec parser.PSpec
	pSpec.Pos = len(ps)
	for _, p := range ps {
		pSpec.PosR = append(pSpec.PosR, p.Regex)
	}
	return pSpec
}

//---------------------------------------------------------------------------//
//-------------------------------- Flag Spec --------------------------------//
//---------------------------------------------------------------------------//

type FlagSpecMethods interface {
	New()
	// Register this struct to cli/sc/c
	Register()
	syntaxify() parser.OptRule
}

type FlagSpec struct {
	Name        string
	Description string
	Short       string
	Long        string
	PArgs       ApiPSpec
}

type ApiFSpec []*FlagSpec

func (fs ApiFSpec) syntaxify() parser.OptRule {
	opt := parser.NewOptRule()
	for _, f := range fs {
		ps := f.PArgs.syntaxify()
		if f.Short != "" {
			opt.Short[f.Short] = ps
		}
		if f.Long != "" {
			opt.Long[f.Long] = ps
		}
	}
	return opt
}
