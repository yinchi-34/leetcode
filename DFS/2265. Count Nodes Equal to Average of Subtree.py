# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfSubtree(self, root: TreeNode) -> int:
        res = 0

        def dfs(node: Optional[TreeNode]) -> tuple[int, int]:
            nonlocal res
            if node is None:
                return 0, 0

            leftSum, leftSize = dfs(node.left)
            rightSum, rightSize = dfs(node.right)
            Size = leftSize + rightSize + 1
            Sum = leftSum + rightSum + node.val
            if Size and Sum // Size == node.val:
                res += 1
            return Sum, Size

        dfs(root)
        return res