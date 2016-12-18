package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"mat"
	"matrix"
	"os"
	"stats"
	"strconv"
)

func main() {

	f, err := os.Open("/home/vhlinka/testdata/ex1data2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "linreg: %v\n", err)
	}

	// Use strings.NewReader.
	// ... This creates a new Reader for passing to csv.NewReader.
	r := csv.NewReader(bufio.NewReader(f))
	// Read all records.
	result, _ := r.ReadAll()

	fmt.Printf("Lines: %v", len(result))
	fmt.Println()

	nrows := len(result)
	ncols := len(result[0])
	fmt.Printf("array 'result' has %v ROWS and %v COLS\n", nrows, ncols)

	// make a float64 version of the array of data read from file
	farray := make([][]float64, nrows)
	frow := make([]float64, nrows*ncols)
	for i := range farray {
		farray[i], frow = frow[:ncols], frow[ncols:]
	}

	for i := range result {
		for j := range result[i] {
			farray[i][j], err = strconv.ParseFloat(result[i][j], 64)
		}
	}

	// ========================= start using the stats packae
	//	mean, stdev := stats.FeatureNormalize(farray)
	//	fmt.Println(mean)
	//	fmt.Println(stdev)

	/*
		for i := range farray {
			for j := range farray[i] {
				fmt.Printf("%v, ", farray[i][j])
			}
		}
	*/
	// create a new matrix that has a new column of 1s added

	/*
		newarray, err := matrix.AddOnesColumn(farray, 0)
		fmt.Println(newarray)

		theta := []float64{1, 1, 1}
		fmt.Println(theta)

		subarray, err := matrix.SubMatrix(newarray, 0, 47, 0, 4)
		fmt.Println(subarray)

		subarray, err = matrix.SubMatrix(newarray, 0, -1, 2, -1)
		fmt.Println(subarray)
	*/

	/*	fmt.Println()
		fmt.Println("Test Matrix Multiply")
		a := [][]float64{{1, 3}, {4, 2}, {1, 1}, {6, 4}, {3, 2}}
		b := [][]float64{{1, 2}, {3, 4}}
		c, err := matrix.MatrixMultiply(a, b)
		fmt.Println(c)
	*/

	//	fmt.Println()
	//	fmt.Println("Test Elemtwise Matrix Multiply")
	//	a := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//	b := [][]float64{{10, 20, 30}, {10, 20, 30}, {10, 20, 30}}
	//	c, err := matrix.MatrixAdd(a, b)
	//	c, err := matrix.MatrixElementwiseMultiply(a, b)
	//	c, err := mat.ElementwiseOperation(a, b, mat.DivisionOperation)
	//	fmt.Println(c)

	//////// ========================= try to implement the linear Regression using Gradient Decent

	//%% Load Data
	//data = load('ex1data2.txt');
	//X = data(:, 1:2);
	X, err := matrix.SubMatrix(farray, 0, matrix.UntilEndOfDimension, 0, 2)
	//y = data(:, 3);
	y, err := matrix.SubMatrix(farray, 0, matrix.UntilEndOfDimension, 2, 3)

	//m = length(y);

	//% Scale features and set them to zero mean
	//[X mu sigma] = featureNormalize(X);

	mean, stdev := stats.FeatureNormalize(X)
	fmt.Println(mean)
	fmt.Println(stdev)

	//% Add intercept term to X
	//X = [ones(m, 1) X];

	newarray, err := matrix.AddOnesColumn(X, 0)
	//	fmt.Println(newarray)

	//alpha = 0.1;
	//num_iters = 50;
	//	alpha := 0.1
	//	numiters := 50

	//% Init Theta and Run Gradient Descent
	//theta = zeros(3, 1);
	theta := [][]float64{{0}, {0}, {0}}

	alpha := 0.1
	m := float64(len(y))
	sc := (alpha * 1.0 / m)

	for iter := 0; iter < 50; iter++ {
		// Error_Term = ( X * theta ) - y;
		errorterm, err := matrix.MatrixMultiply(newarray, theta)
		errorterm, err = matrix.MatrixSubtract(errorterm, y)

		// 	Error_Term_X = X .* Error_Term;
		errortermx, err := mat.ElementwiseOperation(newarray, errorterm, mat.MultiplicationOperation)

		// 	Sum_Errors = sum(Error_Term_X);
		sumerrors, err := mat.SumArray(errortermx, mat.ColumnDimension)

		// 	theta = theta - (( alpha * 1/m ) * Sum_Errors)';
		newtheta, err := mat.ScalerOperation(sumerrors, sc, mat.MultiplicationOperation)
		newtheta, err = mat.Transpose(newtheta)
		theta, err = mat.ElementwiseOperation(theta, newtheta, mat.SubtractionOperation)

		if err != nil {
			fmt.Println()
			fmt.Println("===== ERROR DETECTED ========")
		}

	}

	fmt.Println("===== NEW THETA ========")
	fmt.Println(theta)

	// use the trained model to predict the price of a new house
	// price = [1 ([1650 3] - mu) ./ sigma] * theta;

	newFeature := [][]float64{{1650.0, 3.0}}
	fmt.Println(newFeature)
	stats.ApplyNormalizeParameters(newFeature, mean, stdev)
	newFeature, err = matrix.AddOnesColumn(newFeature, 0)
	fmt.Println(newFeature)

	answer, err := matrix.MatrixMultiply(newFeature, theta)

	fmt.Println(answer)
	//Predicted price of a 1650 sq-ft, 3 br house (using gradient descent):
	//  $292748.085232
}
