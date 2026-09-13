class Solution(object):
    def largestOverlap(self, A, B):
        N = len(A)
        conut = collections.Counter()

        for i, row in enumerate(A):
            for j, v in enumerate(row):
                if v:
                    for i2, row2 in enumerate(B):
                        for j2, v2 in enumerate(row2):
                            if v2:
                                d = [i - i2, j - j2]
                                count[d] += 1

        return max(count.values() or [0])
