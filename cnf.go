package cnf

import (
	"slices"
	"strconv"
)

type Formula []Clause

func NewFormula(x [][]int) Formula {
	cs := make([]Clause, len(x))
	for i, v := range x {
		cs[i] = NewClause(v)
	}
	return Formula(cs)
}

func (f Formula) Int() [][]int {
	result := make([][]int, len(f))
	for i, c := range f {
		result[i] = c.Int()
	}
	return result
}

func (f Formula) SortBySize() {
	slices.SortFunc(f, func(a, b Clause) int {
		aa := len([]Lit(a))
		bb := len([]Lit(b))

		switch {
		case aa < bb:
			return -1
		case aa == bb:
			return 0
		default:
			return 1
		}
	})
}

type Clause []Lit

func NewClause(x []int) Clause {
	lits := make([]Lit, len(x))
	for i, v := range x {
		lits[i] = NewLit(v)
	}
	return Clause(lits)
}

func (c Clause) Int() []int {
	result := make([]int, len(c))
	for i, l := range c {
		result[i] = l.Int()
	}
	return result
}

type Lit int32

func NewLit(v int) Lit {
	s := v < 0
	if s {
		v = -v
	}
	return Lit(v + v + b2i(s))
}

func (l Lit) String() string { return strconv.Itoa(l.Int()) }
func (l Lit) Sign() bool     { return l&1 == 1 }
func (l Lit) Var() int       { return int(l >> 1) }
func (l Lit) Neg() Lit       { return l ^ 1 }

func (l Lit) Int() int {
	x := l.Var()
	if l.Sign() {
		x = -x
	}
	return x
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
