from collections import deque
from typing import Optional

import pytest

from testutil import load_solution

solution = load_solution(__file__)
TreeNode = solution.TreeNode
max_depth = solution.Solution().maxDepth


def build_tree(vals: list[Optional[int]]) -> Optional[TreeNode]:
    """Build a tree from LeetCode's level-order format, None for a missing
    child; trailing Nones may be omitted."""
    if not vals or vals[0] is None:
        return None
    root = TreeNode(vals[0])
    queue = deque([root])
    i = 1
    while i < len(vals):
        node = queue.popleft()
        if vals[i] is not None:
            node.left = TreeNode(vals[i])
            queue.append(node.left)
        i += 1
        if i < len(vals) and vals[i] is not None:
            node.right = TreeNode(vals[i])
            queue.append(node.right)
        i += 1
    return root


def chain(n: int, left: bool) -> Optional[TreeNode]:
    """n nodes linked through left children only, or right children only."""
    root = None
    for i in range(n):
        root = TreeNode(i % 201 - 100, root, None) if left else TreeNode(i % 201 - 100, None, root)
    return root


CASES = [
    ("example_1", lambda: build_tree([3, 9, 20, None, None, 15, 7]), 3),
    # Only a right child: looking only at the left subtree would give 1.
    ("example_2_right_only", lambda: build_tree([1, None, 2]), 2),
    ("empty_tree", lambda: None, 0),
    ("single_node", lambda: build_tree([0]), 1),
    # The left subtree is deeper: looking only at the right would give 2.
    ("left_deeper", lambda: build_tree([1, 2, 3, 4, None, None, None, 5]), 4),
    # One side empty: writing min instead of max would give 1.
    ("one_side_empty", lambda: build_tree([1, None, 2, None, 3, None, 4]), 4),
    ("perfect_tree_depth_3", lambda: build_tree([1, 2, 3, 4, 5, 6, 7]), 3),
    # 10^4 nodes in a chain: the recursion is 10^4 levels deep, past
    # CPython's default limit of 1000.
    ("left_chain_max_nodes", lambda: chain(10_000, True), 10_000),
    ("right_chain_max_nodes", lambda: chain(10_000, False), 10_000),
]


@pytest.mark.parametrize(
    "make_tree, expected", [c[1:] for c in CASES], ids=[c[0] for c in CASES]
)
def test_max_depth(make_tree, expected):
    assert max_depth(make_tree()) == expected
