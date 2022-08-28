package rn

import (
	"fmt"
	"math"

	"github.com/add1609/lin/scalar"
)

type Vec struct {
	N int
	X []float64
}

func (o Vec) String() (str string) {
	for i := 0; i < o.N-1; i++ {
		if i == 0 {
			str += fmt.Sprintf("[%v, ", scalar.RoundTo(o.X[i], 4))
		} else {
			str += fmt.Sprintf("%v, ", scalar.RoundTo(o.X[i], 4))
		}
	}
	str += fmt.Sprintf("%v]", scalar.RoundTo(o.X[o.N-1], 4))
	return
}

// MakeVec returns a Vec object that has dimension n and all elements set to val
//
// Parameters:
//
//	n int - dimension of the vector
//	val float64 - value of the elements of the vector
//
// Returns:
//
//	Vec: Vec object with dimension n and all elements set to val
func MakeVec(n int, val float64) (vec Vec) {
	if n < 1 {
		panic(errNegativeDimension)
	}
	vec = Vec{N: n, X: make([]float64, n)}
	if val != 0 {
		for i := range vec.X {
			vec.X[i] = val
		}
	}
	return
}

// Equal returns true if o and q are equal
//
// Parameters:
//
//	o *Vec - vector to compare to q
//	q Vec - vector to compare to o
//
// Returns:
//
//	bool - true if o and q are equal
func (o *Vec) Equal(q Vec) bool {
	if o.N != q.N {
		return false
	}
	if o.N < 1 || q.N < 1 {
		panic(errZeroLengthVec)
	}
	for i, v := range o.X {
		if v != q.X[i] && !(math.IsNaN(v) || math.IsNaN(q.X[i])) {
			return false
		}
	}
	return true
}

// Get returns the value of the vector o at index i
//
// Parameters:
//
//	i int - index of the element to return
//
// Returns:
//
//	val float64 - value of the element at index i
func (o *Vec) Get(i int) (val float64) {
	if o.N <= i {
		panic(errVectorAccess)
	}
	return o.X[i]
}

// Set sets the value of the vector o at index i to val
//
// Parameters:
//
//	i int - index of the element to set
//	val float64 - value to set the element to
//
// Returns:
//
// none
func (o *Vec) Set(i int, val float64) {
	if o.N <= i {
		panic(errVectorAccess)
	}
	o.X[i] = val
}

// SetSlice sets the values of the vector o to the values of the slice s
//
// Parameters:
//
//	s []float64 - slice of values to set the vector to
//
// Returns:
//
// none
func (o *Vec) SetSlice(s []float64) {
	if o.N < len(s) {
		panic(errSliceLengthMismatch)
	}
	copy(o.X, s)
}

// Abs returns the absolute value of o
//
// Parameters:
//
//	o *Vec - vector to take the absolute value of
//
// Returns:
//
//	u Vec - absolute value of o
func (o *Vec) Abs() (u Vec) {
	if o.N < 1 {
		panic(errZeroLengthVec)
	}
	u = Vec{N: o.N, X: make([]float64, o.N)}
	for i := 0; i < len(u.X); i++ {
		u.Set(i, math.Abs(o.X[i]))
	}
	return
}

// Add returns the vector sum of o and q
//
// Parameters:
//
//	o *Vec - vector to add to
//	q Vec - vector to add
//
// Returns:
//
//	u Vec - vector sum of o and q
func (o *Vec) Add(q Vec) (u Vec) {
	if o.N != q.N {
		panic(errShape)
	}
	if o.N < 1 || q.N < 1 {
		panic(errZeroLengthVec)
	}
	u = Vec{N: o.N, X: make([]float64, o.N)}
	for i := 0; i < len(u.X); i++ {
		u.Set(i, o.X[i]+q.X[i])
	}
	return
}

