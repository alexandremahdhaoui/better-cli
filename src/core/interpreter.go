package core

import (
	"log"
	"main/core/lexer"
	"main/core/parser"
	"main/core/tokenizer"
)

type F func(*parser.AST)

// Interpret
//	Takes a String, a Func & an SRT as input
//	1. Tokenize
//	2. Lexify
//	3. Parse
//	4. Execute Functional Program with AST & SRT
func Interpret(s string, f F, srt *parser.SRT) {
	//	1. Tokenize
	t, err := tokenizer.Tokenize(s)
	checkErr(err)

	//	2. Lexify
	l, err := lexer.Lexify(t)
	checkErr(err)

	//	3. Parse
	p := parser.New(l, srt)
	ast := p.Parse()

	// 4. Execute
	f(ast)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
