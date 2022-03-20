package parser

import (
	"log"
	"main/core/lexer"
	"main/core/regex"
)

// Parser -
type Parser struct {
	q   *Queue
	ast *AST
	srt *SRT
}

func New(l []lexer.LExpr, srt *SRT) *Parser {
	q := Queue{l}
	ast := AST{}
	return &Parser{&q, &ast, srt}
}

func newFromQueue(q *Queue, srt *SRT) *Parser {
	ast := AST{}
	return &Parser{q, &ast, srt}
}

// Parse - The job of the parser is to detect syntax errors and build a parse tree.
// 	A parser takes the stream of tokens from the lexer
//	 and turns it into an abstract syntax tree representing
//   the program.
// 	 AST is then fed to the core that will execute the programm
//
// 	This algorithm is a variant of Recursive Descent Parsing
//
// 	To parse []LExpr, our parser needs a set of syntax rules called SRT (Syntax Rule Tree).
// 		The algorithm walked down each element of []LExpr,
//  	 in order to build the AST recursively.
//		Parser:
//		 Instantiate AST
//		Parse
//		 Pops the first LExpr out of the []LExpr
//		 Checks if SRT[LExpr.content] finds a match
//		If not match:
//			return LExpr.content
//		If match found
//			rule := SRT[LExpr.content]
//			AST[LExpr.content] = Parse(l, rule, AST)
//
func (p Parser) Parse() *AST {
	// If called with an empty Queue returns the ast
	if p.q.IsEmpty() {
		return p.ast
	}
	newQ, l := p.q.Pop()
	q := &newQ
	guardLExprPArg(q)
	// Get Rule & PSpec from specified entry of SRT[l.Content]
	rule := p.srt.getRule(l.Content)
	spec := &rule.Spec
	opt := &rule.Opt
	ast := p.ast
	// Set ast.name as l.Content
	ast.Name = l.Content
	// Parse Positional Arguments if Rule.PSpec.pos > 1
	ast.Pos = posParse(q, &PSpec{Pos: spec.Pos, PosR: spec.PosR})
	// Parse Optional Arguments if Rule.PSpec.opt == true
	if spec.Opt {
		ast.Opt = optParse(q, opt)
	} else if spec.Srt {
		// Create new parser to recursively parse the remaining tokens
		// Call Parser.New with the sub-srt of our current SRT
		recP := newFromQueue(q, &rule.Srt)
		// Parse & Assign the recursive AST to the current AST.ast
		ast.Ast = recP.Parse()
	}
	return ast
}

//----------------------------------------------------------------------------//
//------------------------ Parse Positional Arguments ------------------------//
//----------------------------------------------------------------------------//
func posParse(q *Queue, spec *PSpec) []string {
	var pos []string
	for i := 0; i < spec.Pos; i++ {
		guardEmptyQueuePArg(q, spec.Pos, i)
		guardLExprPArg(q)
		l := q.MutPop()
		guardRegex(l.Content, spec.PosR[i])
		pos = append(pos, l.Content)
	}
	return pos
}

//----------------------------------------------------------------------------//
//------------------------- Parse Optional Arguments -------------------------//
//----------------------------------------------------------------------------//
// optParse assumes the whole q is consumable
//
func optParse(q *Queue, opt *OptRule) *AbstractOpt {
	absOpt := AbstractOpt{}
	if q.IsEmpty() {
		return &absOpt
	}
	guardOptArgQueue(q)
	for {
		if q.IsEmpty() {
			break
		}
		l := q.MutPop()
		// Check if l in OptRule.
		switch l.Kind {
		case "long":
			absOpt.Long = append(absOpt.Long, flagParse(l.Content, q, opt.Long))
		case "short":
			absOpt.Short = append(absOpt.Short, flagParse(l.Content, q, opt.Short))
		}
	}
	return &absOpt
}

func flagParse(s string, q *Queue, f FSpec) *AbstractFlag {
	spec := f[s]
	pos := posParse(q, &spec)
	return &AbstractFlag{Name: s, Pos: pos}
}

//---------------------------------------------------------------------------//
//----------------------------- Check Functions -----------------------------//
//---------------------------------------------------------------------------//

// guardEmptyQueuePArg throws err if q empty while expecting positional arguments
func guardEmptyQueuePArg(q *Queue, e int, r int) {
	if q.IsEmpty() {
		log.Fatalf("Syntax Error: Expected %d positional arguments, Received %d\n",
			e, r)
	}
}

// guardLExprPArg checks if l is indeed a positional argument
func guardLExprPArg(q *Queue) {
	if q.items[0].Kind != "pos" {
		log.Fatalf("Syntax Error: Expected Positional Arguments, Received %s-flag", q.items[0].Kind)
	}
}

// guardRegex checks if l.content satisfies spec.posR
func guardRegex(s, r string) {
	if regex.IsAlias(s, r) {
		return
	}
	log.Fatalf("Syntax Error: Expected Positional Arguments with RegexType=`%s`, Received %s", r, s)
}

func guardOptArgQueue(q *Queue) {
	if q.items[0].Kind == "pos" {
		log.Fatalf("Syntax Error: Expected `Optional` Arguments, Received `Positional`")
	}
}
