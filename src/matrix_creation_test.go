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
