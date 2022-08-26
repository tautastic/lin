package rn

import "fmt"

type Line struct {
	A, B Vec
}

func (l Line) String() string {
	return fmt.Sprintf("%v + r * %v", l.A.X, l.B.X)
}

// LineFromTo returns a line from p to q
//
//	l: x = p + r * (q - p).
func LineFromTo(p, q Vec) Line {
	return Line{p, Sub(q, p)}
}

// At returns the vector v = l(r)
func (l *Line) At(r float64) Vec {
	return Add(l.A, Scale(r, l.B))
}
