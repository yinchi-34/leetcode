from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        """Key each string by its letters in sorted order. Within a group,
        strings keep their input order; groups come out in the order their
        first member appears, because dict keeps insertion order."""
        groups: defaultdict[str, list[str]] = defaultdict(list)
        for s in strs:
            groups["".join(sorted(s))].append(s)
        return list(groups.values())
