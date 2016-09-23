/*
   build a matrix and perform an operation on it.
   give a row size, a column size, and an (optional) initialization
   you don't have to initialize the matrix, it will default to 0's.
*/

package main

import "fmt"

// Matrix is a matrix initialization it has a row and column size and an initialization
type Matrix struct {
	iSize, jSize   int
	initialization []int
}

func main() {
	initializer := []int{1, 2, 3, 4, 5, 6}
	matrix := BuildMatrix(Matrix{
		iSize:          2,
		jSize:          3,
		initialization: initializer,
	})
	fmt.Println(matrix)
}

// BuildMatrix builds a matrix of a given size with a slice of numbers
func BuildMatrix(m Matrix) [][]int {
	var output [][]int
	output = make([][]int, m.iSize)
	for i := 0; i < m.iSize; i++ {
		innerArray := make([]int, m.jSize)
		if len(m.initialization) > 0 {
			for position, val := range m.initialization[:m.jSize] {
				innerArray[position] = val
			}
			m.initialization = m.initialization[m.jSize:]
		}
		output[i] = innerArray
	}
	return output
}
