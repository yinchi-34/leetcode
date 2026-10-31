class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        @cache  # 缓存装饰器，避免重复计算 dfs（一行代码实现记忆化）
        def dfs(i: int, j: int, k: int) -> int:
            if i < 0 or j < 0:  # 出界
                return 0
            if i == 0 and j == 0:  # 到达起点
                return 1
            if grid[i][j] == 0:  # 没有镜子，随便走
                return (dfs(i, j - 1, 0) + dfs(i - 1, j, 1)) % 1_000_000_007
            if k == 0:  # 从右边过来
                return dfs(i - 1, j, 1)  # 反射到上边
            # 从下边过来
            return dfs(i, j - 1, 0)  # 反射到左边

        m, n = len(grid), len(grid[0])
        return dfs(m - 1, n - 1, 0)  # 从终点出发