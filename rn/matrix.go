package rn

import (
	"fmt"
	"math"
)

// Mat implements a column-major representation of a matrix
//
//	Example:
//	           _      _
//	          |  0  3  |
//	      A = |  1  4  |
//	          |_ 2  5 _|(m x n)
//
//	Data[i+j*m] = A[i][j]
type Mat struct {
	M, N int       // M, N => col size, row size
	Data []float64 // column-major data array
}

func (o Mat) String() (str string) {
	for i := 0; i < o.M; i++ {
		if i > 0 {
			str += "\n"
		}
		for j := 0; j < o.N; j++ {
			str += fmt.Sprintf("%9g ", roundTo(o.Get(i, j), 3))
		}
	}
	return
}

// MakeMat returns a Mat object with row size of M, column size of N and all elements set to val
func MakeMat(m, n int, val float64) (o Mat) {
	if m < 1 || n < 1 {
		panic(errNegativeDimension)
	}
	o = Mat{M: m, N: n, Data: make([]float64, m*n)}
	if val != 0 {
		for k := range o.Data {
			o.Data[k] = val
		}
	}
	return
}

// GetCopy returns a copy of this matrix
func (o *Mat) GetCopy() (clone Mat) {
	if o.M < 1 || o.N < 1 {
		panic(errZeroLengthVec)
	}
	clone = MakeMat(o.M, o.N, 0)
	copy(clone.Data, o.Data)
	return
}

// Get gets the value at A[i][j]
func (o *Mat) Get(i, j int) float64 {
	if o.M <= i {
		panic(errRowAccess)
	}
	if o.N <= j {
		panic(errColAccess)
	}
	return o.Data[i+j*o.M]
}

// Set sets the value at A[i][j]
func (o *Mat) Set(i, j int, val float64) {
	if o.M <= i {
		panic(errRowAccess)
	}
	if o.N <= j {
		panic(errColAccess)
	}
	o.Data[i+j*o.M] = val
}

// GetCol returns column j of this matrix
func (o *Mat) GetCol(j int) (col Vec) {
	if o.N <= j {
		panic(errColAccess)
	}
	col = Vec{N: o.M, X: make([]float64, o.M)}
	copy(col.X, o.Data[j*o.M:(j+1)*o.M])
	return
}

// SetCol sets the values of a column j to the values of a Vec v
func (o *Mat) SetCol(j int, v Vec) {
	if v.N != o.M {
		panic(errColLength)
	}
	if o.N <= j {
		panic(errColAccess)
	}
	copy(o.Data[j*o.M:(j+1)*o.M], v.X)
}

// GetRow returns row i of this matrix
func (o *Mat) GetRow(i int) (row Vec) {
	if o.M < 1 || o.N < 1 {
		panic(errNegativeDimension)
	}
	if o.M <= i {
		panic(errRowAccess)
	}
	row = Vec{N: o.N, X: make([]float64, o.N)}
	for j := 0; j < o.N; j++ {
		row.Set(j, o.Data[i+j*o.M])
	}
	return
}

// SetRow sets the values of a row i to the values of a Vec v
func (o *Mat) SetRow(i int, v Vec) {
	if v.N != o.N {
		panic(errRowLength)
	}
	if o.M <= i {
		panic(errRowAccess)
	}
	if o.M < 1 || o.N < 1 {
		panic(errNegativeDimension)
	}
	for j := 0; j < o.N; j++ {
		o.Data[i+j*o.M] = v.Get(j)
	}
}

// SwapRows sets the values of a row i to the values of a row j
func (o *Mat) SwapRows(i, j int) {
	if o.M <= i {
		panic(errRowAccess)
	}
	if o.N <= j {
		panic(errColAccess)
	}
	tmp := o.GetRow(i)
	o.SetRow(i, o.GetRow(j))
	o.SetRow(j, tmp)
}

// Largest returns the largest component |a[ij]| of this matrix and it's index
//
//	largest := |a[ij]|
func (o *Mat) Largest() (largest float64, idx int) {
	if o.M < 1 || o.N < 1 {
		panic(errNegativeDimension)
	}
	largest = math.Abs(o.Data[0])
	for k := 1; k < o.M*o.N; k++ {
		tmp := math.Abs(o.Data[k])
		if tmp > largest {
			largest = tmp
			idx = k
		}
	}
	return
}
