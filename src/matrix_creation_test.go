package main

import (
	"testing"
)

func TestMatrixCreation(t *testing.T) {
	t.Log("Testing matrix_creation...")
	matrix := BuildMatrix(Matrix{
		iSize: 2,
		jSize: 3,
	})
	if len(matrix) < 2 {
		t.Errorf("Expected a matrix of length 2, but it was %d instead.", len(matrix))
	}
}

func TestMatrixCreationInitialized(t *testing.T) {
	t.Log("Testing matrix_creation...")
	initializer := []int{1, 2, 3, 4, 5, 6}
	matrix := BuildMatrix(Matrix{
		iSize:          2,
		jSize:          3,
		initialization: initializer,
	})
	if len(matrix) < 2 {
		t.Errorf("Expected a matrix of length 2, but it was %d instead.", len(matrix))
	}
}
