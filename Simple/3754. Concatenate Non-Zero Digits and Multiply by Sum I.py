class Solution:
    def sumAndMultiply(self, n: int) -> int:
        sum = 0
        res = 0

        for i in str(n):
            if int(i) != 0:
                sum += int(i)
                res = res * 10 + int(i)

        return res * sum