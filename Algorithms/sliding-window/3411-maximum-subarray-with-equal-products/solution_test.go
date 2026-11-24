package solution

import "testing"

func TestMaxLength(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 1, 2, 1, 1, 1},
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 3, 4, 5, 6},
			expected: 3,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3, 1, 4, 5, 1},
			expected: 5,
		},
		{
			name:     "Two equal non-coprime elements",
			nums:     []int{2, 2},
			expected: 2,
		},
		{
			name:     "No coprime neighbors",
			nums:     []int{6, 4, 10, 8},
			expected: 2,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Whole array pairwise coprime",
			nums:     []int{1, 2, 3, 5, 7},
			expected: 5,
		},
		{
			name:     "Largest coprime product",
			nums:     []int{8, 9, 5, 7, 1, 1},
			expected: 6,
		},
		{
			name:     "Window must drop several elements",
			nums:     []int{3, 5, 7, 2, 10, 3},
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxLength(test.nums)
			if result != test.expected {
				t.Errorf("maxLength(%v) = %d, want %d", test.nums, result, test.expected)
			}
		})
	}
}
