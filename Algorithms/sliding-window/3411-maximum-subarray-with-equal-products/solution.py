from math import gcd


class Solution:
    def maxLength(self, nums: list[int]) -> int:
        res = 2  # any subarray of length 2 satisfies prod == lcm * gcd
        prod = 1
        left = 0
        for right, x in enumerate(nums):
            # Shrink until x is coprime to every element in nums[left:right].
            while gcd(prod, x) > 1:
                prod //= nums[left]
                left += 1
            prod *= x
            res = max(res, right - left + 1)
        return res
