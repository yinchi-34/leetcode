import random

import pytest

from testutil import load_solution

sorting = load_solution(__file__, "sorting")
threeway = load_solution(__file__, "three-way-quickselect")
lomuto = load_solution(__file__, "lomuto-quickselect")

SOLUTIONS = {
    "sorting": sorting.Solution().minMoves2,
    "three_way_quickselect": threeway.Solution().minMoves2,
    "lomuto_quickselect": lomuto.Solution().minMoves2,
}

SELECTORS = {
    "three_way_quickselect": threeway.quick_select,
    "lomuto_quickselect": lomuto.quick_select,
}

CASES = [
    ("example_1", [1, 2, 3], 2),
    ("example_2", [1, 10, 2, 9], 16),
    ("single_element", [7], 0),
    ("all_equal", [5, 5, 5, 5, 5], 0),
    ("even_length", [1, 1000], 999),
    ("outliers", [1, 2, 3, 1000, 1001], 1998),
    ("negative_values", [-5, 0, 3, -2], 10),
    ("many_duplicates", [0, 1, 0, 1, 1, 0, 1], 3),
    ("extreme_values", [-1000000000, 1000000000, 0], 2000000000),
]


@pytest.mark.parametrize("solution", SOLUTIONS)
@pytest.mark.parametrize(
    "nums, expected", [c[1:] for c in CASES], ids=[c[0] for c in CASES]
)
def test_min_moves2(solution, nums, expected):
    # Pass a copy: the quickselect solutions reorder their input.
    assert SOLUTIONS[solution](nums.copy()) == expected


def test_sorting_keeps_input():
    nums = [3, 1, 2]
    sorting.Solution().minMoves2(nums)
    assert nums == [3, 1, 2]


@pytest.mark.parametrize("selector", SELECTORS)
def test_quick_select(selector):
    """Check every rank k against sorting on arrays with many duplicates,
    and check that the result is a permutation of the input."""
    rng = random.Random(462)
    for _ in range(2000):
        n = rng.randint(1, 10)
        nums = [rng.randint(-2, 2) for _ in range(n)]
        expected = sorted(nums)
        for k in range(n):
            a = nums.copy()
            assert SELECTORS[selector](a, 0, n - 1, k) == expected[k], (nums, k)
            assert sorted(a) == expected, (nums, k)


def brute_force(nums: list[int]) -> int:
    """Try every element as the common target; an optimal target can always
    be chosen among the elements."""
    return min(sum(abs(num - x) for num in nums) for x in nums)


@pytest.mark.parametrize("solution", SOLUTIONS)
def test_against_brute_force(solution):
    rng = random.Random(4620)
    for _ in range(1000):
        nums = [rng.randint(-20, 20) for _ in range(rng.randint(1, 8))]
        assert SOLUTIONS[solution](nums.copy()) == brute_force(nums), nums


def test_all_equal_large():
    """Three-way partitioning finishes all-equal input in one pass. The
    Lomuto solution is left out on purpose, as in the Go tests: it is
    quadratic on this input."""
    assert threeway.Solution().minMoves2([0] * 100000) == 0
