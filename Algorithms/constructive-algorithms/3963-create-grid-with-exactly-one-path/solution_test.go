package solution

import "testing"

// countPaths 统计只向右或向下、只经过 '.' 的从左上到右下的路径数。
func countPaths(grid []string) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' {
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// checkGrid 检查网格尺寸、字符集，以及路径是否恰好一条。
func checkGrid(t *testing.T, m, n int, grid []string) {
	t.Helper()
	if len(grid) != m {
		t.Fatalf("createGrid(%d, %d) has %d rows, want %d", m, n, len(grid), m)
	}
	for i, row := range grid {
		if len(row) != n {
			t.Fatalf("createGrid(%d, %d) row %d has length %d, want %d", m, n, i, len(row), n)
		}
		for j := 0; j < n; j++ {
			if row[j] != '.' && row[j] != '#' {
				t.Fatalf("createGrid(%d, %d) has invalid char %q at (%d, %d)", m, n, row[j], i, j)
			}
		}
	}
	if got := countPaths(grid); got != 1 {
		t.Errorf("createGrid(%d, %d) = %v has %d paths, want 1", m, n, grid, got)
	}
}

func TestCreateGrid(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
	}{
		{name: "example_2x3", m: 2, n: 3},
		{name: "example_3x3", m: 3, n: 3},
		{name: "example_1x4", m: 1, n: 4},
		{name: "single_cell", m: 1, n: 1},
		{name: "single_column", m: 4, n: 1},
		{name: "max_size", m: 25, n: 25},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkGrid(t, test.m, test.n, createGrid(test.m, test.n))
		})
	}
}

func TestCreateGridAllSizes(t *testing.T) {
	for m := 1; m <= 25; m++ {
		for n := 1; n <= 25; n++ {
			checkGrid(t, m, n, createGrid(m, n))
		}
	}
}

func TestCountPaths(t *testing.T) {
	tests := []struct {
		name     string
		grid     []string
		expected int
	}{
		{name: "all_free_2x2", grid: []string{"..", ".."}, expected: 2},
		{name: "all_free_3x3", grid: []string{"...", "...", "..."}, expected: 6},
		{name: "blocked_end", grid: []string{"..", ".#"}, expected: 0},
		{name: "example_answer", grid: []string{"..#", "#..", "##."}, expected: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := countPaths(test.grid); got != test.expected {
				t.Errorf("countPaths(%v) = %d, want %d", test.grid, got, test.expected)
			}
		})
	}
}
