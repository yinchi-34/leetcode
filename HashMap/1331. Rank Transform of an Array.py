"""
数据是静态还是动态？
        ↓
是否需要保持大小关系？
        ↓
值域大还是不同值数量大？
        ↓
输入是否已经有序？
        ↓
元素是否可哈希？
        ↓
排名是否允许并列？
        ↓
升序还是降序？
        ↓
在内存处理还是数据库处理？

"""

from typing import List


class Solution:
    def arrayRankTransform(self, arr: List[int]) -> List[int]:
        # 哈希（索引）+ 排序
        # 适用场景
        ranks = {v:i for i, v in enumerate(sorted(set(arr)), 1)}
        return [ranks[value] for value in arr]