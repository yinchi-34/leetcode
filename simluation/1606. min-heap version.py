import heapq
from sortedcontainers import SortedList

class Solution:
    def busiestServers(self, k: int, arrival: list[int], load: list[int]) -> list[int]:
        # 第一种方法
        # 有序集合+优先队列
        # 有序集合存储空闲服务器的编号
        # 优先队列存储忙碌服务器的结束时间编号
        available = SortedList(range(k))
        busy = []
        count = [0] * k

        for i in range(len(arrival)):
            # 先更新空闲服务器
            while busy and busy[0][0] <= arrival[i]:
                available.add(busy[0][1])
                heapq.heappop(busy)
            if len(available) < 1:
                continue

            # 先找到位置，找在它第一个大于等于应该所在的服务器，应用二分
            j = available.bisect_left(i % k)
            # 如果没找到当前元素bisect _left返回长度，整个服务器是环，所以j=0
            if j == len(available):
                j = 0

            id = available[j]
            count[id] += 1
            heapq.heappush(busy, (arrival[i] + load[i], id))
            available.remove(id)

        max_count = max(count)
        return [k for k, count in enumerate(count) if count == max_count]