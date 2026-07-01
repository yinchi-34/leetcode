from collections import deque
from typing import List

class Solution:
    def maximumSafenessFactor(self, grid: List[List[int]]) -> int:
        n = len(grid)
        dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        # 1. 多源 BFS：计算每个格子到最近小偷的距离
        dist = [[-1] * n for _ in range(n)]
        q = deque()

        for i in range(n):
            for j in range(n):
                if grid[i][j] == 1:
                    dist[i][j] = 0
                    q.append((i, j))

        while q:
            x, y = q.popleft()
            for dx, dy in dirs:
                nx, ny = x + dx, y + dy
                if 0 <= nx < n and 0 <= ny < n and dist[nx][ny] == -1:
                    dist[nx][ny] = dist[x][y] + 1
                    q.append((nx, ny))

        # 2. 判断安全系数 limit 是否可行
        def check(limit: int) -> bool:
            if dist[0][0] < limit:
                return False

            visited = [[False] * n for _ in range(n)]
            q = deque([(0, 0)])
            visited[0][0] = True

            while q:
                x, y = q.popleft()
                if x == n - 1 and y == n - 1:
                    return True

                for dx, dy in dirs:
                    nx, ny = x + dx, y + dy
                    if (
                        0 <= nx < n
                        and 0 <= ny < n
                        and not visited[nx][ny]
                        and dist[nx][ny] >= limit
                    ):
                        visited[nx][ny] = True
                        q.append((nx, ny))

            return False

        # 3. 二分最大安全系数
        left, right = 0, max(max(row) for row in dist)
        ans = 0

        while left <= right:
            mid = (left + right) // 2
            if check(mid):
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans