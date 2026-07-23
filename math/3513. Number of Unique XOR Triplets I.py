"""
You are given an integer array nums of length n, where nums is a permutation of the numbers in the range [1, n].

A XOR triplet is defined as the XOR of three elements nums[i] XOR nums[j] XOR nums[k] where i <= j <= k.

Return the number of unique XOR triplet values from all possible triplets (i, j, k).

"""


class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        n = len(nums)
        if n <= 2:
            return n
        res = 1

        while res <= n:
            res <<= 1
        return res