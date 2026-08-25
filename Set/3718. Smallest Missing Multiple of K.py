class Solution:
    def missingMultiple(self, nums: List[int], k: int) -> int:
        #  是否存在最好用set，因为底层逻辑是hash
        #  如果为list，循环的时候为O（N）
        #  如果为set，循环是O（1）

        seen = set(nums)

        res = k
        while res in seen:
            res += k
        return res