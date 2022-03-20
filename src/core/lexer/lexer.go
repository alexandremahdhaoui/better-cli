package lexer

import (
	"errors"
	"fmt"
	"main/core/regex"
)

type LExpr struct {
	Kind    string
	Content string
}

// Lexify - A lexer is basically a tokenizer,
// 	but it usually attaches extra context to the tokens:
// This token is a number, that token is a string literal,
// 	this other token is an equality operator.
func Lexify(t []string) ([]LExpr, error) {
	var l []LExpr
	for _, s := range t {
		switch {
		case regex.IsAlias(s, "long_flag"):
			l = append(l, LExpr{"long", cleanLong(s)})
		case regex.IsAlias(s, "short_flag"):
			l = append(l, LExpr{"short", cleanShort(s)})
		case regex.IsAlias(s, "word_dash_dot"):
			l = append(l, LExpr{"pos", s})
		default:
			return nil, errors.New(fmt.Sprintf("Lexer cannot resolve Kind for token: %s.", s))
		}
	}
	return l, nil
}

func cleanShort(s string) string {
	return s[1:]
}

func cleanLong(s string) string {
	return s[2:]
}
