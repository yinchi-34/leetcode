class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        m, n = len(grid), len(grid[0])

        visited = [[False] * n for _ in range(m)]
        res = 0

        for r in range(m):
            for c in range(n):
                if grid[r][c] != 0 and visited[r][c] == False:
                    stack = [(r, c)]
                    visited[r][c] = True

                    total = 0
                    while stack:
                        x, y = stack.pop()
                        total += grid[x][y]
                        for nx, ny in [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]:
                            if 0 <= nx < m and 0 <= ny < n and grid[nx][ny] != 0 and not visited[nx][ny]:
                                stack.append((nx, ny))
                                visited[nx][ny] = True
                    if (total % k) == 0:
                        res += 1
        return res