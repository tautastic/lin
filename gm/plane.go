package gm

import "C"
import (
	"fmt"

	"github.com/add1609/lin/rn"
)

type Plane struct {
	A, B, C rn.Vec
}

func (o Plane) String() (str string) {
	str += fmt.Sprintf("x = %v + λ * %v + μ * %v", o.A, o.B, o.C)
	return
}

// MakePlane returns a Plane object that passes through the given points
//
// Parameters:
//
//	P1 rn.Vec - The first point
//	P2 rn.Vec - The second point
//	P3 rn.Vec - The third point
//
// Returns:
//
//	plane Plane - The plane passing through the given points
func MakePlane(P1, P2, P3 rn.Vec) (plane Plane) {
	plane.A = P1
	plane.B = P2.Sub(P1)
	plane.C = P3.Sub(P1)
	return
}

// Equal returns true if o and q are equal
//
// Parameters:
//
//	o *Plane - plane to compare to q
//	q Plane - plane to compare to o
//
// Returns:
//
//	bool - true if o and q are equal
func (o *Plane) Equal(q Plane) bool {
	return o.A.Equal(q.A) && o.B.Equal(q.B) && o.C.Equal(q.C)
}

// IntersectLine returns the intersection point of a plane and a line
//
// Parameters:
//
//	o *Plane - The plane
//	l Line - The line
//
// Returns:
//
//	x rn.Vec - The value of λ and μ that satisfy the equation of the line and the plane
//	P rn.Vec - The intersection point of the plane and the line
func (o *Plane) IntersectLine(l Line) (x rn.Vec, P rn.Vec) {
	lgs := rn.MakeMat(3, 4, 0)
	lgs.SetCol(0, l.B)
	lgs.SetCol(1, o.B.Scale(-1))
	lgs.SetCol(2, o.C.Scale(-1))
	lgs.SetCol(3, o.A.Sub(l.A))
	_, x = lgs.GaussSolve()
	P = l.A.Add(l.B.Scale(x.Get(0)))
	return
}
