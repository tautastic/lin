package gm

import (
	"fmt"

	"github.com/add1609/lin/rn"
)

type Line struct {
	A, B rn.Vec
}

func (o Line) String() (str string) {
	str += fmt.Sprintf("x = %v + λ * %v", o.A, o.B)
	return
}

// MakeLine returns a Line object that passes through the given points
//
// Parameters:
//
//	P1 rn.Vec - The first point
//	P2 rn.Vec - The second point
//
// Returns:
//
//	line Line - The line passing through the given points
func MakeLine(P1, P2 rn.Vec) (line Line) {
	line.A = P1
	line.B = P2.Sub(P1)
	return
}

// Equal returns true if o and q are equal
//
// Parameters:
//
//	o *Line - line to compare to q
//	q Line - line to compare to o
//
// Returns:
//
//	bool - true if o and q are equal
func (o *Line) Equal(q Line) bool {
	return o.A.Equal(q.A) && o.B.Equal(q.B)
}

// IntersectPlane returns the intersection point of a plane and a line
//
// Parameters:
//
//	o *Line - The line
//	p Plane - The plane
//
// Returns:
//
//	x rn.Vec - The value of λ and μ that satisfy the equation of the line and the plane
//	P rn.Vec - The intersection point
func (o *Line) IntersectPlane(p Plane) (x rn.Vec, P rn.Vec) {
	lgs := rn.MakeMat(3, 4, 0)
	lgs.SetCol(0, o.B)
	lgs.SetCol(1, p.B.Scale(-1))
	lgs.SetCol(2, p.C.Scale(-1))
	lgs.SetCol(3, p.A.Sub(o.A))
	_, x = lgs.GaussSolve()
	P = o.A.Add(o.B.Scale(x.Get(0)))
	return
}
