import pytest

from testutil import load_solution

score = load_solution(__file__).Solution().scoreOfParentheses


@pytest.mark.parametrize(
    "s, expected",
    [("()", 1), ("(())", 2), ("()()", 2), ("(()(()))", 6), ("((()))()", 5)],
)
def test_score_of_parentheses(s, expected):
    assert score(s) == expected
