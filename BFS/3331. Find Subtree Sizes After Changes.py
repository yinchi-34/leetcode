class Solution:
    def findSubtreeSizes(self, parent: List[int], s: str) -> List[int]:
        n = len(parent)

        g_old = defaultdict(list)
        for i in range(1, n):
            g_old[parent[i]].append(i)

        records = defaultdict(list)
        res = [0] * n

        def operate(u: int):
            res[u] += 1
            records[s[u]].append(u)
            for v in g_old[u]:
                operate(v)
                u_new = records[s[v]][-1] if records[s[v]] else u
                res[u_new] += res[v]

            records[s[u]].pop()

        operate(0)
        return res