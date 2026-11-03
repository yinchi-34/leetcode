package solution

import (
	"math/rand"
	"slices"
	"strings"
	"testing"

	countkey "leetcode-go/data-structures/hash-table/0049-group-anagrams/count-key"
	sortedkey "leetcode-go/data-structures/hash-table/0049-group-anagrams/sorted-key"
)

var solutions = []struct {
	name string
	fn   func([]string) [][]string
}{
	{"sorted_key", sortedkey.GroupAnagrams},
	{"count_key", countkey.GroupAnagrams},
}

// normalize puts a grouping into a canonical order so that two groupings can
// be compared: the groups are sorted lexicographically, and the strings inside
// each group keep the order the solution returned them in. The expected
// groups list strings in input order, so the tests also check that both
// solutions keep input order inside a group.
func normalize(groups [][]string) [][]string {
	res := make([][]string, len(groups))
	for i, g := range groups {
		res[i] = slices.Clone(g)
	}
	slices.SortFunc(res, func(a, b []string) int {
		return slices.Compare(a, b)
	})
	return res
}

func equalGroupings(a, b [][]string) bool {
	return slices.EqualFunc(normalize(a), normalize(b), slices.Equal[[]string])
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "Example 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}},
		},
		{
			name:     "Example 2: single empty string",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3: single letter",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "Several empty strings form one group",
			strs:     []string{"", "b", ""},
			expected: [][]string{{"", ""}, {"b"}},
		},
		{
			name:     "Duplicate strings stay in one group",
			strs:     []string{"abc", "abc", "cba"},
			expected: [][]string{{"abc", "abc", "cba"}},
		},
		{
			name:     "Same letter set, different multiplicities",
			strs:     []string{"aab", "abb", "bab", "aba"},
			expected: [][]string{{"aab", "aba"}, {"abb", "bab"}},
		},
		{
			name:     "Different lengths never group",
			strs:     []string{"a", "aa", "aaa"},
			expected: [][]string{{"a"}, {"aa"}, {"aaa"}},
		},
		{
			name:     "No two strings are anagrams",
			strs:     []string{"ab", "ac", "bc"},
			expected: [][]string{{"ab"}, {"ac"}, {"bc"}},
		},
		{
			// Counts (11, 0, ...) and (1, 10, 0, ...) both print as "110" followed
			// by 24 zeros, so a key built by concatenating counts without a
			// separator would wrongly merge these two strings.
			name:     "Letter counts that collide if concatenated without separators",
			strs:     []string{strings.Repeat("a", 11), "a" + strings.Repeat("b", 10)},
			expected: [][]string{{strings.Repeat("a", 11)}, {"a" + strings.Repeat("b", 10)}},
		},
		{
			name:     "All 26 letters, shuffled",
			strs:     []string{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", "abcdefghijklmnopqrstuvwxyy"},
			expected: [][]string{{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}, {"abcdefghijklmnopqrstuvwxyy"}},
		},
	}

	for _, sol := range solutions {
		for _, test := range tests {
			t.Run(sol.name+"/"+test.name, func(t *testing.T) {
				result := sol.fn(test.strs)
				if !equalGroupings(result, test.expected) {
					t.Errorf("GroupAnagrams(%q) = %q, want %q (any group order)", test.strs, result, test.expected)
				}
			})
		}
	}
}

// isAnagram reports whether a and b are anagrams by counting letters.
func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var cnt [26]int
	for i := 0; i < len(a); i++ {
		cnt[a[i]-'a']++
		cnt[b[i]-'a']--
	}
	return cnt == [26]int{}
}

// bruteForce groups strs by comparing each string with the first string of
// every existing group, in O(n^2 k) time. It is the reference for random tests.
func bruteForce(strs []string) [][]string {
	var res [][]string
	for _, s := range strs {
		placed := false
		for i := range res {
			if isAnagram(res[i][0], s) {
				res[i] = append(res[i], s)
				placed = true
				break
			}
		}
		if !placed {
			res = append(res, []string{s})
		}
	}
	return res
}

func TestGroupAnagramsRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(49))
	for iter := 0; iter < 500; iter++ {
		n := 1 + rng.Intn(30)
		alphabet := 1 + rng.Intn(4) // few letters, so anagrams are common
		strs := make([]string, n)
		for i := range strs {
			b := make([]byte, rng.Intn(5))
			for j := range b {
				b[j] = byte('a' + rng.Intn(alphabet))
			}
			strs[i] = string(b)
		}
		want := bruteForce(strs)
		for _, sol := range solutions {
			got := sol.fn(strs)
			if !equalGroupings(got, want) {
				t.Fatalf("%s: GroupAnagrams(%q) = %q, want %q (any group order)", sol.name, strs, got, want)
			}
		}
	}
}
