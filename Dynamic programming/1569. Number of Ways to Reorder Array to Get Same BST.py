class Solution:
    def numOfWays(self, nums: list[int]) -> int:
        # 排序
        MOD = 1_000_000_007
        n = len(nums)

        # dp[i][j] = C(i, j)
        dp = [[0] * (n + 1) for _ in range(n + 1)]

        for i in range(n + 1):
            dp[i][0] = dp[i][i] = 1

            for j in range(1, i):
                dp[i][j] = (
                    dp[i - 1][j - 1] + dp[i - 1][j]
                ) % MOD

        def dfs(seq):
            if len(seq) <= 2:
                return 1

            root = seq[0]

            # 保留原有相对顺序，且不包含 root
            left = [x for x in seq[1:] if x < root]
            right = [x for x in seq[1:] if x > root]

            # 左右子树的交错排列数
            ways = dp[len(seq) - 1][len(left)]

            # 左右子树各自的排列数
            ways = ways * dfs(left) % MOD
            ways = ways * dfs(right) % MOD

            return ways

        # 排除原始排列
        return (dfs(nums) - 1) % MOD