package solution

import (
	"math/rand"
	"reflect"
	"testing"
)

// grid returns an m x n matrix holding 1, 2, ..., m*n in row-major order.
func grid(m, n int) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = i*n + j + 1
		}
	}
	return matrix
}

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:     "Example 2",
			matrix:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:     "Single cell",
			matrix:   [][]int{{7}},
			expected: []int{7},
		},
		{
			name:     "Single row",
			matrix:   [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Single column",
			matrix:   [][]int{{1}, {2}, {3}},
			expected: []int{1, 2, 3},
		},
		// Each case below fails if one of the four boundary checks is removed.
		{
			name:     "Wide 2x3, needs the bottom check",
			matrix:   grid(2, 3),
			expected: []int{1, 2, 3, 6, 5, 4},
		},
		{
			name:     "Tall 4x2, needs the left check",
			matrix:   grid(4, 2),
			expected: []int{1, 2, 4, 6, 8, 7, 5, 3},
		},
		{
			name:     "Tall 3x2",
			matrix:   grid(3, 2),
			expected: []int{1, 2, 4, 6, 5, 3},
		},
		{
			name:     "Square 4x4",
			matrix:   grid(4, 4),
			expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			name:     "Negative values",
			matrix:   [][]int{{-1, -2}, {-3, -4}},
			expected: []int{-1, -2, -4, -3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := spiralOrder(test.matrix)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("spiralOrder(%v) = %v, want %v", test.matrix, result, test.expected)
			}
		})
	}
}

// bruteForce walks with a direction array and a visited grid, turning right
// when the next cell is outside the matrix or already visited; it stops after
// m*n steps.
func bruteForce(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	res := make([]int, 0, m*n)
	i, j, d := 0, 0, 0
	for step := 0; step < m*n; step++ {
		res = append(res, matrix[i][j])
		seen[i][j] = true
		ni, nj := i+dirs[d][0], j+dirs[d][1]
		if ni < 0 || ni >= m || nj < 0 || nj >= n || seen[ni][nj] {
			d = (d + 1) % 4
		}
		i, j = i+dirs[d][0], j+dirs[d][1]
	}
	return res
}

// The constraints allow 1 <= m, n <= 10, so check every shape.
func TestSpiralOrderEveryShape(t *testing.T) {
	for m := 1; m <= 10; m++ {
		for n := 1; n <= 10; n++ {
			matrix := grid(m, n)
			if result, want := spiralOrder(matrix), bruteForce(matrix); !reflect.DeepEqual(result, want) {
				t.Errorf("%dx%d: spiralOrder = %v, want %v", m, n, result, want)
			}
		}
	}
}

func TestSpiralOrderRandomValues(t *testing.T) {
	rng := rand.New(rand.NewSource(54))
	for k := 0; k < 200; k++ {
		m, n := rng.Intn(10)+1, rng.Intn(10)+1
		matrix := make([][]int, m)
		for i := range matrix {
			matrix[i] = make([]int, n)
			for j := range matrix[i] {
				matrix[i][j] = rng.Intn(201) - 100
			}
		}
		if result, want := spiralOrder(matrix), bruteForce(matrix); !reflect.DeepEqual(result, want) {
			t.Errorf("spiralOrder(%v) = %v, want %v", matrix, result, want)
		}
	}
}
