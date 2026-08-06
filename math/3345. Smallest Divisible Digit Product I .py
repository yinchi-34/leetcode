class Solution:
    def smallestNumber(self, n: int, t: int) -> int:
        # 计算数位乘积
        def check(num):
            product = 1
            num = n
            while num > 0:
                product *= num % 10
                num //= 10
                if product == 0:
                    break
            return product % t == 0

        while not check(n):
            n += 1

        return n