// Sub returns the vector difference of o and q
//
// Parameters:
//
//	o *Vec - vector to subtract from
//	q Vec - vector to subtract
//
// Returns:
//
//	u Vec - vector difference of o and q
func (o *Vec) Sub(q Vec) (u Vec) {
	if o.N != q.N {
		panic(errShape)
	}
	if o.N < 1 || q.N < 1 {
		panic(errZeroLengthVec)
	}
	u = Vec{N: o.N, X: make([]float64, o.N)}
	for i := 0; i < len(u.X); i++ {
		u.Set(i, o.X[i]-q.X[i])
	}
	return
}

// Scale returns the vector o scaled by s
//
// Parameters:
//
//	o *Vec - vector to scale
//	s float64 - scalar to scale o by
//
// Returns:
//
//	u Vec - vector o scaled by s
func (o *Vec) Scale(r float64) (u Vec) {
	if o.N < 1 {
		panic(errZeroLengthVec)
	}
	u = Vec{N: o.N, X: make([]float64, o.N)}
	for i := 0; i < len(u.X); i++ {
		u.Set(i, r*o.X[i])
	}
	return
}

// Dot returns the dot product of o and q
//
// Parameters:
//
//	o *Vec - vector to take the dot product of
//	q Vec - vector to take the dot product of
//
// Returns:
//
//	d float64 - dot product of o and q
func (o *Vec) Dot(q Vec) (d float64) {
	if o.N != q.N {
		panic(errShape)
	}
	if o.N < 1 || q.N < 1 {
		panic(errZeroLengthVec)
	}
	for i := 0; i < o.N; i++ {
		d += o.X[i] * q.X[i]
	}
	return
}

// Cross returns the cross product of o and q
//
// Parameters:
//
//	o *Vec - vector to take the cross product of
//	q Vec - vector to take the cross product of
//
// Returns:
//
//	u Vec - cross product of o and q
func (o *Vec) Cross(q Vec) (u Vec) {
	if o.N != q.N {
		panic(errShape)
	}
	if o.N < 1 || q.N < 1 {
		panic(errZeroLengthVec)
	}
	switch o.N {
	default:
		panic(errOrder)
	case 3:
		u = Vec{N: 3, X: make([]float64, 3)}
		u.X[0] = (o.X[1] * q.X[2]) - (o.X[2] * q.X[1])
		u.X[1] = (o.X[2] * q.X[0]) - (o.X[0] * q.X[2])
		u.X[2] = (o.X[0] * q.X[1]) - (o.X[1] * q.X[0])
		return
	}
}

// Norm returns the norm of o
//
// Parameters:
//
//	o *Vec - vector to take the norm of
//
// Returns:
//
//	nrm float64 - norm of o
func (o *Vec) Norm() (nrm float64) {
	nrm = math.Sqrt(o.Dot(*o))
	return
}

// Dist returns the distance between o and q
//
// Parameters:
//
//	o *Vec - vector to take the distance from
//	q Vec - vector to take the distance to
//
// Returns:
//
//	dist float64 - distance between o and q
func (o *Vec) Dist(q Vec) (dist float64) {
	sub := q.Sub(*o)
	dist = sub.Norm()
	return
}

// Cos returns the cosine of the angle between o and q
//
// Parameters:
//
//	o *Vec - vector to take the cosine of
//	q Vec - vector to take the cosine of
//
// Returns:
//
//	cos float64 - cosine of the angle between o and q
func (o *Vec) Cos(q Vec) (cos float64) {
	cos = o.Dot(q) / (o.Norm() * q.Norm())
	return
}

// Largest returns the largest element of o
//
// Parameters:
//
//	o *Vec - vector to find the largest element of
//	begin int - index to begin searching from
//	end int - index to end searching at
//
// Returns:
//
//	val float64 - largest element of o
//	idx int - index of the largest element of o
func (o *Vec) Largest(begin, end int) (val float64, idx int) {
	if o.N < 1 {
		panic(errZeroLengthVec)
	}
	val = math.Abs(o.X[begin])
	idx = begin
	for k := begin + 1; k < o.N && k < end; k++ {
		tmp := math.Abs(o.X[k])
		if tmp > val {
			val = tmp
			idx = k
		}
	}
	return
}
