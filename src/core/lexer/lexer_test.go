package lexer

import (
	"reflect"
	"testing"
)

// go test core/lexer/

// testCase checks equality between result & reference
func testCase(g []LExpr, w []LExpr, err error, t *testing.T) {
	if !reflect.DeepEqual(g, w) || err != nil {
		t.Fatalf(
			"\ngot: %q\nwant: %q\nerr: %v",
			g, w, err,
		)
	}
}

// TestLex calls Lexify
func TestLex(t *testing.T) {
	x := []string{
		"txt", "ls", "folder", "-i", "-a", "-l", "--template", "empty",
	}
	g, err := Lexify(x)
	w := []LExpr{
		{"pos", "txt"},
		{"pos", "ls"},
		{"pos", "folder"},
		{"short", "i"},
		{"short", "a"},
		{"short", "l"},
		{"long", "template"},
		{"pos", "empty"},
	}
	testCase(g, w, err, t)
}
