package chapter1

func zeroMatrix(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])

	rowsToZero := make([]bool, rows)
	columnsToZero := make([]bool, columns)

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if matrix[row][column] == 0 {
				rowsToZero[row] = true
				columnsToZero[column] = true
			}
		}
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if rowsToZero[row] || columnsToZero[column] {
				matrix[row][column] = 0
			}
		}
	}
}

func zeroMatrixV2(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])

	toZero := make([]bool, rows+columns)

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if matrix[row][column] == 0 {
				toZero[row] = true
				toZero[rows+column] = true
			}
		}
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if toZero[row] || toZero[rows+column] {
				matrix[row][column] = 0
			}
		}
	}
}

func zeroMatrixV3(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if matrix[row][column] == 0 {
				matrix[0][column] = 0
				matrix[row][0] = 0
			}
		}
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if matrix[0][column] == 0 || matrix[row][0] == 0 {
				matrix[row][column] = 0
			}
		}
	}
}
