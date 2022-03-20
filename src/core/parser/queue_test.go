package parser

import (
	"main/core/lexer"
	"reflect"
	"testing"
)

// testCases checks equality between result & reference
func testCaseQ(g Queue, w Queue, err error, t *testing.T) {
	if !reflect.DeepEqual(g, w) || err != nil {
		t.Fatalf(
			"\ngot: %q\nwant: %q\nerr: %v",
			g, w, err,
		)
	}
}

func testCaseL(g lexer.LExpr, w lexer.LExpr, err error, t *testing.T) {
	if !reflect.DeepEqual(g, w) || err != nil {
		t.Fatalf(
			"\ngot: %q\nwant: %q\nerr: %v",
			g, w, err,
		)
	}
}

func testCaseBool(g bool, w bool, err error, t *testing.T) {
	if !reflect.DeepEqual(g, w) || err != nil {
		t.Fatalf(
			"\ngot: %t\nwant: %t\nerr: %v",
			g, w, err,
		)
	}
}

func TestIsEmpty(t *testing.T) {
	// test true
	q := Queue{}
	g1 := q.IsEmpty()
	// test false
	q = Queue{
		[]lexer.LExpr{
			{"pos", "txt"},
			{"short", "l"},
			{"long", "template"},
		},
	}
	g2 := q.IsEmpty()
	testCaseBool(g1, true, nil, t)
	testCaseBool(g2, false, nil, t)
}

func TestPop(t *testing.T) {
	g1 := Queue{
		[]lexer.LExpr{
			{"pos", "txt"},
			{"short", "l"},
			{"long", "template"},
		},
	}
	g2, g3 := g1.Pop()
	w1 := Queue{
		[]lexer.LExpr{
			{"pos", "txt"},
			{"short", "l"},
			{"long", "template"},
		},
	}
	w2 := Queue{
		[]lexer.LExpr{
			{"short", "l"},
			{"long", "template"},
		},
	}
	w3 := lexer.LExpr{Kind: "pos", Content: "txt"}
	testCaseQ(g1, w1, nil, t)
	testCaseQ(g2, w2, nil, t)
	testCaseL(g3, w3, nil, t)
}

func TestMutPop(t *testing.T) {
	g1 := Queue{
		[]lexer.LExpr{
			{"pos", "txt"},
			{"short", "l"},
			{"long", "template"},
		},
	}
	g2 := g1.MutPop()
	w1 := Queue{
		[]lexer.LExpr{
			{"short", "l"},
			{"long", "template"},
		},
	}
	w2 := lexer.LExpr{Kind: "pos", Content: "txt"}
	testCaseQ(g1, w1, nil, t)
	testCaseL(g2, w2, nil, t)
}
