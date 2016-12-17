package matrix

import (
	"fmt"
)

// insert a column of 1.0 values at the specified zero-based column
// note: the new column will be added to the "left" of the column specified
//		  a value of 0 will cause a new column to be inserted as the first colum.
//		  a value of
//
func AddOnesColumn(a [][]float64, pos int) ([][]float64, error) {

	var err error
	err = nil

	nrows := len(a)
	ncols := len(a[0])
	targetcols := ncols + 1 // target will have one more column than the source

	if pos < 0 || pos > targetcols-1 {
		err = fmt.Errorf("Invalid Position value specified. Must between 0 and less than %v", targetcols-1)
		return nil, err
	}
	// creat a new target matrix
	newarray := make([][]float64, nrows)
	frow := make([]float64, nrows*(targetcols))
	for i := range newarray {
		newarray[i], frow = frow[:targetcols], frow[targetcols:]
	}

	sourcecol := 0
	// now copy the values - and add the 1.0 values in the new column
	for j := 0; j < targetcols; j++ {
		if j == pos {
			for i := 0; i < nrows; i++ {
				newarray[i][j] = 1.0
			}
		} else {
			for i := 0; i < nrows; i++ {
				newarray[i][j] = a[i][sourcecol]
			}
			sourcecol++
		}
	}

	return newarray, err
}
