from typing import List


class Solution:
    def busiestServers(
            self,
            k: int,
            arrival: list[int],
            load: list[int]
    ) -> list[int]:
        end_time = [0] * k
        count = [0] * k

        for i in range(len(arrival)):
            start = i % k

            for offset in range(k):
                server = (start + offset) % k
                if end_time[server] <= arrival[i]:
                    end_time[server] = arrival[i] + load[i]
                    count[server] += 1
                    break

        max_count = max(count)
        res = []
        for j in range(k):
            if count[j] == max_count:
                res.append(j)

        return res
