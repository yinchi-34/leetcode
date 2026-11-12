import pytest

from testutil import load_solution

longest_consecutive = load_solution(__file__).Solution().longestConsecutive

CASES = [
    ("example_1", [100, 4, 200, 1, 3, 2], 4),
    ("example_2", [0, 3, 7, 2, 5, 8, 4, 6, 0, 1], 9),
    ("example_3", [1, 0, 1, 2], 3),
    ("empty_array", [], 0),
    ("single_element", [7], 1),
    ("all_duplicates", [5, 5, 5], 1),
    ("negative_numbers", [-1, -3, -2, 1, 0], 5),
    ("two_runs_longer_second", [10, 11, 1, 2, 3, 4], 4),
]


@pytest.mark.parametrize(
    "nums, expected", [c[1:] for c in CASES], ids=[c[0] for c in CASES]
)
def test_longest_consecutive(nums, expected):
    assert longest_consecutive(nums) == expected
