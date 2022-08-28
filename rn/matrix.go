package rn

import (
	"fmt"
	"math"

	"github.com/add1609/lin/scalar"
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
			str += fmt.Sprintf("%9g ", scalar.RoundTo(o.Get(i, j), 3))
		}
	}
	return
}

// MakeMat returns a new matrix with m rows and n columns and all elements set to val
//
// Parameters:
//
//	m int - number of rows
//	n int - number of columns
//	val float64 - value to initialize the matrix with
//
// Returns:
//
//	mat Mat - a new matrix with m rows and n columns and all elements set to val
func MakeMat(m, n int, val float64) (mat Mat) {
	if m < 1 || n < 1 {
		panic(errNegativeDimension)
	}
	mat = Mat{M: m, N: n, Data: make([]float64, m*n)}
	if val != 0 {
		for k := range mat.Data {
			mat.Data[k] = val
		}
	}
	return
}

// GetCopy returns a copy of this matrix
//
// Parameters:
//
//	o *Mat - matrix to copy
//
// Returns:
//
//	clone Mat - a copy of this matrix
func (o *Mat) GetCopy() (clone Mat) {
	if o.M < 1 || o.N < 1 {
		panic(errZeroLengthVec)
	}
	clone = MakeMat(o.M, o.N, 0)
	copy(clone.Data, o.Data)
	return
}

// Get returns the value at A[i][j]
//
// Parameters:
//
//	o *Mat - matrix to get value at A[i][j] of
//	i int - row index
//	j int - column index
//
// Returns:
//
//	val float64 - the value at A[i][j]
func (o *Mat) Get(i, j int) (val float64) {
	if o.M <= i {
		panic(errRowAccess)
	}
	if o.N <= j {
		panic(errColAccess)
	}
	return o.Data[i+j*o.M]
}

// Set sets the value at A[i][j]
//
// Parameters:
//
//	o *Mat - matrix to set value at A[i][j] of
//	i int - row index
//	j int - column index
//	val float64 - the value to set at A[i][j]
//
// Returns:
//
//	none
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
//
// Parameters:
//
//	o *Mat - matrix to get column j of
//	j int - column index
//
// Returns:
//
//	col Vec - column j of this matrix
func (o *Mat) GetCol(j int) (col Vec) {
	if o.N <= j {
		panic(errColAccess)
	}
	col = Vec{N: o.M, X: make([]float64, o.M)}
	copy(col.X, o.Data[j*o.M:(j+1)*o.M])
	return
}

// SetCol sets the values of a column j to the values of a Vec v
//
// Parameters:
//
//	o *Mat - matrix to set column j of
//	j int - column index
//	v Vec - column values
//
// Returns:
//
//	none
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
//
// Parameters:
//
//	o *Mat - matrix to get row i of
//	i int - row index
//
// Returns:
//
//	row Vec - row i of this matrix
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
//
// Parameters:
//
//	o *Mat - matrix to set row i of
//	i int - row index
//	v Vec - row values
//
// Returns:
//
//	none
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

// SwapRows swaps row i with row j
//
// Parameters:
//
//	o *Mat - matrix to swap rows i and j of
//	i int - row index
//	j int - row index
//
// Returns:
//
//	none
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

// Largest returns the largest element in this matrix
//
// Parameters:
//
//	o *Mat - matrix to find largest element of
//
// Returns:
//
//	val float64 - the largest element in this matrix
//	idx int - the index of the largest element in this matrix
func (o *Mat) Largest() (val float64, idx int) {
	if o.M < 1 || o.N < 1 {
		panic(errNegativeDimension)
	}
	val = math.Abs(o.Data[0])
	for k := 1; k < o.M*o.N; k++ {
		tmp := math.Abs(o.Data[k])
		if tmp > val {
			val = tmp
			idx = k
		}
	}
	return
}
