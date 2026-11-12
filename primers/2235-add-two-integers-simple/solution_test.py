import pytest

from testutil import load_solution

sum_ = load_solution(__file__).Solution().sum


@pytest.mark.parametrize(
    "num1, num2, expected",
    [(5, 10, 15), (-5, -10, -15), (-5, 10, 5), (0, 0, 0)],
    ids=["positive_numbers", "negative_numbers", "mixed_numbers", "zero"],
)
def test_sum(num1, num2, expected):
    assert sum_(num1, num2) == expected
