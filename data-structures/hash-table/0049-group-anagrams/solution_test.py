import random

import pytest

from testutil import load_solution

SOLUTIONS = {
    "sorted_key": load_solution(__file__, "sorted-key").Solution().groupAnagrams,
    "count_key": load_solution(__file__, "count-key").Solution().groupAnagrams,
}


def normalize(groups: list[list[str]]) -> list[list[str]]:
    """Sort the groups so two groupings can be compared. Strings inside a
    group keep their order, so the tests also check input order there."""
    return sorted(groups)


CASES = [
    ("example_1", ["eat", "tea", "tan", "ate", "nat", "bat"],
     [["bat"], ["tan", "nat"], ["eat", "tea", "ate"]]),
    ("example_2_single_empty_string", [""], [[""]]),
    ("example_3_single_letter", ["a"], [["a"]]),
    ("several_empty_strings", ["", "b", ""], [["", ""], ["b"]]),
    ("duplicate_strings", ["abc", "abc", "cba"], [["abc", "abc", "cba"]]),
    ("same_letters_different_counts", ["aab", "abb", "bab", "aba"],
     [["aab", "aba"], ["abb", "bab"]]),
    ("different_lengths", ["a", "aa", "aaa"], [["a"], ["aa"], ["aaa"]]),
    ("no_anagrams", ["ab", "ac", "bc"], [["ab"], ["ac"], ["bc"]]),
    # Counts (11, 0, ...) and (1, 10, 0, ...) would collide in a key built by
    # concatenating counts without separators; a tuple key cannot collide.
    ("counts_collide_if_concatenated", ["a" * 11, "a" + "b" * 10],
     [["a" * 11], ["a" + "b" * 10]]),
    ("all_26_letters_shuffled",
     ["abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", "abcdefghijklmnopqrstuvwxyy"],
     [["abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"], ["abcdefghijklmnopqrstuvwxyy"]]),
]


@pytest.mark.parametrize("solution", SOLUTIONS)
@pytest.mark.parametrize(
    "strs, expected", [c[1:] for c in CASES], ids=[c[0] for c in CASES]
)
def test_group_anagrams(solution, strs, expected):
    assert normalize(SOLUTIONS[solution](strs)) == normalize(expected)


def brute_force(strs: list[str]) -> list[list[str]]:
    """Compare each string with the first string of every group, O(n^2 k)."""
    res: list[list[str]] = []
    for s in strs:
        for group in res:
            if sorted(group[0]) == sorted(s):
                group.append(s)
                break
        else:
            res.append([s])
    return res


@pytest.mark.parametrize("solution", SOLUTIONS)
def test_group_anagrams_random(solution):
    rng = random.Random(49)
    for _ in range(500):
        alphabet = "abcd"[: rng.randint(1, 4)]  # few letters, so anagrams are common
        strs = ["".join(rng.choice(alphabet) for _ in range(rng.randint(0, 4)))
                for _ in range(rng.randint(1, 30))]
        assert normalize(SOLUTIONS[solution](strs)) == normalize(brute_force(strs)), strs
