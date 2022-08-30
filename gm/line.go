package gm

import (
	"fmt"

	"github.com/add1609/lin/rn"
)

type Line struct {
	V1, V2 rn.Vec
}

func (o Line) String() (str string) {
	str += fmt.Sprintf("x = %v + λ * %v", o.V1, o.V2)
	return
}

// MakeLine returns a Line object that passes through P1 and has direction D1
//
// Parameters:
//
//	P1 rn.Vec - The point
//	D1 rn.Vec - The direction
//
// Returns:
//
//	line Line - The line that passes through P1 and has direction D1
func MakeLine(P1, D1 rn.Vec) (line Line) {
	line.V1 = P1
	line.V2 = D1
	return
}

// MakeLineByPoints returns a Line object that passes through the given points
//
// Parameters:
//
//	P1 rn.Vec - The first point
//	P2 rn.Vec - The second point
//
// Returns:
//
//	line Line - The line that passes through the given points
func MakeLineByPoints(P1, P2 rn.Vec) (line Line) {
	line.V1 = P1
	line.V2 = P2.Sub(P1)
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
	return o.V1.Equal(q.V1) && o.V2.Equal(q.V2)
}

// At returns the point at which the line arrives when scaled by s
//
// Parameters:
//
//	s float64 - The scalar
//
// Returns:
//
//	P1 rn.Vec - The point at which the line arrives when scaled by s
func (o *Line) At(s float64) (P1 rn.Vec) {
	P1 = o.V1.Add(o.V2.Scale(s))
	return
}

// IntersectPlane returns the intersection point of a line and a plane (if it exists)
//
// Parameters:
//
//	o *Line - The line
//	plane Plane - The plane
//
// Returns:
//
//	x rn.Vec - The value of λ and μ that satisfy the equation
//	P1 rn.Vec - The intersection point
//	err error - An error if one occurred
func (o *Line) IntersectPlane(plane Plane) (x, P1 rn.Vec, err error) {
	lgs := rn.MakeMat(3, 4, 0)
	lgs.SetCol(0, o.V2)
	lgs.SetCol(1, plane.V2.Scale(-1))
	lgs.SetCol(2, plane.V3.Scale(-1))
	lgs.SetCol(3, plane.V1.Sub(o.V1))
	_, x, err = lgs.GaussSolve()
	if err != nil {
		P1 = o.V1.Add(o.V2.Scale(x.Get(0)))
	}
	return
}

// IntersectLine returns the intersection point of two lines (if it exists)
//
// Parameters:
//
//	o *Line - The first line
//	q Line - The second line
//
// Returns:
//
//	x rn.Vec - The value of λ and μ that satisfy the equation
//	P1 rn.Vec - The intersection point
//	err error - An error if one occurred
func (o *Line) IntersectLine(q Line) (x, P1 rn.Vec, err error) {
	lgs := rn.MakeMat(3, 3, 0)
	lgs.SetCol(0, o.V2)
	lgs.SetCol(1, q.V2.Scale(-1))
	lgs.SetCol(2, q.V1.Sub(o.V1))
	_, x, err = lgs.GaussSolve()
	if err != nil {
		P1 = o.V1.Add(o.V2.Scale(x.Get(0)))
	}
	return
}
