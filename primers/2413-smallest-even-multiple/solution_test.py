import pytest

from testutil import load_solution

smallest_even_multiple = load_solution(__file__).Solution().smallestEvenMultiple


@pytest.mark.parametrize("n, expected", [(5, 10), (6, 6), (1, 2), (150, 150)])
def test_smallest_even_multiple(n, expected):
    assert smallest_even_multiple(n) == expected
