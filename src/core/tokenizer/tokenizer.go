package tokenizer

import (
	"fmt"
	"main/core/regex"
	"strings"
)

// unpack takes a []string and appends its content to the end of
//	the token slice
func extend(t []string, t_ []string) []string {
	for _, s := range t_ {
		t = append(t, s)
	}
	return t
}

// Tokenize reads tokens one at a time from the input stream
// 	and pass the tokens to the lexer.
func Tokenize(s string) ([]string, error) {
	pre, err := fromInput(s)
	var t []string
	for _, s := range pre {
		switch {
		case regex.IsAlias(s, "chained_short_flag"):
			t_, _ := fromChainedShortFlags(s)
			t = extend(t, t_)
		case regex.IsAlias(s, "long_flag_eq") ||
			regex.IsAlias(s, "long_flag_eq_quote"):
			t_, _ := fromLongFlagEq(s)
			t = extend(t, t_)
		default:
			t = append(t, s)
		}
	}
	return t, err
}

// fromInput
// 	Tokenize input strings
func fromInput(s string) ([]string, error) {
	t := strings.Split(s, " ")
	return t, nil
}

// fromChainedShortFlags
//	Takes chained short flags and tokenize them into
// 	several short flags
func fromChainedShortFlags(s string) ([]string, error) {
	var t []string
	for _, flag := range s[1:] {
		t = append(t, fmt.Sprintf("-%s", string(flag)))
	}
	return t, nil
}

// fromLongFlagEq
// 	Takes long flag with equal sign and/or quotes,
//	clean it and return an array containing the flag and its value
func fromLongFlagEq(s string) ([]string, error) {
	s = strings.ReplaceAll(s, `"`, "")
	t := strings.Split(s, "=")
	return t, nil
}
