import random

import pytest

from testutil import load_solution

spiral_order = load_solution(__file__).Solution().spiralOrder


def grid(m: int, n: int) -> list[list[int]]:
    """An m x n matrix holding 1, 2, ..., m*n in row-major order."""
    return [[i * n + j + 1 for j in range(n)] for i in range(m)]


CASES = [
    ("example_1", [[1, 2, 3], [4, 5, 6], [7, 8, 9]], [1, 2, 3, 6, 9, 8, 7, 4, 5]),
    ("example_2", [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]],
     [1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]),
    ("single_cell", [[7]], [7]),
    ("single_row", [[1, 2, 3]], [1, 2, 3]),
    ("single_column", [[1], [2], [3]], [1, 2, 3]),
    # Each case below fails if one of the four boundary checks is removed.
    ("wide_2x3_needs_bottom_check", grid(2, 3), [1, 2, 3, 6, 5, 4]),
    ("tall_4x2_needs_left_check", grid(4, 2), [1, 2, 4, 6, 8, 7, 5, 3]),
    ("tall_3x2", grid(3, 2), [1, 2, 4, 6, 5, 3]),
    ("square_4x4", grid(4, 4), [1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10]),
    ("negative_values", [[-1, -2], [-3, -4]], [-1, -2, -4, -3]),
]


@pytest.mark.parametrize(
    "matrix, expected", [c[1:] for c in CASES], ids=[c[0] for c in CASES]
)
def test_spiral_order(matrix, expected):
    assert spiral_order(matrix) == expected


def brute_force(matrix: list[list[int]]) -> list[int]:
    """Walk with a direction array and a visited grid, turning right when the
    next cell is outside the matrix or already visited; stop after m*n steps."""
    m, n = len(matrix), len(matrix[0])
    dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    seen = [[False] * n for _ in range(m)]
    res: list[int] = []
    i = j = d = 0
    for _ in range(m * n):
        res.append(matrix[i][j])
        seen[i][j] = True
        ni, nj = i + dirs[d][0], j + dirs[d][1]
        if not (0 <= ni < m and 0 <= nj < n) or seen[ni][nj]:
            d = (d + 1) % 4
        i, j = i + dirs[d][0], j + dirs[d][1]
    return res


def test_every_shape_up_to_10x10():
    """The constraints allow 1 <= m, n <= 10, so check every shape."""
    for m in range(1, 11):
        for n in range(1, 11):
            matrix = grid(m, n)
            assert spiral_order(matrix) == brute_force(matrix), (m, n)


def test_random_values():
    rng = random.Random(54)
    for _ in range(200):
        m, n = rng.randint(1, 10), rng.randint(1, 10)
        matrix = [[rng.randint(-100, 100) for _ in range(n)] for _ in range(m)]
        assert spiral_order(matrix) == brute_force(matrix), matrix
