package regex

import (
	"testing"
)

// go test core/regex/

// testCase checks equality between result & reference
func testCase(res bool, f string, t *testing.T) {
	if res != true {
		t.Fatalf("Testing: `%s`\nExpected: `true`\nReceived `%t`\n",
			f, res)
	}
}

func TestIsRegex(t *testing.T) {
	var s, f string
	// IsEmptyString
	s, f = "", "IsEmptyString"
	testCase(IsAlias(s, "empty_string"), f, t)
	// IsInteger
	s, f = "465", "IsInteger"
	testCase(IsAlias(s, "int"), f, t)
	// IsFloat
	s, f = "3.14", "IsFloat"
	testCase(IsAlias(s, "float"), f, t)
	// IsLowerAlphabetical
	s, f = "kubectl", "IsLowerAlphabetical"
	testCase(IsAlias(s, "lower_alpha"), f, t)
	// IsLowerAlphabeticalDash
	s, f = "fro-mage", "IsLowerAlphabeticalDash"
	testCase(IsAlias(s, "lower_alpha_dash"), f, t)
	// IsAlphabetical
	s, f = "FroMage", "IsAlphabetical"
	testCase(IsAlias(s, "alpha"), f, t)
	// IsAlphanumerical
	s, f = "Fr0m4gE", "IsAlphanumerical"
	testCase(IsAlias(s, "alphanum"), f, t)
	// IsAlphanumericalDash
	s, f = "Fr0-M4ge", "IsAlphanumericalDash"
	testCase(IsAlias(s, "aphanum_dash"), f, t)
	// IsWord
	s, f = "Fro_M4ge", "IsWord"
	testCase(IsAlias(s, "word"), f, t)
	// IsWordDash
	s, f = "From_-_--4ge", "IsWordDash"
	testCase(IsAlias(s, "word_dash"), f, t)
	// IsWordDashDot
	s, f = "Fr0.m_-._--aGe.com", "IsWordDashDot"
	testCase(IsAlias(s, "word_dash_dot"), f, t)
	// IsShortFlags
	s, f = "-f", "IsShortFlags"
	testCase(IsAlias(s, "short_flag"), f, t)
	// IsChainedShortFlag
	s, f = "-ial", "IsChainedShortFlag"
	testCase(IsAlias(s, "chained_short_flag"), f, t)
	// IsLongFlag
	s, f = "--dry-run", "IsLongFlag"
	testCase(IsAlias(s, "long_flag"), f, t)
	// IsLongFlagWithEqualSign
	s, f = "--example=Fr0-M4ge", "IsLongFlagEq"
	testCase(IsAlias(s, "long_flag_eq"), f, t)
	// IsLongFlagWithEqualSign
	s, f = "--example=\"Fr0-M4ge\"", "IsLongFlagEqQuote"
	testCase(IsAlias(s, "long_flag_eq_quote"), f, t)
}
