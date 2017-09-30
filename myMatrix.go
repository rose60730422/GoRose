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
	// 1. 讀取 matrix1, matrix2 的長寬
	r := matrix1.numberRow
	c := matrix1.numberColumn

	// 2. 宣告用來回傳的 MyMatrix
	roseAns := MyMatrix{}
	roseAns.numberColumn = c
	roseAns.numberColumn = r

	// 3. 計算 MyMatrix 中的 matrix
	roseAns.matrix = make([][]int, r)
	for i := 0 ; i < r ; i++ {
		roseAns.matrix[i] = make([]int , c)
		for j := 0 ; j < c ; j++ {
			roseAns.matrix[i][j] = matrix1.matrix[i][j] + matrix2.matrix[i][j] 
		}
	}

	// 4. 回傳
	return roseAns
}

func matrixDot (matrix1, matrix2 MyMatrix) MyMatrix {
	r := matrix1.numberRow
	c := matrix2.numberColumn

	roseAnsMatrix := MyMatrix{}
	roseAnsMatrix.numberColumn = c
	roseAnsMatrix.numberRow = r
	roseAnsMatrix.matrix = make([][]int, r)

	for index := range roseAnsMatrix.matrix {
		roseAnsMatrix.matrix[index] = make([]int, c)
	}

	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			roseAnsMatrix.matrix[i][j] = dotIJ(i , j , matrix1 , matrix2)
		}
	}
	return roseAnsMatrix
}

func dotIJ (i int, j int, matrix1 MyMatrix, matrix2 MyMatrix) int {
	var sum int
	for k := 0 ; k < matrix1.numberColumn ; k++ {
		sum += matrix1.matrix [i][k] * matrix2.matrix[k][j]
	}
	return sum 
}

func transposeMatrix (inputMatrix MyMatrix) MyMatrix {
	r := inputMatrix.numberRow
	c := inputMatrix.numberColumn

	transpose := MyMatrix{}
	transpose.numberColumn = r
	transpose.numberRow = c
	transpose.matrix = make([][]int , c)
	for index := range transpose.matrix {
		transpose.matrix[index] = make([]int , r)
	}

	for i := 0 ; i < c ; i++ {
		for j := 0 ; j < r ; j++{
			transpose.matrix[i][j] = inputMatrix.matrix[j][i]
		}
	}
	return transpose
}





func matrixScale(number int, inputMatrix MyMatrix) MyMatrix {
	// 1.
	r := inputMatrix.numberRow
	c := inputMatrix.numberColumn
	
	//2.
	ansMatrix := MyMatrix{}
	ansMatrix.numberColumn = c
	ansMatrix.numberRow = r
	ansMatrix.matrix = make([][]int, r)
	// fmt.Println(ansMatrix)

	// 3.
	for index := range ansMatrix.matrix {
		ansMatrix.matrix[index] = make([]int, c)
		// fmt.Println(ansMatrix)
	}

	// originalMatrix := inputMatrix.matrix

	// 4.
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			ansMatrix.matrix[i][j] = number * inputMatrix.matrix[i][j] 
		}
	}
	return ansMatrix
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

	fmt.Println("transpose")
	transposeMatrix := transposeMatrix(myMatrix4)
	transposeMatrix.show()

}