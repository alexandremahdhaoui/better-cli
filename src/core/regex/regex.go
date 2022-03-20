package regex

import "regexp"

type R map[string]string

func getR(alias string) string {
	m := R{
		"empty_string":       `^$`,
		"int":                `^(\d+)$`,
		"float":              `^(\d+)\.(\d+)$`,
		"lower_alpha":        `^([a-z]+)$`,
		"lower_alpha_dash":   `^([a-z-]+)$`,
		"alpha":              `^([A-Za-z]+)$`,
		"alphanum":           `^([a-zA-Z0-9]+)$`,
		"aphanum_dash":       `^([a-zA-Z0-9-]+)$`,
		"word":               `^(\w+)$`,
		"word_dash":          `^([\w-]+)$`,
		"word_dash_dot":      `^([\w-\.]*)$`,
		"short_flag":         `^-([A-Za-z])$`,
		"chained_short_flag": `^-([A-Za-z]+[A-Za-z])$`,
		"long_flag":          `^--([\w-]+)$`,
		"long_flag_eq":       `^--([\w-]+)=([\w-]+)$`,
		"long_flag_eq_quote": `^--([\w-]+)=\"([\w-]*)\"$`,
	}
	return m[alias]
}

func IsR(s, regex string) bool {
	b, _ := regexp.MatchString(regex, s)
	return b
}

func IsAlias(s, alias string) bool {
	return IsR(s, getR(alias))
}

//"long_flag_eq":       `^--([a-z-]+)=([\w-]+)$`,
//"long_flag_eq_quote": `^--([a-z-]+)=\"([\w-]*)\"$`,
