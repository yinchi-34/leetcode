class Solution:
    def maxSubarrayLength(self, nums: List[int], k: int) -> int:
        n, occ = len(nums), Counter()
        right, res = -1, 0

        for left in range(n):
            if left > 0:
                # 初始设置为-1，状态统一
                occ[nums[left - 1]] -= 1
            while right + 1 < n and occ[nums[right + 1]] < k:
                right += 1
                occ[nums[right]] += 1

            res = max(res, right - left + 1)
        return res