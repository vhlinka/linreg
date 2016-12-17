//
// This file contains a series of element-wise operations - such as addition, subtraction, division
// and multiplication
//
//
package matrix

import (
	"fmt"
)

//
// Subtract Matrix a from b and return as a new matrix
//
func MatrixSubtract(a [][]float64, b [][]float64) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

	if nrows != len(b) || ncols != len(b[0]) {
		err = fmt.Errorf("Number of Columns and Rows must be the same between the two matrix")
		return nil, err
	}

	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(ncols))
	for i := range newarray {
		newarray[i], frow = frow[:ncols], frow[ncols:]
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			newarray[i][j] = a[i][j] - b[i][j]
		}
	}

	return newarray, err
}

//
// Element-wise devide Matrix a by b and return as a new matrix
//
func MatrixElementwiseDevide(a [][]float64, b [][]float64) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

	if nrows != len(b) || ncols != len(b[0]) {
		err = fmt.Errorf("Number of Columns and Rows must be the same between the two matrix")
		return nil, err
	}

	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(ncols))
	for i := range newarray {
		newarray[i], frow = frow[:ncols], frow[ncols:]
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			newarray[i][j] = a[i][j] / b[i][j]
		}
	}

	return newarray, err
}

//
// Element-wise devide Matrix a by b and return as a new matrix
// NOTE: This code implements the Octave "Broadcast" feature
//
//
func MatrixElementwiseMultiply(a [][]float64, b [][]float64, broadcast ...bool) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

	if nrows != len(b) || ncols != len(b[0]) {
		err = fmt.Errorf("Number of Columns and Rows must be the same between the two matrix")
		return nil, err
	}

	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(ncols))
	for i := range newarray {
		newarray[i], frow = frow[:ncols], frow[ncols:]
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			newarray[i][j] = a[i][j] * b[i][j]
		}
	}

	return newarray, err
}

//
// Element-wise devide Matrix a by b and return as a new matrix
// NOTE: This code implements the Octave "Broadcast" feature
//
//
func ElementwiseScaler(a [][]float64, s float64) ([][]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])

	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(ncols))
	for i := range newarray {
		newarray[i], frow = frow[:ncols], frow[ncols:]
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			newarray[i][j] = a[i][j] * s
		}
	}

	return newarray, err
}
