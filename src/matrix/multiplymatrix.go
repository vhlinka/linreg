package matrix

import (
	"fmt"
)

//
// Multiply an Matrix - a = a * b
//
func MatrixMultiply(a [][]float64, b [][]float64) ([][]float64, error) {

	var err error

	arows := len(a)
	acols := len(a[0])
	brows := len(b)
	bcols := len(b[0])

	err = nil
	if acols != brows {
		err = fmt.Errorf("Number of Columns in the Matrix a must equal number rows in b")
		return nil, err
	}

	// creat a new target matrix
	newarray := make([][]float64, arows)
	frow := make([]float64, arows*(bcols))
	for i := range newarray {
		newarray[i], frow = frow[:bcols], frow[bcols:]
	}

	for i := range newarray {
		for j := range newarray[i] {
			for k := 0; k < brows; k++ {
				newarray[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return newarray, err
}

//
// Multiply an Matrix by a vector
//
func MatrixMultiplyVector(a [][]float64, v []float64) ([]float64, error) {

	var err error

	nrows := len(a)
	ncols := len(a[0])
	vlen := len(v)

	err = nil
	if ncols != vlen {
		err = fmt.Errorf("Number of Columns in the Matrix must equal number rows in Vector")
		return nil, err
	}

	// create a new array/vector with the same number of rows as the matrix

	newvector := make([]float64, nrows)
	for i := range a {
		for j := range a[i] {
			newvector[i] += a[i][j] * v[j]
		}
	}

	return newvector, err
}
