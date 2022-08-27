package rn

import (
	"fmt"
	"math"
)

type Vec struct {
	N int
	X []float64
}

func roundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func (o Vec) String() (str string) {
	for i := 0; i < o.N-1; i++ {
		if i == 0 {
			str += fmt.Sprintf("[%v, ", roundTo(o.X[i], 4))
		} else {
			str += fmt.Sprintf("%v, ", roundTo(o.X[i], 4))
		}
	}
	str += fmt.Sprintf("%v]", roundTo(o.X[o.N-1], 4))
	return
}

// MakeVec returns a Vec object that has length of n and all elements set to val
func MakeVec(n int, val float64) (o Vec) {
	if n < 1 {
		panic(errNegativeDimension)
	}
	o = Vec{N: n, X: make([]float64, n)}
	if val != 0 {
		for i := range o.X {
			o.X[i] = val
		}
	}
	return
}

// Equal returns true if o.N == q.N and o.X == q.X
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
func (o *Vec) Get(i int) float64 {
	if o.N <= i {
		panic(errVectorAccess)
	}
	return o.X[i]
}

// Set sets the value of o[i] = val
func (o *Vec) Set(i int, val float64) {
	if o.N <= i {
		panic(errVectorAccess)
	}
	o.X[i] = val
}

// SetSlice copies the float64 slice s onto o.X
func (o *Vec) SetSlice(s []float64) {
	if o.N < len(s) {
		panic(errSliceLengthMismatch)
	}
	copy(o.X, s)
}

// Abs returns the absolute value of o
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

// Sub returns the vector sum of o and -q
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

// Scale returns the vector o scaled by r
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

// Norm returns the Euclidean norm of o
//
//	||o|| = sqrt(o · o)
func (o *Vec) Norm() (nrm float64) {
	nrm = math.Sqrt(o.Dot(*o))
	return
}

// Dist returns the distance between o and q
//
//	||q - o||
func (o *Vec) Dist(q Vec) (dist float64) {
	sub := q.Sub(*o)
	dist = sub.Norm()
	return
}

// Cos returns the cosine of the opening angle between o and q
func (o *Vec) Cos(q Vec) (cos float64) {
	cos = o.Dot(q) / (o.Norm() * q.Norm())
	return
}

// Largest returns the largest component |a[ij]| of this vector and it's index
func (o *Vec) Largest(begin, end int) (largest float64, idx int) {
	if o.N < 1 {
		panic(errZeroLengthVec)
	}
	largest = math.Abs(o.X[begin])
	idx = begin
	for k := begin + 1; k < o.N && k < end; k++ {
		tmp := math.Abs(o.X[k])
		if tmp > largest {
			largest = tmp
			idx = k
		}
	}
	return
}
