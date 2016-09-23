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

// TestMatrixCreationLarge makes a 1 million row by 1000 column matrix ~2.1 seconds on a desktop
// Larger ones (1 billion took 71 seconds) overran the stack
//runtime: VirtualAlloc of 1048576 bytes failed with errno=1455 fatal error: runtime: cannot map pages in arena address space
func TestMatrixCreationLarge(t *testing.T) {
	t.Log("Testing matrix_creation...")
	matrix := BuildMatrix(Matrix{
		iSize: 1000000,
		jSize: 1000,
	})
	if len(matrix) < 2 {
		t.Errorf("Expected a matrix of length 2, but it was %d instead.", len(matrix))
	}
}
