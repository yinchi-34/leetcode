from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        """Key each string by its 26 letter counts. A list cannot be a dict
        key, so the counts are frozen into a tuple. strs must contain only
        lowercase English letters."""
        groups: defaultdict[tuple[int, ...], list[str]] = defaultdict(list)
        for s in strs:
            cnt = [0] * 26
            for c in s:
                cnt[ord(c) - ord("a")] += 1
            groups[tuple(cnt)].append(s)
        return list(groups.values())
