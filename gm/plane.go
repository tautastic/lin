package gm

import (
	"fmt"

	"github.com/add1609/lin/rn"
)

type Plane struct {
	V1, V2, V3 rn.Vec
}

func (o Plane) String() (str string) {
	str += fmt.Sprintf("x = %v + λ * %v + μ * %v", o.V1, o.V2, o.V3)
	return
}

// MakePlane returns a plane that passes through P1 and has directions D1 and D2
//
// Parameters:
//
//	P rn.Vec - The point
//	D1 rn.Vec - The first direction
//	D2 rn.Vec - The second direction
//
// Returns:
//
//	plane Plane - The plane that passes through P1 and has directions D1 and D2
func MakePlane(P1, D1, D2 rn.Vec) (plane Plane) {
	plane.V1 = P1
	plane.V2 = D1
	plane.V3 = D2
	return
}

// MakePlaneByPoints returns a plane that passes through the given points
//
// Parameters:
//
//	P1 rn.Vec - The first point
//	P2 rn.Vec - The second point
//	P3 rn.Vec - The third point
//
// Returns:
//
//	plane Plane - The plane that passes through the given points
func MakePlaneByPoints(P1, P2, P3 rn.Vec) (plane Plane) {
	plane.V1 = P1
	plane.V2 = P2.Sub(P1)
	plane.V3 = P3.Sub(P1)
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
	return o.V1.Equal(q.V1) && o.V2.Equal(q.V2) && o.V3.Equal(q.V3)
}

// IntersectLine returns the intersection point of a plane and a line (if it exists)
//
// Parameters:
//
//	o *Plane - The plane
//	line Line - The line
//
// Returns:
//
//	x rn.Vec - The value of λ and μ that satisfy the equation
//	P1 rn.Vec - The intersection point
//	err error - An error if one occurred
func (o *Plane) IntersectLine(line Line) (x, P1 rn.Vec, err error) {
	lgs := rn.MakeMat(3, 4, 0)
	lgs.SetCol(0, line.V2)
	lgs.SetCol(1, o.V2.Scale(-1))
	lgs.SetCol(2, o.V3.Scale(-1))
	lgs.SetCol(3, o.V1.Sub(line.V1))
	_, x, err = lgs.GaussSolve()
	if err != nil {
		P1 = line.V1.Add(line.V2.Scale(x.Get(0)))
	}
	return
}
