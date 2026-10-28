class Solution:
    def minCost(self, grid: list[list[int]]) -> int:
        visited = set()
        res = inf

        def dfs(i: int, j: int, xor: int) -> None:
            nonlocal res
            if res == 0 or i < 0 or j < 0 or (i, j, xor) in visited:
                return

            visited.add((i, j, xor))
            xor ^= grid[i][j]

            if i == 0 and j == 0:
                res = min(res, xor)
                return
            dfs(i - 1, j, xor)
            dfs(i, j - 1, xor)

        dfs(len(grid) - 1, len(grid[0]) - 1, 0)

        return res