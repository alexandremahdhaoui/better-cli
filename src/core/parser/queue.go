package parser

import (
	"main/core/lexer"
)

type Queue struct {
	items []lexer.LExpr
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) < 1
}

// Pop item from the Queue and ensures Immutability
func (q *Queue) Pop() (Queue, lexer.LExpr) {
	l := q.items[0]
	q_ := Queue{q.items[1:]}
	return q_, l
}

// MutPop - Pop item from the Queue and ensures Mutability
func (q *Queue) MutPop() lexer.LExpr {
	l := q.items[0]
	q.items = q.items[1:]
	return l
}
