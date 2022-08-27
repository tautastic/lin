package rn

// Error represents matrix/vector handling errors
type Error struct{ string }

func (err Error) Error() string { return err.string }

var (
	errPivot               = Error{"lin: malformed pivot list"}
	errOrder               = Error{"lin: invalid order for matrix"}
	errShape               = Error{"lin: dimension mismatch"}
	errSquare              = Error{"lin: expect square matrix"}
	errSingular            = Error{"lin: matrix is singular"}
	errColLength           = Error{"lin: col length mismatch"}
	errNormOrder           = Error{"lin: invalid norm order for matrix"}
	errColAccess           = Error{"lin: column index out of range"}
	errRowAccess           = Error{"lin: row index out of range"}
	errRowLength           = Error{"lin: row length mismatch"}
	errVectorAccess        = Error{"lin: vector index out of range"}
	errZeroLengthMat       = Error{"lin: zero length in matrix dimension"}
	errZeroLengthVec       = Error{"lin: zero length in vector dimension"}
	errIndexOutOfRange     = Error{"lin: index out of range"}
	errNegativeDimension   = Error{"lin: negative dimension"}
	errSliceLengthMismatch = Error{"lin: input slice length mismatch"}
)
