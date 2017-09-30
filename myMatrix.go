package main

import (
	"fmt"
)

type MyMatrix struct {
	numberRow int
	numberColumn int
	matrix [][]int
}

func matrixAdd(matrix1, matrix2 MyMatrix) MyMatrix {
	r := matrix1.numberRow
	c := matrix1.numberColumn
	roseMatrix := make([][]int, r)
	for i := 0 ; i < r ; i++ {
		roseMatrix[i] = make([]int , c)
		for j := 0 ; j < c ; j++ {
			roseMatrix[i][j] = matrix1.matrix[i][j] + matrix2.matrix[i][j] 
		}
	}
	var roseAns = MyMatrix {
		r,
		c,
		roseMatrix,
	}
	return roseAns
}

func matrixDot (matrix1, matrix2 MyMatrix) MyMatrix {
	r := matrix1.numberRow
	c := matrix2.numberColumn

	ansMatrixDot := make([][]int, r)
	for index := range ansMatrixDot {
		ansMatrixDot[index] = make([]int, c)
	}

	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			ansMatrixDot[i][j] = dotIJ(i , j , matrix1 , matrix2)
		}
	}
	var roseAnsMatrix MyMatrix
	roseAnsMatrix.numberRow = r
	roseAnsMatrix.numberColumn = c
	roseAnsMatrix.matrix = ansMatrixDot
	return roseAnsMatrix
}

func dotIJ (i int, j int, matrix1 MyMatrix, matrix2 MyMatrix) int {
	var sum int
	for k := 0 ; k < matrix1.numberColumn ; k++ {
		sum += matrix1.matrix [i][k] * matrix2.matrix[k][j]
	}
	return sum 
}







func matrixScale(number int, inputMatrix MyMatrix) MyMatrix {
	length := inputMatrix.numberRow
	width := inputMatrix.numberColumn

	ansMatrix := make([][]int, length)
	// fmt.Println(ansMatrix)
	for index := range ansMatrix {
		ansMatrix[index] = make([]int, width)
		// fmt.Println(ansMatrix)
	}

	// originalMatrix := inputMatrix.matrix
	for i := 0 ; i < length ; i++ {
		for j := 0 ; j < width ; j++ {
			ansMatrix[i][j] = number * inputMatrix.matrix[i][j] 
		}
	}

	var ansMyMatrix MyMatrix
	ansMyMatrix.numberColumn = width
	ansMyMatrix.numberRow = length
	ansMyMatrix.matrix = ansMatrix
	return ansMyMatrix
	// ansMatrix[0][1] = number * matrix[0][1]
	// ansMatrix[0][0] = number * matrix[0][0]
	// ansMatrix[0][2] = number * matrix[0][2]
	// ansMatrix[1][0] = number * matrix[1][0]
	// ansMatrix[1][1] = number * matrix[1][1]
	// ansMatrix[1][2] = number * matrix[1][2]
	// ansMatrix[2][0] = number * matrix[2][0]
	// ansMatrix[2][1] = number * matrix[2][1]
	// ansMatrix[2][2] = number * matrix[2][2]
}

func (mtx *MyMatrix) show() {
	//for rowIndex:= range mtx.matrix {
	for rowIndex := 0; rowIndex < len(mtx.matrix); rowIndex++ {
		fmt.Println(mtx.matrix[rowIndex])
	}
}

func main() {

	width := 3
	height := 3
	matrix := make([][]int, 3)
	row1 := []int{1, 0, 0}
	row2 := []int{0, 1, 0}
	row3 := []int{0, 0, 1}
	matrix[0] = row1
	matrix[1] = row2
	matrix[2] = row3

	identityMatrix := MyMatrix{height, width, matrix}

	
	// identityMatrix.show();

	danielMatrix := matrixScale(3, identityMatrix)
	danielMatrix.show()

	matrix1 := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
	}
	myMatrix1 := MyMatrix{2, 3, matrix1}

	matrix2 := [][]int{
		[]int{5, 1, 3},
		[]int{7, 0, 8},
	}
	myMatrix2 := MyMatrix{2, 3, matrix2}

	roseMatrix := matrixAdd(myMatrix1, myMatrix2)
	roseMatrix.show()

	matrix3 := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
	}
	myMatrix3 := MyMatrix{2, 3, matrix3}

	matrix4 := [][]int{
		[]int{5, 1},
		[]int{7, 0},
		[]int{3, 4},
	}
	myMatrix4 := MyMatrix{3, 2, matrix4}

	roseMatrix = matrixDot(myMatrix3, myMatrix4)
	roseMatrix.show()
}