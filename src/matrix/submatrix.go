package matrix

import (
	"fmt"
)

//
// return a subset of the source Matrix
//
//	return a[row1:row2, col1:col2]
//		values for row2 and col2 can be set to -1 - telling the routine to use MAX
//
func SubMatrix(a [][]float64, row1 int, row2 int, col1 int, col2 int) ([][]float64, error) {
	var err error
	err = nil

	nrows := len(a)
	ncols := len(a[0])

	if row2 == -1 {
		row2 = nrows
	}
	if col2 == -1 {
		col2 = ncols
	}
	targetrows := row2 - row1
	targetcols := col2 - col1

	if targetrows < 0 || targetrows > nrows {
		err = fmt.Errorf("Invalid row specifications")
		return nil, err
	}

	if targetcols < 0 || targetcols > ncols {
		err = fmt.Errorf("Invalid column specifications")
		return nil, err
	}

	// creat a new target matrix
	newarray := make([][]float64, targetrows)
	frow := make([]float64, targetrows*(targetcols))
	for i := range newarray {
		newarray[i], frow = frow[:targetcols], frow[targetcols:]
	}

	for i, trow := row1, 0; i < row2; i, trow = i+1, trow+1 {
		for j, tcol := col1, 0; j < col2; j, tcol = j+1, tcol+1 {
			newarray[trow][tcol] = a[i][j]
		}
	}

	return newarray, err

}
