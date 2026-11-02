package solution

import (
	"math/rand"
	"testing"
)

var minMovesTests = []struct {
	name     string
	nums     []int
	expected int
}{
	{name: "example_1", nums: []int{1, 2, 3}, expected: 3},
	{name: "example_2", nums: []int{1, 1, 1}, expected: 0},
	{name: "single_element", nums: []int{7}, expected: 0},
	{name: "unsorted", nums: []int{5, 1, 7, 2}, expected: 11},
	{name: "negative_values", nums: []int{-3, -1, 2}, expected: 7},
	{name: "repeated_minimum", nums: []int{2, 2, 5}, expected: 3},
	{name: "large_gap", nums: []int{-1000000000, 1000000000}, expected: 2000000000},
}

func TestMinMoves(t *testing.T) {
	for _, test := range minMovesTests {
		t.Run(test.name, func(t *testing.T) {
			result := minMoves(test.nums)
			if result != test.expected {
				t.Errorf("minMoves(%v) = %d, want %d", test.nums, result, test.expected)
			}
		})
	}
}

func TestMinMovesBatched(t *testing.T) {
	for _, test := range minMovesTests {
		t.Run(test.name, func(t *testing.T) {
			input := append([]int(nil), test.nums...)
			result := minMovesBatched(input)
			if result != test.expected {
				t.Errorf("minMovesBatched(%v) = %d, want %d", test.nums, result, test.expected)
			}
			for i := range input {
				if input[i] != test.nums[i] {
					t.Fatalf("minMovesBatched modified its input: got %v, want %v", input, test.nums)
				}
			}
		})
	}
}

// simulate performs the moves one at a time, always skipping a maximum.
func simulate(nums []int) int {
	a := append([]int(nil), nums...)
	moves := 0
	for {
		mx, mn := 0, 0
		for i := range a {
			if a[i] > a[mx] {
				mx = i
			}
			if a[i] < a[mn] {
				mn = i
			}
		}
		if a[mx] == a[mn] {
			return moves
		}
		for i := range a {
			if i != mx {
				a[i]++
			}
		}
		moves++
	}
}

func TestAgainstSimulation(t *testing.T) {
	rng := rand.New(rand.NewSource(453))
	for iter := 0; iter < 500; iter++ {
		n := 1 + rng.Intn(6)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(21) - 10
		}
		want := simulate(nums)
		if got := minMoves(nums); got != want {
			t.Fatalf("minMoves(%v) = %d, simulation = %d", nums, got, want)
		}
		if got := minMovesBatched(nums); got != want {
			t.Fatalf("minMovesBatched(%v) = %d, simulation = %d", nums, got, want)
		}
	}
}
