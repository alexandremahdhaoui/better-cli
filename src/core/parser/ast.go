package parser

type AbstractFlag struct {
	Name string
	Pos  []string
}

type AbstractOpt struct {
	Short []*AbstractFlag
	Long  []*AbstractFlag
}

// AST - Abstract Syntax Tree
type AST struct {
	Name string
	Pos  []string
	Opt  *AbstractOpt
	Ast  *AST
}
