package solution

import "slices"

// GroupAnagrams groups the strings in strs that are anagrams of each other.
// Each string is keyed by its letters in sorted order, so two strings share a
// key exactly when they are anagrams. Within a group, strings keep their input
// order; the order of the groups is unspecified.
func GroupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, s := range strs {
		b := []byte(s)
		slices.Sort(b)
		key := string(b)
		groups[key] = append(groups[key], s)
	}

	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}
