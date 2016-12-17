package matrix

import (
	"fmt"
)

//
// Add Matrix a and b and return as a new matrix
//
func MatrixAdd(a [][]float64, b [][]float64) ([][]float64, error) {

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
			newarray[i][j] = a[i][j] + b[i][j]
		}
	}

	return newarray, err
}
