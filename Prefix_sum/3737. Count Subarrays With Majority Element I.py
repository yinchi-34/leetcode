"""
目标：
target 出现次数 > 子数组长度 / 2


"""
from typing import List


class Count_majority_element_version1:
    def countMajoritySubarrays(self, nums: List[int], target: int) -> int:
        result = 0
        n = len(nums)

        for i in range(n):
            tmp = 0
            for j in range(i, n):
                tmp += (1 if nums[j] == target else -1)
                if tmp > 0:
                    result += 1

        return result