class Solution:
    def resultArray(self, nums: List[int], k: int) -> List[int]:
        """
        题意理解：
            即有多少个连续子数组，元素乘积 % k = x
            隐藏条件：x 的区间是[0, k - 1]

        概念上的暴力代码：
        result = [0] * k
        for left in range(len(nums)):
            product = 1

            for right in range(left, len(nums)):
                product *= nums[right]

                x = product % k
                result[x] += 1

        改成遍历到当前元素，有多少个子数组，以及当前子数组的乘积余数是多少
            即DP
        """
        result = [0] * k
        dp = [0] * k

        for num in nums:
            new_dp = [0] * k

            for r in range(k):
                new_r = (r * num) % k
                new_dp[new_r] += dp[r]

            # 单独的 [num]
            new_dp[num % k] += 1

            # 当前这些子数组加入总答案
            for r in range(k):
                result[r] += new_dp[r]

            dp = new_dp

        return result