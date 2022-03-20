package tokenizer

import (
	"reflect"
	"testing"
)

// go test ./src/util/core/tokenizer/

// testCase checks equality between result & reference
func testCase(g []string, w []string, err error, t *testing.T) {
	if !reflect.DeepEqual(g, w) || err != nil {
		t.Fatalf(
			"\ngot: %q\nwant: %q\n%v",
			g, w, err,
		)
	}
}

// TestTokenize calls Tokenize, checks for a valid return value
func TestTokenize(t *testing.T) {
	test := `txt ls folder -ial --template="empty"`
	g, err := Tokenize(test)
	w := []string{
		"txt", "ls", "folder", "-i", "-a", "-l", "--template", "empty",
	}
	testCase(g, w, err, t)
}

// TestFromInput calls fromInput, checks for a valid return value
func TestFromInput(t *testing.T) {
	test := `txt new filename.txt -f --template="empty"`
	g, err := fromInput(test)
	w := []string{"txt", "new", "filename.txt", "-f", "--template=\"empty\""}
	testCase(g, w, err, t)
}

// TestFromChainedShortFlags calls fromChainedShortFlags,
// checks for a valid return values
func TestFromChainedShortFlags(t *testing.T) {
	test := "-ial"
	g, err := fromChainedShortFlags(test)
	w := []string{"-i", "-a", "-l"}
	testCase(g, w, err, t)
}

// TestFromLongFlag calls fromLongFlagEq, checks for a valid return value
func TestFromLongFlag(t *testing.T) {
	test := "--template=\"empty\""
	g, err := fromLongFlagEq(test)
	w := []string{"--template", "empty"}
	testCase(g, w, err, t)
}
