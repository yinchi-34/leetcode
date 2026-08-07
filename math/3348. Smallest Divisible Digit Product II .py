class Solution:
    def smallestNumber(self, s: str, t: int) -> str:
        '''
        res 不含数位0
        res >= num
        且数位乘积能整除t

        t因式分解
        '''
        # 对 t 进行因式分解，看会不会出现超过1位数的质数
        num = t
        for d in [7, 5, 3, 2]:
            while num % d == 0:
                num //= d
        # 出现超过1位数的质数，不满足题意
        if num != 1:
            return "-1"

        # 哨兵补全
        s = '0' * 50 + s
        n = len(s)

        res = []

        # 第i位,t，flag是否比s大
        @cache
        def dfs(i, t, flag):
            if i == n:
                # 满足了
                if t == 1:
                    return True
                return False

            w = int(s[i])
            # 未超过s
            if not flag:
                if w == 0:
                    # 前 50 位可以从0开始填
                    if i < 50:
                        start = 0
                    # 否则要从1开始填
                    else:
                        start = 1
                else:
                    start = w
            # 超过了
            else:
                start = 1

            for nw in range(start, 10, 1):
                res.append(str(nw))
                # 根据gcd(t,nw)，更新 t
                if not nw:
                    nt = t
                else:
                    nt = t // gcd(nw, t)

                if flag or nw != w:
                    nflag = True
                else:
                    nflag = False

                if dfs(i + 1, nt, nflag):
                    return True
                res.pop()

            return False

        dfs(0, t, False)
        dfs.cache_clear()

        # 去掉前缀 0
        for i in range(len(res)):
            if res[i] != '0':
                return ''.join(res[i:])

