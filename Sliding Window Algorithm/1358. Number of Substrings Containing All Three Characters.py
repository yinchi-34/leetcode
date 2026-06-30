from collections import defaultdict

class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        """
        滑动窗口
        它没有存前缀数组，但利用了前缀区间 [0, left - 1] 的数量。

        枚举每个右端点，维护一个尽量左缩的窗口。
        当窗口内同时包含 a、b、c 时，不断移动 left，直到窗口不再满足条件。

        此时所有起点在 [0, left - 1] 的子字符串，
        都可以和当前右端点组成合法子字符串，因此 res += left。
        """
        res, left = 0, 0
        count = defaultdict(int)

        for c in s:
            count[c] += 1
            while len(count) == 3:
                out = s[left]
                count[out] -= 1
                if count[out] == 0:
                    del count[out]

                left += 1
            res += left

        return res