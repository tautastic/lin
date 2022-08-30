package rn

import (
	"github.com/add1609/lin/errors"
	"github.com/add1609/lin/scalar"
)

// BackSubstitution returns the solution vector x of an upper triangular matrix
//
// Parameters:
//
//	o *Mat - An upper triangular matrix
//
// Returns:
//
//	x Vec - The solution vector x
func (o *Mat) BackSubstitution() (x Vec) {
	uN := o.N - 1
	x = MakeVec(uN, 0)
	for i := uN - 1; 0 <= i; i-- {
		yI := o.Get(i, o.N-1)
		uII := o.Get(i, i)
		for j := i + 1; j < uN; j++ {
			if uIJ := o.Get(i, j); uIJ == 0 {
				continue
			} else {
				yI -= uIJ * x.Get(j)
			}
		}
		x.Set(i, scalar.RoundTo(yI/uII, 13))
	}
	return
}

// GaussSolve takes an upper triangular matrix and returns it's solution vector x
//
// Parameters:
//
//	o *Mat - An upper triangular matrix
//
// Returns:
//
//	mat Mat - The modified matrix
//	x Vec - The solution vector x
//	err error - An error if one occurred
func (o *Mat) GaussSolve() (mat Mat, x Vec, err error) {
	var rowPivot, colPivot int
	mat = o.GetCopy()

	if o.N < o.M {
		err = errors.ErrShape
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err = r.(errors.Error)
			return
		}
	}()

	for {
		if mat.N-1 < colPivot || mat.M-1 < rowPivot {
			x = mat.BackSubstitution()
			return
		}
		col := mat.GetCol(colPivot)
		colMax, colMaxIdx := col.Largest(rowPivot, col.N)
		if colMax == 0 {
			rowPivot += 1
		} else {
			mat.SwapRows(rowPivot, colMaxIdx)
			for i := rowPivot + 1; i < mat.M; i++ {
				f := mat.Get(i, colPivot) / mat.Get(rowPivot, colPivot)
				if i < mat.N {
					mat.Set(i, colPivot, 0)
				}
				for j := colPivot + 1; j < mat.N; j++ {
					mat.Set(i, j, mat.Get(i, j)-mat.Get(rowPivot, j)*f)
				}
			}
			rowPivot += 1
			colPivot += 1
		}
	}
}
