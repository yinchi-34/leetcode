import pytest

from testutil import load_solution

max_length = load_solution(__file__).Solution().maxLength

CASES = [
    ("example_1", [1, 2, 1, 2, 1, 1, 1], 5),
    ("example_2", [2, 3, 4, 5, 6], 3),
    ("example_3", [1, 2, 3, 1, 4, 5, 1], 5),
    ("two_equal_non_coprime", [2, 2], 2),
    ("no_coprime_neighbors", [6, 4, 10, 8], 2),
    ("all_ones", [1, 1, 1, 1], 4),
    ("whole_array_pairwise_coprime", [1, 2, 3, 5, 7], 5),
    ("largest_coprime_product", [8, 9, 5, 7, 1, 1], 6),
    ("window_drops_several", [3, 5, 7, 2, 10, 3], 4),
]


@pytest.mark.parametrize(
    "nums, expected", [c[1:] for c in CASES], ids=[c[0] for c in CASES]
)
def test_max_length(nums, expected):
    assert max_length(nums) == expected
