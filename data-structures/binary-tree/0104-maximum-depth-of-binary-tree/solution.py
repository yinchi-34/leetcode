import sys
from typing import Optional

# The tree can be a chain of 10^4 nodes, so the recursion can be 10^4 levels
# deep; CPython stops at 1000 by default. Go needs nothing like this, because
# goroutine stacks grow on demand.
sys.setrecursionlimit(20_000)


class TreeNode:
    """The binary tree node LeetCode provides."""

    def __init__(self, val: int = 0, left: "Optional[TreeNode]" = None,
                 right: "Optional[TreeNode]" = None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0
        return max(self.maxDepth(root.left), self.maxDepth(root.right)) + 1
